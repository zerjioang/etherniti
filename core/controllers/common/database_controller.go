package common

import (
	"errors"
	"fmt"

	"github.com/zerjioang/etherniti/core/api"
	"github.com/zerjioang/etherniti/core/data"
	"github.com/zerjioang/etherniti/core/db"
	"github.com/zerjioang/etherniti/core/logger"
	"github.com/zerjioang/etherniti/core/util/str"
	"github.com/zerjioang/etherniti/shared/mixed"
	"github.com/zerjioang/etherniti/thirdparty/echo"
)

type DatabaseController struct {
	storage *db.BadgerStorage
	name    string
	// data model generator function. it will create a new struct
	modelGenerator func() mixed.DatabaseObjectInterface
	// data model used in this collection
	model mixed.DatabaseObjectInterface
	// path prepend string
	pathPrepend string
}

// constructor like function
func NewDatabaseController(pathPrepend string, collection string, modelGenerator func() mixed.DatabaseObjectInterface) (DatabaseController, error) {
	dbctl := DatabaseController{}
	if collection == "" {
		return dbctl, errors.New("invalid collection name provided")
	}

	lastChar := collection[len(collection)-1:]
	if lastChar != "s" {
		// check if ends with y. example: proxy, registry
		if lastChar == "y" {
			dbctl.name = collection
		} else {
			// collection name does not end with plural. add the 's'
			dbctl.name = collection + "s"
		}
	} else {
		dbctl.name = collection
	}
	dbctl.modelGenerator = modelGenerator
	dbctl.model = modelGenerator()

	// create path prepend data
	if pathPrepend != "" {
		if pathPrepend[0] == '/' {
			// developer entered path prepend starts with slash
			dbctl.pathPrepend = pathPrepend
		} else {
			// developer entered path prepend does not starts with slash
			dbctl.pathPrepend = "/" + pathPrepend
		}
	}

	// database initialization method
	var err error
	dbctl.storage, err = db.NewCollection(collection)
	if err != nil {
		logger.Error("failed to initialize database db collection: ", err)
	}
	return dbctl, err
}

func (ctl *DatabaseController) Create(c *echo.Context) error {
	requestedItem, err := ctl.modelGenerator().Bind(c)
	if err.Occur() {
		return api.StackError(c, err)
	}
	if requestedItem != nil {
		canWriteErr := requestedItem.CanWrite(c)
		if canWriteErr == nil {
			// check if current user Id is valid, exists.
			// source: auth-jwt-token
			authId := c.AuthenticatedUserUuid()
			if authId == "" {
				return errors.New("unauthorized operation detected")
			}
			key := ctl.buildCompositeId(authId, string(requestedItem.Key()))
			value := requestedItem.Value()
			fmt.Println(requestedItem)
			writeErr := ctl.storage.PutUniqueKeyValue(key, value)
			if writeErr != nil {
				return api.Error(c, writeErr)
			} else {
				return api.Success(c, data.SuccessfullyCreated, value)
			}
		} else {
			return api.Error(c, canWriteErr)
		}
	}
	return api.ErrorStr(c, data.FailedToProcess)
}

func (ctl *DatabaseController) GetKey(key []byte) ([]byte, error) {
	return ctl.storage.Get(key)
}

func (ctl *DatabaseController) SetKey(key []byte, value []byte) error {
	return ctl.storage.Set(key, value)
}

func (ctl *DatabaseController) SetUniqueKey(key []byte, value []byte) error {
	return ctl.storage.PutUniqueKeyValue(key, value)
}

func (ctl *DatabaseController) Update(c *echo.Context) error {
	modelId := c.Param("id")
	if modelId != "" {
		canUpdateErr := ctl.model.CanUpdate(c, modelId)
		if canUpdateErr == nil {
			// check if current user Id is valid, exists.
			// source: auth-jwt-token
			authId := c.AuthenticatedUserUuid()
			if authId == "" {
				return errors.New("unauthorized operation detected")
			}
			key := ctl.buildCompositeId(authId, modelId)
			// read current content stored in given key
			projectData, readErr := ctl.storage.Get(key)
			if readErr != nil {
				return api.Error(c, readErr)
			}
			// decode byte content into model
			readedObject, objErr := ctl.model.Decode(projectData)
			if objErr.Occur() {
				return api.StackError(c, objErr)
			}
			// now update readed object content with user provided content in request body
			requestedItem, bindErr := ctl.modelGenerator().Bind(c)
			if bindErr.Occur() {
				return api.StackError(c, bindErr)
			}
			// update item
			updatedItem, updateErr := readedObject.Update(requestedItem)
			if updateErr.Occur() {
				return api.StackError(c, updateErr)
			}
			// save result
			storeErr := ctl.storage.Set(key, updatedItem.Value())
			if storeErr != nil {
				return api.Error(c, storeErr)
			}
			return api.SendSuccessBlob(c, projectData)
		}
		return api.Error(c, canUpdateErr)
	} else {
		return api.ErrorStr(c, data.ProvideId)
	}
}

func (ctl *DatabaseController) Read(c *echo.Context) error {
	modelId := c.Param("id")
	if modelId != "" {
		canReadErr := ctl.model.CanRead(c, modelId)
		if canReadErr == nil {
			// check if current user Id is valid, exists.
			// source: auth-jwt-token
			authId := c.AuthenticatedUserUuid()
			if authId == "" {
				return errors.New("unauthorized operation detected")
			}
			key := ctl.buildCompositeId(authId, modelId)
			projectData, readErr := ctl.storage.Get(key)
			if readErr != nil {
				return api.Error(c, readErr)
			}
			return api.SendSuccessBlob(c, projectData)
		}
		return api.Error(c, canReadErr)
	} else {
		return api.ErrorStr(c, data.ProvideId)
	}
}

func (ctl *DatabaseController) Delete(c *echo.Context) error {
	modelId := c.Param("id")
	if modelId != "" {
		canDeleteErr := ctl.model.CanDelete(c, modelId)
		if canDeleteErr == nil {
			// check if current user Id is valid, exists.
			// source: auth-jwt-token
			authId := c.AuthenticatedUserUuid()
			if authId == "" {
				return errors.New("unauthorized operation detected")
			}
			// build the composite id: authId + modelId
			key := ctl.buildCompositeId(authId, modelId)
			deleteErr := ctl.storage.Delete(key)
			if deleteErr != nil {
				return api.Error(c, deleteErr)
			}
			return api.SendSuccess(c, data.SuccessfullyDeleted, nil)
		}
		return api.Error(c, canDeleteErr)
	} else {
		return api.ErrorStr(c, data.ProvideId)
	}
}

func (ctl *DatabaseController) List(c *echo.Context) error {
	canList := ctl.model.CanList(c)
	if canList == nil {
		results, err := ctl.storage.List("", ctl.model)
		if err != nil {
			return api.Error(c, err)
		} else if results == nil || len(results) == 0 {
			//no data found
			return api.Success(c, str.UnsafeBytes(ctl.name), nil)
		} else {
			return api.SendSuccess(c, str.UnsafeBytes(ctl.name), results)
		}
	}
	return api.ErrorStr(c, data.NotAllowedToList)
}

func (ctl *DatabaseController) ListOwnerOnly(c *echo.Context) error {
	canList := ctl.model.CanList(c)
	if canList == nil {
		// check if current user Id is valid, exists.
		// source: auth-jwt-token
		authId := c.AuthenticatedUserUuid()
		if authId == "" {
			return errors.New("unauthorized operation detected")
		}
		// search for results that start with current user id
		results, err := ctl.storage.List(authId, ctl.model)
		if err != nil {
			return api.Error(c, err)
		} else if results == nil || len(results) == 0 {
			//no data found
			return api.Success(c, str.UnsafeBytes(ctl.name), nil)
		} else {
			return api.SendSuccess(c, str.UnsafeBytes(ctl.name), results)
		}
	}
	return api.ErrorStr(c, data.NotAllowedToList)
}

// todo delegate rather than recall
func (ctl DatabaseController) Model() mixed.DatabaseObjectInterface {
	return ctl.model
}

// automatically register database related basic CRUD operations
func (ctl *DatabaseController) RegisterDatabaseMethods(router *echo.Group) {
	listPostPath := ctl.pathPrepend + "/" + ctl.name
	logger.Info("exposing GET ", listPostPath)
	router.GET(listPostPath, ctl.ListOwnerOnly)
	logger.Info("exposing POST ", listPostPath)
	router.POST(listPostPath, ctl.Create)

	customItemPath := ctl.pathPrepend + "/" + ctl.name + "/:id"
	logger.Info("exposing GET ", customItemPath)
	router.GET(customItemPath, ctl.Read)
	logger.Info("exposing PUT ", customItemPath)
	router.PUT(customItemPath, ctl.Update)
	logger.Info("exposing DELETE ", customItemPath)
	router.DELETE(customItemPath, ctl.Delete)
}

// implemented method from interface RouterRegistrable
func (ctl DatabaseController) RegisterRouters(router *echo.Group) {
	logger.Info("exposing ", ctl.name, " controller methods")
	ctl.RegisterDatabaseMethods(router)
}

// build the composite id: authId + modelId
func (ctl *DatabaseController) buildCompositeId(authId string, modelId string) []byte {
	c := authId + "." + modelId
	key := str.UnsafeBytes(c)
	return key
}
