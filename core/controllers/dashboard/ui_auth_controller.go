// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package dashboard

import (
	"github.com/zerjioang/etherniti/core/controllers/wrap"
	"github.com/zerjioang/etherniti/core/mail"
	"github.com/zerjioang/etherniti/shared"
	"github.com/zerjioang/etherniti/shared/notifier"
	"github.com/zerjioang/go-hpc/lib/db/badgerdb"
	"github.com/zerjioang/go-hpc/lib/mailer/fakesender"
	"github.com/zerjioang/go-hpc/lib/mailer/model"
	"github.com/zerjioang/go-hpc/lib/uuid/randomuuid"
	"github.com/zerjioang/go-hpc/thirdparty/jwt-go"

	"github.com/zerjioang/etherniti/core/config"
	"github.com/zerjioang/go-hpc/lib/checkmail"
	"github.com/zerjioang/go-hpc/lib/radix"

	"github.com/zerjioang/etherniti/core/api"
	"github.com/zerjioang/etherniti/core/controllers/common"
	"github.com/zerjioang/etherniti/core/data"
	"github.com/zerjioang/etherniti/core/logger"
	"github.com/zerjioang/etherniti/core/model/auth"
	"github.com/zerjioang/etherniti/shared/constants"
	"github.com/zerjioang/go-hpc/thirdparty/echo"
)

const (
	MinPasswordLen = 6
)

var (
	// Create the JWT key used to create the signature
	authTokenSecret = []byte("cc03a2bc-4a01-43dd-bdfe-a65f4a6e1f2f ")
	// radix tree of common passwords used
	rdx *radix.Tree
)

// todo remove this global
var (
	// load etherniti proxy configuration
	opts = config.GetDefaultOpts()
)

type UIAuthController struct {
	common.DatabaseController
}

func init() {
	logger.Debug("loading common password database into memory")
	rdx = radix.New()
	rdx.LoadFromRaw(config.BlacklistedPasswordFile, constants.NewLine)
	logger.Debug("blacklisted database loaded")
}

// constructor like function
func NewUIAuthController() *UIAuthController {
	uiCtl := &UIAuthController{}
	var err error
	uiCtl.DatabaseController, err = common.NewDatabaseController("", "", "auth", auth.NewDBAuthModel)
	if err != nil {
		logger.Error("failed to create authentication controller ", err)
	}
	return uiCtl
}

// logins user data and returns access token
func (ctl *UIAuthController) login(c *shared.EthernitiContext) error {

	//new login request
	req := auth.NewEmptyAuthRequest()
	if err := c.Bind(&req); err != nil {
		// return a binding error
		logger.Error(data.FailedToBind, err)
		return api.ErrorBytes(c, data.BindErr)
	}
	if req.Email != "" && req.Password != "" {
		logger.Info("logging user with email: ", req.Email)
		item, readErr := ctl.GetKey([]byte(req.Email))
		if readErr == nil {
			dto := auth.NewEmptyAuthRequest()
			pErr := badgerdb.Unserialize(item, &dto)
			if pErr != nil {
				logger.Error("failed to unserialize data: ", pErr.Error())
				return api.ErrorBytes(c, data.DatabaseError)
			}
			// check if email and password matches
			matches := req.Email == dto.Email && badgerdb.CompareHash(req.Password, dto.Password)
			if matches {
				// 1 if account state is unknown send a confirmation link to email
				// 2 if confirmation is pending show error
				// if account confirmed, create authentication token
				switch req.Status {
				case auth.AccountUnknown:
					go ctl.sendConfirmationLink(req)
					return api.ErrorBytes(c, data.LoginRequestAccountUnknownErr)
				case auth.AccountEmailConfirmationPending:
					return api.ErrorBytes(c, data.LoginRequestAccountConfirmationPendingErr)
				case auth.AccountEmailConfirmed:
					token, err := createToken(dto.Uuid)
					if err != nil || token == "" {
						logger.Error("failed to create authentication token: ", err)
						return api.ErrorBytes(c, data.InvalidLoginData)
					}
					return api.SendSuccess(c, data.UserLogin, auth.NewLoginResponse(token))
				default:
					// do not login. account might be blocked, under investigation or in recovery
				}
			}
			return api.ErrorBytes(c, data.InvalidLoginData)
		}
		//db-badger read error
		// this code is trigger each time user fails a login attempt
		return api.ErrorBytes(c, data.FailedLoginVerification)
	}
	return api.ErrorBytes(c, data.MissingLoginFields)
}

// registers new user data in the api sever
func (ctl *UIAuthController) register(c *shared.EthernitiContext) error {
	//new login request
	req := auth.NewEmptyAuthRequest()
	if err := c.Bind(&req); err != nil {
		// return a binding error
		logger.Error(data.FailedToBind, err)
		return api.ErrorBytes(c, data.BindErr)
	}
	if req.Email != "" && req.Password != "" && req.Username != "" {
		// 1 check password length
		if len(req.Password) < MinPasswordLen {
			logger.Error("etherniti proxy minimum password policy forces to use more characters than provided password")
			return api.ErrorStr(c, "Etherniti proxy minimum password policy forces to use more characters than provided password")
		}
		// 2 check password against common database
		_, found := rdx.Get(req.Password)
		if found {
			logger.Error("etherniti wont allow account registration with provided password")
			return api.ErrorStr(c, "Etherniti wont allow account registration with provided password because it is considered as a weak password. Please use stronger password combination to use Etherniti platform in order to have secure credentials.")
		}

		// 3 check validate email
		if !checkmail.FastEmailCheck(req.Email) {
			logger.Error("invalid email provided in registration")
			return api.ErrorStr(c, "invalid email provided in registration")
		}

		logger.Info("registering user with email: ", req.Email)
		// hash user password
		req.Password = badgerdb.Hash(req.Password)
		req.Uuid = randomuuid.GenerateIDString().UnsafeString()
		req.Role = constants.StandardUser
		code, genErr := req.GenConfirmationCode()
		if genErr != nil {
			logger.Error("failed to create email verification code")
			return api.ErrorBytes(c, data.UserRegisterFailed)
		}
		req.Confirmation = code
		saveErr := ctl.SetUniqueKey(req.PrimaryKey(), badgerdb.Serialize(req))
		if saveErr != nil {
			logger.Error("failed to register new user due to: ", saveErr)
			return api.ErrorBytes(c, data.UserRegisterFailed)
		}
		go func() {
			// send email confirmation link
			logger.Debug("sending user account activation email as background job")

			err := mail.SendActivationEmail(
				&model.AuthMailRequest{
					Username:     req.Username,
					Email:        req.Email,
					Confirmation: req.Confirmation,
				},
				mail.ConfirmEmailTemplate,
				fakesender.FakeEmailSender,
			)
			if err != nil {
				logger.Error("failed to send email confirmation due to: ", err)
			}
		}()
		// send new account created internal event
		notifier.NewDashboardAccountEvent.Emit()
		return api.SendSuccess(c, data.RegistrationSuccess, nil)
	}
	return api.ErrorStr(c, "registration aborted due to missing fields")
}

// account creation confirmation link
func (ctl *UIAuthController) confirm(c *shared.EthernitiContext) error {
	//new account confirmation request
	req := auth.NewEmptyAuthRequest()
	req.Confirmation = c.Param("msg")
	accountEmailId, err := req.IsValidConfirmation()
	if err != nil {
		// return a binding error
		logger.Error(data.FailedToBind, err)
		return api.Error(c, err)
	}
	logger.Error("confirming user account via confirmation link")
	dbkey := []byte(accountEmailId)
	item, readErr := ctl.GetKey(dbkey)
	if readErr == nil {
		pErr := badgerdb.Unserialize(item, &req)
		if pErr != nil {
			logger.Error("failed to unserialize data: ", pErr)
			return api.ErrorBytes(c, data.AccountConfirmDbError)
		}
		// check if account status is pending confirmation
		if !(req.Status == auth.AccountEmailConfirmationPending || req.Status == auth.AccountUnknown) {
			logger.Error("unauthorized account update due to invalid state. The account must be [Unknown, PendincConfirmation]")
			// show 'expiration like'' message to the users
			return api.ErrorBytes(c, data.AccountConfirmDeniedError)
		}
		// update current user status to account confirmed
		req.Status = auth.AccountEmailConfirmed
		// store account updated
		updateErr := ctl.UpdateKey(dbkey, badgerdb.Serialize(req))
		if updateErr != nil {
			logger.Error("failed to update user authentication data: ", updateErr)
			return api.ErrorBytes(c, data.AccountConfirmUpdateError)
		}
	}
	// redirect to dashboard
	return api.Redirect(c, opts.Authentication.ConfirmationRedirectUrl)
}

// generate a token for given user.
// this functions allows to firebase registered users to work with the proxy
func (ctl *UIAuthController) token(c *shared.EthernitiContext) error {
	//new login request
	req := auth.NewEmptyAuthRequest()
	if err := c.Bind(&req); err != nil {
		// return a binding error
		logger.Error(data.FailedToBind, err)
		return api.ErrorBytes(c, data.BindErr)
	}
	logger.Error("failed to generate user token")
	return api.ErrorBytes(c, data.UserTokenFailed)
}

// triggers user account recovery mecanisms
func (ctl *UIAuthController) recover(c *shared.EthernitiContext) error {
	//new recovery request
	req := auth.NewEmptyAuthRequest()
	if err := c.Bind(&req); err != nil {
		// return a binding error
		logger.Error(data.FailedToBind, err)
		return api.ErrorBytes(c, data.BindErr)
	}
	if req.Email != "" {
		logger.Info("recovering user with email: ", req.Email)
		return api.SendSuccess(c, []byte("account recovery in progress"), nil)
	}
	return api.ErrorStr(c, "recovery aborted due to missing fields")
}

// validates recatpcha requests
// more info at: https://www.google.com/recaptcha/admin/site/346227166
func (ctl *UIAuthController) recaptcha(c *shared.EthernitiContext) error {
	return data.ErrNotImplemented
}

func (ctl *UIAuthController) RegisterRouters(router *echo.Group) {
	logger.Debug("exposing ui controller methods")
	router.POST("/auth/login", wrap.Call(ctl.login))
	router.POST("/auth/register", wrap.Call(ctl.register))
	router.GET("/auth/confirm/:msg", wrap.Call(ctl.confirm))
	router.POST("/auth/recover", wrap.Call(ctl.recover))
	router.POST("/auth/token", wrap.Call(ctl.token))
	router.POST("/auth/recaptcha", wrap.Call(ctl.recaptcha))
}

func (ctl *UIAuthController) sendConfirmationLink(req auth.AuthRequest) {
	// todo send email to the user
}

func ParseAuthenticationToken(tokenStr string) (auth.AuthRequest, error) {
	var decoded auth.AuthRequest
	// Parse takes the token string and a function for looking up the key. The latter is especially
	// useful if you use multiple keys for your application.  The standard is to use 'kid' in the
	// head of the token to identify which key to use, but the parsed token (head and claims) is provided
	// to the callback, providing flexibility.
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, data.ErrInvalidSigningMethod
		}
		// return used token secret
		return authTokenSecret, nil
	})
	if err != nil {
		return decoded, err
	}

	mapc, ok := token.Claims.(jwt.MapClaims)
	if !ok || mapc == nil {
		return decoded, data.ErrFailedToRead
	}
	decoded.Uuid = mapc["sid"].(string)
	return decoded, nil
}
