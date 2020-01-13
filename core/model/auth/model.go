package auth

import (
	"errors"

	"github.com/zerjioang/etherniti/shared"

	"github.com/zerjioang/go-hpc/thirdparty/echo/protocol/encoding"

	"github.com/zerjioang/go-hpc/common"
	"github.com/zerjioang/go-hpc/lib/db/badgerdb"
	"github.com/zerjioang/go-hpc/shared/db"

	"github.com/zerjioang/go-hpc/lib/secure"
	"github.com/zerjioang/go-hpc/lib/secure/chacha20"
	"github.com/zerjioang/go-hpc/lib/stack"
	"github.com/zerjioang/go-hpc/util/hex"

	"github.com/zerjioang/etherniti/core/data"
	"github.com/zerjioang/etherniti/core/logger"
	"github.com/zerjioang/etherniti/shared/constants"
	"github.com/zerjioang/go-hpc/util/str"
)

var (
	confirmationEncoder *chacha20.ChachaEncoder
)

type AccountState uint8

const (
	AccountUnknown AccountState = iota
	AccountEmailConfirmationPending
	AccountEmailConfirmed
	AccountBlocked
	AccountUnderInvestigation
	AccountRecoveryRequested
)

// new login request dto
type AuthRequest struct {
	db.DaoInterface `json:"_,omitempty"`
	Uuid            string             `json:"sid,omitempty"`
	Username        string             `json:"name,omitempty" form:"name" query:"name"`
	Role            constants.UserRole `json:"role,omitempty" form:"role" query:"role"`
	Email           string             `json:"email" form:"email" query:"email"`
	Password        string             `json:"pwd" form:"pwd" query:"pwd"`
	// for api key based authentication
	ApiKey    string `json:"key,omitempty" form:"key" query:"key"`
	ApiSecret string `json:"secret,omitempty" form:"secret" query:"secret"`
	// for account confirmation
	Confirmation string `json:"confirmation,omitempty"`
	// for account state management
	Status AccountState `json:"status,omitempty"`
}

func init() {
	logger.Debug("creating confirmation link secret")
	pwd := secure.Keygen256()
	logger.Debug("creating confirmation link encoder")
	confirmationEncoder = chacha20.NewChachaEncoderParams([]byte(pwd))
}

// implementation of interface DaoInterface
func (req *AuthRequest) Key() []byte {
	return str.UnsafeBytes(str.ToLowerAscii(req.Email))
}
func (req *AuthRequest) Value(serializer common.Serializer) []byte {
	return encoding.GetBytesFromSerializer(serializer, req)
}

// custom validation logic for read operation
// return nil if everyone can read
func (req *AuthRequest) CanRead(context *shared.EthernitiContext, key string) error {
	return nil
}

// custom validation logic for update operation
// return nil if everyone can update
func (req *AuthRequest) CanUpdate(context *shared.EthernitiContext, key string) error {
	return data.ErrNotAuthorized
}

// custom validation logic for delete operation
// return nil if everyone can delete
func (req *AuthRequest) CanDelete(context *shared.EthernitiContext, key string) error {
	return nil
}

// custom validation logic for write operation
// return nil if everyone can write
func (req *AuthRequest) CanWrite(context *shared.EthernitiContext) error {
	if req.Email != "" && req.Password != "" && req.Username != "" {
		logger.Info("registering user with email: ", req.Email)
		// hash user password
		req.Password = badgerdb.Hash(req.Password)
		return nil
	}
	return errors.New("you have to provide a valid email, password and username")
}

// custom validation logic for list operation
// return nil if everyone can list
func (req *AuthRequest) CanList(context *shared.EthernitiContext) error {
	return data.ErrListingNotSupported
}

func (req *AuthRequest) Bind(context *shared.EthernitiContext) (db.DaoInterface, stack.Error) {
	if err := context.Bind(&req); err != nil {
		// return a binding error
		logger.Error(data.FailedToBind, err)
		return nil, stack.Ret(err)
	}
	return nil, data.ErrBind
}

// IsValidConfirmation detects whether given confirmation string is valid or not
func (req *AuthRequest) IsValidConfirmation() (string, error) {
	if req == nil {
		return "", errors.New("missing account confirmation payload")
	}
	if req.Confirmation == "" {
		return "", errors.New("missing account confirmation data")
	}
	// chacha20 nonce is always 48 hex encoded, so thats the minimum length allowed in the confirmation link
	if len(req.Confirmation) <= 48 {
		return "", errors.New("invalid or corrupted confirmation data")
	}
	// split nonce and encrypted content
	nonce := req.Confirmation[0:48]
	ciphertext := req.Confirmation[48:]
	nonceRaw, _ := hex.DecodeString(nonce)
	cipherRaw, _ := hex.DecodeString(ciphertext)
	plaintext, err := confirmationEncoder.Decrypt(cipherRaw, nonceRaw)
	if err != nil {
		logger.Error("failed to verify confirmation link due to: ", err)
		return "", errors.New("confirmation link is not valid or has expired")
	}
	if plaintext == nil || len(plaintext) == 0 {
		logger.Error("failed to verify confirmation link message due to: ", err)
		return "", errors.New("confirmation link message is not valid")
	}
	return string(plaintext), nil
}

// genConfirmationCode generates an email account confirmation code for current suer
func (req *AuthRequest) GenConfirmationCode() (string, error) {
	cipher, nonce, err := confirmationEncoder.Encrypt([]byte(req.Email))
	if err != nil {
		logger.Error("failed to generate account confirmation link code due to: ", err)
		return "", err
	}
	// return confirmation code as nonce:cipher
	code := hex.EncodeString(nonce) + hex.EncodeString(cipher)
	return code, nil
}

func (req *AuthRequest) PrimaryKey() []byte {
	return str.UnsafeBytes(req.Email)
}

func NewEmptyAuthRequestPtr() *AuthRequest {
	return new(AuthRequest)
}

func NewEmptyAuthRequest() AuthRequest {
	return AuthRequest{}
}

func NewDBAuthModel() db.DaoInterface {
	return NewEmptyAuthRequestPtr()
}
