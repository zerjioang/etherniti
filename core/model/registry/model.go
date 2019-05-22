package registry

import (
	"github.com/zerjioang/etherniti/core/data"
	"github.com/zerjioang/etherniti/core/logger"
	"github.com/zerjioang/etherniti/core/model/metadata"
	"github.com/zerjioang/etherniti/core/modules/stack"
	"github.com/zerjioang/etherniti/core/util/id"
	"github.com/zerjioang/etherniti/core/util/str"
	"github.com/zerjioang/etherniti/shared/mixed"
	"github.com/zerjioang/etherniti/thirdparty/echo"
)

type Registry struct {
	// implement interface to be a rest-db-crud able struct
	mixed.DatabaseObjectInterface `json:"_,omitempty"`

	// unique registry indetifier used for database storage
	Uuid string `json:"uuid"`
	// registry name
	Name string `json:"name"`
	// registry description
	Description string `json:"description"`

	Address string `json:"address"`
	Version string `json:"version"`

	Metadata *metadata.Metadata `json:"metadata"`
}

func (r Registry) Id() string {
	return r.Name + "-" + r.Version
}

func (r Registry) Validate() stack.Error {
	if r.Name == "" {
		return stack.New("you must supply a valid contract name")
	}
	if r.Address == "" {
		return stack.New("you must supply a valid contract address starting with 0x")
	}
	if r.Version == "" {
		return stack.New("you must supply a valid contract version")
	}
	return stack.Nil()
}

// implementation of interface DatabaseObjectInterface
func (r Registry) Key() []byte {
	return str.UnsafeBytes(r.Id())
}
func (r Registry) Value() []byte {
	return str.GetJsonBytes(r)
}
func (r Registry) New() mixed.DatabaseObjectInterface {
	return NewEmptyRegistry()
}

// custom validation logic for read operation
// return nil if everyone can read
func (r Registry) CanRead(context *echo.Context, key string) error {
	// todo check if current r id belongs to current user
	return nil
}

// custom validation logic for update operation
// return nil if everyone can update
func (r Registry) CanUpdate(context *echo.Context, key string) error {
	// todo check if current r id belongs to current user
	return nil
}

// custom validation logic for delete operation
// return nil if everyone can delete
func (r Registry) CanDelete(context *echo.Context, key string) error {
	// todo check if current r id belongs to current user
	return nil
}

// custom validation logic for write operation
// return nil if everyone can write
func (r Registry) CanWrite(context *echo.Context) error {
	return nil
}

// custom validation logic for list operation
// return nil if everyone can list
func (r Registry) CanList(context *echo.Context) error {
	// todo check if current r id belongs to current user
	return nil
}

func (r Registry) Bind(context *echo.Context) (mixed.DatabaseObjectInterface, stack.Error) {
	//new registry creation request
	if err := context.Bind(&r); err != nil {
		// return a binding error
		logger.Error("failed to bind request data to model: ", err)
		return nil, data.ErrBind
	}
	e := r.Validate()
	if e.Occur() {
		logger.Error("failed to validate new registry object data: ", e.Error())
		return nil, e
	} else {
		// get required data to build a new registry item
		intIP := context.IntIp()
		// get user uuid
		projectOwner := context.AuthenticatedUserUuid()

		if intIP == 0 || projectOwner == "" {
			logger.Error("failed to create new registry: missing data")
			return nil, data.StackErrInvalidData
		} else {
			return r, stack.Nil()
		}
	}
}

func NewEmptyRegistry() Registry {
	return Registry{}
}

func NewDBRegistry() mixed.DatabaseObjectInterface {
	return NewEmptyRegistry()
}

func NewRegistry(ctx *echo.Context, name string, description string) *Registry {
	p := new(Registry)
	p.Uuid = id.GenerateIDString().String()
	p.Name = name
	p.Description = description
	p.Metadata = metadata.NewMetadata(ctx)
	return p
}
