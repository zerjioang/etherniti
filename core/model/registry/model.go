package registry

import (
	"encoding/json"

	"github.com/zerjioang/go-hpc/lib/uuid/randomuuid"

	"github.com/zerjioang/etherniti/shared"

	"github.com/zerjioang/go-hpc/thirdparty/echo/protocol/encoding"

	"github.com/zerjioang/go-hpc/shared/db"

	"github.com/zerjioang/etherniti/core/model"
	"github.com/zerjioang/go-hpc/common"

	"github.com/zerjioang/etherniti/core/data"
	"github.com/zerjioang/etherniti/core/logger"
	"github.com/zerjioang/etherniti/core/model/metadata"
	"github.com/zerjioang/etherniti/core/model/registry/version"
	"github.com/zerjioang/go-hpc/lib/eth"
	"github.com/zerjioang/go-hpc/lib/stack"
	"github.com/zerjioang/go-hpc/util/str"
)

type Registry struct {
	// implement interface to be a rest-db-badger-crud able struct
	db.DaoInterface `json:"_,omitempty"`

	// unique registry identifier used for database storage
	Uuid [8]byte `json:"sid"`
	// registry name
	Name string `json:"name"`
	// registry description
	Description string `json:"description"`

	// contract address for this version
	Address string `json:"address"`
	// contract version
	Version version.ContractVersion `json:"version"`
	// contract metadata
	Metadata *metadata.Metadata `json:"metadata,omitempty"`
}

func (r Registry) Id() string {
	return r.Name + "-" + r.Version.String()
}

func (r Registry) Validate() stack.Error {
	if r.Name == "" {
		return stack.New("you must supply a valid contract name")
	}
	if !eth.IsValidAddressLow(r.Address) {
		return stack.New("you must supply a valid contract address starting with 0x")
	}
	if !r.Version.Valid() {
		return stack.New("you must supply a valid contract version")
	}
	return stack.Nil()
}

// implementation of interface DaoInterface
func (r Registry) Key() []byte {
	return str.UnsafeBytes(r.Id())
}
func (r Registry) Value(serializer common.Serializer) []byte {
	return encoding.GetBytesFromSerializer(serializer, r)
}
func (r Registry) NewDao() db.DaoInterface {
	return NewEmptyRegistryDao()
}

// custom validation logic for read operation
// return nil if everyone can read
func (r Registry) CanRead(context *shared.EthernitiContext, key string) error {
	// todo check if current r id belongs to current user
	return nil
}

// custom validation logic for update operation
// return nil if everyone can update
func (r Registry) CanUpdate(context *shared.EthernitiContext, key string) error {
	// todo check if current r id belongs to current user
	return nil
}

// custom validation logic for delete operation
// return nil if everyone can delete
func (r Registry) CanDelete(context *shared.EthernitiContext, key string) error {
	// todo check if current r id belongs to current user
	return nil
}

// custom validation logic for write operation
// return nil if everyone can write
func (r Registry) CanWrite(context *shared.EthernitiContext) error {
	return nil
}

// custom validation logic for list operation
// return nil if everyone can list
func (r Registry) CanList(context *shared.EthernitiContext) error {
	// todo check if current r id belongs to current user
	return nil
}

func (r Registry) Bind(context *shared.EthernitiContext) (db.DaoInterface, stack.Error) {
	//new registry creation request
	if err := context.Bind(&r); err != nil {
		// return a binding error
		logger.Error(data.FailedToBind, err)
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
		projectOwner := context.UserId()

		if intIP == 0 || projectOwner == "" {
			logger.Error("failed to create new registry: missing data")
			return nil, data.ErrInvalidData
		} else {
			r.Metadata = metadata.NewMetadata(context)
			return r, stack.Nil()
		}
	}
}

func (r Registry) Update(o db.DaoInterface) (db.DaoInterface, stack.Error) {
	newRg, ok := o.(*Registry)
	if newRg == nil || !ok {
		return nil, model.UnsupportedDataErr
	}
	//if new name is provided, update it
	r.Name = model.ConditionalStringUpdate(newRg.Name, r.Name, "")
	return r, stack.Nil()
}

func NewEmptyRegistryDao() db.DaoInterface {
	return &Registry{}
}

// converts byte sequence to go registry struct
func (r Registry) Decode(data []byte) (db.DaoInterface, stack.Error) {
	o := NewEmptyRegistryDao()
	err := json.Unmarshal(data, &o)
	return o, stack.Ret(err)
}

func NewRegistry(name string, description string, major int, minor int, mtdt *metadata.Metadata) *Registry {
	p := new(Registry)
	p.Uuid = randomuuid.GenerateSnowFlakeId()
	p.Name = name
	p.Description = description
	p.Version.Major = major
	p.Version.Minor = minor
	p.Metadata = mtdt
	return p
}
