package common

import (
	"errors"
	"strings"

	"github.com/zerjioang/etherniti/core/controllers/wrap"
	"github.com/zerjioang/go-hpc/lib/stack"

	"github.com/zerjioang/go-hpc/thirdparty/echo/protocol"
	"github.com/zerjioang/go-hpc/thirdparty/echo/protocol/encoding"

	"github.com/zerjioang/go-hpc/lib/db/badgerdb"
	"github.com/zerjioang/go-hpc/shared/db"

	"github.com/zerjioang/etherniti/core/api"
	"github.com/zerjioang/etherniti/core/data"
	"github.com/zerjioang/etherniti/core/logger"
	"github.com/zerjioang/etherniti/shared"
	"github.com/zerjioang/go-hpc/thirdparty/echo"
	"github.com/zerjioang/go-hpc/util/str"
)

var (
	errInvalidCollectionName = errors.New("invalid collection name provided")
	errUnauthorizedOp        = errors.New("unauthorized operation detected")
	s                        = []byte(".")
	dbSerializer, _          = encoding.EncodingModeSelector(protocol.ModeJson)
)

type DatabaseController struct {
	storage *badgerdb.BadgerStorage
	name    string
	// data model generator function. it will create a new struct
	modelGenerator func() db.DaoInterface
	// data model used in this collection
	model db.DaoInterface
	// path prepend string
	pathPrepend string
}

// constructor like function
func NewDatabaseController(pathPrepend, dbPath, collection string, modelGenerator func() db.DaoInterface) (DatabaseController, error) {
	dbctl := DatabaseController{}
	if collection == "" {
		return dbctl, errInvalidCollectionName
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
	dbctl.storage, err = badgerdb.NewCollection(dbPath, collection)
	if err != nil {
		logger.Error("failed to initialize database db-badger collection: ", err)
	}
	return dbctl, err
}

func (ctl *DatabaseController) Create(c *shared.EthernitiContext) error {
	daoItem := ctl.modelGenerator()
	daoItem.Decode(c.Body())
	// convert dao bject to dao-with-api object
	requestedItem, err := ctl.toRequestItem(daoItem)
	if err.Occur() {
		return api.StackError(c, err)
	}

	if requestedItem != nil {
		canWriteErr := requestedItem.CanWrite(c)
		if canWriteErr == nil {
			// check if current user Id is valid, exists.
			// source: auth-jwt-token
			authId := c.AuthUserUuid()
			if authId == "" {
				return errUnauthorizedOp
			}
			key := ctl.buildCompositeId(authId, string(daoItem.Key()))
			value := daoItem.Value(dbSerializer)
			writeErr := ctl.storage.PutUniqueKeyValue(key, value)
			if writeErr != nil {
				return api.Error(c, writeErr)
			} else {
				return api.SendSuccess(c, data.SuccessfullyCreated, requestedItem)
			}
		} else {
			return api.Error(c, canWriteErr)
		}
	}
	return api.ErrorBytes(c, data.FailedToProcess)
}

func (ctl *DatabaseController) GetKey(key []byte) ([]byte, error) {
	return ctl.storage.Get(key)
}

func (ctl *DatabaseController) SetKey(key []byte, value []byte) error {
	return ctl.storage.SetRawKey(key, value)
}

func (ctl *DatabaseController) UpdateKey(key []byte, value []byte) error {
	return ctl.storage.SetRawKey(key, value)
}

func (ctl *DatabaseController) SetUniqueKey(key []byte, value []byte) error {
	return ctl.storage.PutUniqueKeyValue(key, value)
}

func (ctl *DatabaseController) Update(c *shared.EthernitiContext) error {
	modelId := c.Param("id")
	if modelId != "" {
		daoItem := ctl.modelGenerator()
		// convert dao bject to dao-with-api object
		requestedItem, err := ctl.toRequestItem(daoItem)
		if err.Occur() {
			return api.StackError(c, err)
		}
		canUpdateErr := requestedItem.CanUpdate(c, modelId)
		if canUpdateErr == nil {
			// check if current user Id is valid, exists.
			// source: auth-jwt-token
			authId := c.AuthUserUuid()
			if authId == "" {
				return errUnauthorizedOp
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
			requestedItem, bindErr := daoItem.Decode(c.Body())
			if bindErr.Occur() {
				return api.StackError(c, bindErr)
			}
			// update item
			updatedItem, updateErr := readedObject.Update(requestedItem)
			if updateErr.Occur() {
				return api.StackError(c, updateErr)
			}
			// save result
			storeErr := ctl.storage.SetRawKey(key, updatedItem.Value(dbSerializer))
			if storeErr != nil {
				return api.Error(c, storeErr)
			}
			return api.SendSuccessBlob(c, updatedItem.Value(dbSerializer))
		}
		return api.Error(c, canUpdateErr)
	} else {
		return api.ErrorBytes(c, data.ProvideId)
	}
}

func (ctl *DatabaseController) Read(c *shared.EthernitiContext) error {
	modelId := c.Param("id")
	if modelId != "" {
		daoItem := ctl.modelGenerator()
		// convert dao bject to dao-with-api object
		requestedItem, err := ctl.toRequestItem(daoItem)
		if err.Occur() {
			return api.StackError(c, err)
		}
		canReadErr := requestedItem.CanRead(c, modelId)
		if canReadErr == nil {
			// check if current user Id is valid, exists.
			// source: auth-jwt-token
			authId := c.AuthUserUuid()
			if authId == "" {
				return errUnauthorizedOp
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
		return api.ErrorBytes(c, data.ProvideId)
	}
}

func (ctl *DatabaseController) Delete(c *shared.EthernitiContext) error {
	modelId := c.Param("id")
	if modelId != "" {
		daoItem := ctl.modelGenerator()
		// convert dao bject to dao-with-api object
		requestedItem, err := ctl.toRequestItem(daoItem)
		if err.Occur() {
			return api.StackError(c, err)
		}
		canDeleteErr := requestedItem.CanDelete(c, modelId)
		if canDeleteErr == nil {
			// check if current user Id is valid, exists.
			// source: auth-jwt-token
			authId := c.AuthUserUuid()
			if authId == "" {
				return errUnauthorizedOp
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
		return api.ErrorBytes(c, data.ProvideId)
	}
}

func (ctl *DatabaseController) List(c *shared.EthernitiContext) error {
	daoItem := ctl.modelGenerator()
	// convert dao bject to dao-with-api object
	requestedItem, err := ctl.toRequestItem(daoItem)
	if err.Occur() {
		return api.StackError(c, err)
	}
	canListErr := requestedItem.CanList(c)
	if canListErr == nil {
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
	return api.ErrorBytes(c, data.NotAllowedToList)
}

func (ctl *DatabaseController) ListOwnerOnly(c *shared.EthernitiContext) error {
	daoItem := ctl.modelGenerator()
	// convert dao bject to dao-with-api object
	requestedItem, err := ctl.toRequestItem(daoItem)
	if err.Occur() {
		return api.StackError(c, err)
	}
	canListErr := requestedItem.CanList(c)
	if canListErr == nil {
		// check if current user Id is valid, exists.
		// source: auth-jwt-token
		authId := c.AuthUserUuid()
		if authId == "" {
			return errUnauthorizedOp
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
	return api.ErrorBytes(c, data.NotAllowedToList)
}

// todo delegate rather than recall
func (ctl DatabaseController) Model() db.DaoInterface {
	return ctl.model
}

// automatically register database related basic CRUD operations
func (ctl *DatabaseController) RegisterDatabaseMethods(router *echo.Group) {
	listPostPath := ctl.pathPrepend + "/" + ctl.name
	logger.Info("exposing GET ", listPostPath)
	router.GET(listPostPath, wrap.Call(ctl.ListOwnerOnly))
	logger.Info("exposing POST ", listPostPath)
	router.POST(listPostPath, wrap.Call(ctl.Create))

	customItemPath := ctl.pathPrepend + "/" + ctl.name + "/:id"
	logger.Info("exposing GET ", customItemPath)
	router.GET(customItemPath, wrap.Call(ctl.Read))
	logger.Info("exposing PUT ", customItemPath)
	router.PUT(customItemPath, wrap.Call(ctl.Update))
	logger.Info("exposing DELETE ", customItemPath)
	router.DELETE(customItemPath, wrap.Call(ctl.Delete))
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

// build the composite id: authId + modelId
func (ctl *DatabaseController) buildCompositeId2(authId string, modelId string) []byte {
	return []byte(strings.Join([]string{authId, modelId}, "."))
}

// build the composite id: authId + modelId
func (ctl *DatabaseController) buildCompositeId3(authId string, modelId string) []byte {
	var b strings.Builder
	b.WriteString(authId)
	b.WriteString(".")
	b.WriteString(modelId)
	return []byte(b.String())
}

// build the composite id: authId + modelId
func (ctl *DatabaseController) buildCompositeId4(authId string, modelId string) db.ItemKey {
	return db.ItemKey{
		Left:  []byte(authId),
		Sep:   []byte("."),
		Right: []byte(modelId),
	}
}

// build the composite id: authId + modelId
func (ctl *DatabaseController) buildCompositeId5(authId string, modelId string) db.ItemKey {
	return db.ItemKey{
		Left:  str.UnsafeBytes(authId),
		Sep:   s,
		Right: str.UnsafeBytes(modelId),
	}
}

// build the composite id: authId + modelId
func (ctl *DatabaseController) buildCompositeId6(authId string, modelId string) *db.ItemKey {
	return &db.ItemKey{
		Left:  str.UnsafeBytes(authId),
		Sep:   str.UnsafeBytes("."),
		Right: str.UnsafeBytes(modelId),
	}
}

// build the composite id: authId + modelId
func (ctl *DatabaseController) buildCompositeId7(authId []byte, modelId []byte) *db.ItemKey {
	return &db.ItemKey{
		Left:  authId,
		Sep:   s,
		Right: modelId,
	}
}

// build the composite id: authId + modelId
func (ctl *DatabaseController) buildCompositeId8(authId []byte, modelId []byte) db.ItemKey {
	return db.ItemKey{
		Left:  authId,
		Sep:   s,
		Right: modelId,
	}
}

func (ctl *DatabaseController) toRequestItem(item db.DaoInterface) (db.ControllerDBPolicyInterface, stack.Error) {
	if item == nil {
		return nil, stack.New("no source data found in the request")
	}
	policy, ok := item.(db.ControllerDBPolicyInterface)
	if ok && policy != nil {
		return policy, stack.Nil()
	}
	return nil, stack.New("forbidden data provided")
}
