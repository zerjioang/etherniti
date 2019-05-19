package common

import (
	"errors"

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
	// data model used in this collection
	model mixed.DatabaseObjectInterface
}

// constructor like function
func NewDatabaseController(collection string, model mixed.DatabaseObjectInterface) (DatabaseController, error) {
	dbctl := DatabaseController{}
	dbctl.name = collection
	dbctl.model = model
	if collection == "" {
		return dbctl, errors.New("invalid collection name provided")
	}
	// database initialization method
	var err error
	dbctl.storage, err = db.NewCollection(collection)
	if err != nil {
		logger.Error("failed to initialize projects db collection: ", err)
	}
	return dbctl, err
}

func (ctl *DatabaseController) Create(c *echo.Context) error {
	requestedItem, err := ctl.model.New().Bind(c)
	if err.Occur() {
		return api.StackError(c, err)
	}
	if requestedItem != nil {
		canWriteErr := requestedItem.CanWrite(c)
		if canWriteErr == nil {
			requestedItem.Key()
			writeErr := ctl.storage.PutUniqueKeyValue(requestedItem.Key(), requestedItem.Value())
			if writeErr != nil {
				return api.Error(c, writeErr)
			} else {
				return api.SendSuccess(c, data.SuccessfullyCreated, requestedItem)
			}
		} else {
			return api.Error(c, canWriteErr)
		}
	}
	return api.ErrorStr(c, data.FailedToProcess)
}

func (ctl *DatabaseController) Read(c *echo.Context) error {
	projectId := c.Param("id")
	if projectId != "" {
		canReadErr := ctl.model.CanRead(c, projectId)
		if canReadErr == nil {
			key := str.UnsafeBytes(projectId)
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
	projectId := c.Param("id")
	if projectId != "" {
		canUpdateErr := ctl.model.CanUpdate(c, projectId)
		if canUpdateErr == nil {
			key := str.UnsafeBytes(projectId)
			projectData, readErr := ctl.storage.Get(key)
			if readErr != nil {
				return api.Error(c, readErr)
			}
			return api.SendSuccessBlob(c, projectData)
		}
		return api.Error(c, canUpdateErr)
	} else {
		return api.ErrorStr(c, data.ProvideId)
	}
}

func (ctl *DatabaseController) Delete(c *echo.Context) error {
	projectId := c.Param("id")
	if projectId != "" {
		canDeleteErr := ctl.model.CanDelete(c, projectId)
		if canDeleteErr == nil {
			key := str.UnsafeBytes(projectId)
			deleteErr := ctl.storage.Delete(key)
			if deleteErr != nil {
				return api.Error(c, deleteErr)
			}
			return api.SendSuccessBlob(c, data.SuccessfullyDeleted)
		}
		return api.Error(c, canDeleteErr)
	} else {
		return api.ErrorStr(c, data.ProvideId)
	}
}

func (ctl *DatabaseController) List(c *echo.Context) error {
	canList := ctl.model.CanList(c)
	if canList == nil {
		results, err := ctl.storage.List("")
		if err != nil {
			return api.Error(c, err)
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
	listPostPath := "/" + ctl.name
	router.GET(listPostPath, ctl.List)
	router.POST(listPostPath, ctl.Create)

	customItemPath := "/" + ctl.name + "/:id"
	router.GET(customItemPath, ctl.Read)
	router.PUT(customItemPath, ctl.Update)
	router.DELETE(customItemPath, ctl.Delete)
}

// implemented method from interface RouterRegistrable
func (ctl DatabaseController) RegisterRouters(router *echo.Group) {
	logger.Info("exposing project controller methods")
	ctl.RegisterDatabaseMethods(router)
}
