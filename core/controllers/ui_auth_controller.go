// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package controllers

import (
	"time"

	"github.com/zerjioang/etherniti/core/api"
	"github.com/zerjioang/etherniti/core/controllers/common"
	"github.com/zerjioang/etherniti/core/data"
	"github.com/zerjioang/etherniti/core/db"
	"github.com/zerjioang/etherniti/core/logger"
	"github.com/zerjioang/etherniti/core/model/auth"
	"github.com/zerjioang/etherniti/core/util/id"
	"github.com/zerjioang/etherniti/core/util/str"
	"github.com/zerjioang/etherniti/shared/constants"
	"github.com/zerjioang/etherniti/thirdparty/echo"
	"github.com/zerjioang/etherniti/thirdparty/jwt-go"
)

var (
	// Create the JWT key used to create the signature
	authTokenSecret = []byte(" cc03a2bc-4a01-43dd-bdfe-a65f4a6e1f2f ")
)

type UIAuthController struct {
	common.DatabaseController
}

// constructor like function
func NewUIAuthController() UIAuthController {
	uiCtl := UIAuthController{}
	var err error
	uiCtl.DatabaseController, err = common.NewDatabaseController("auth", auth.NewDBAuthModel())
	if err != nil {
		logger.Error("failed to create database controller ", err)
	}
	return uiCtl
}

// logins user data and returns access token
func (ctl UIAuthController) login(c *echo.Context) error {

	//new login request
	req := auth.NewEmptyAuthRequest()
	if err := c.Bind(&req); err != nil {
		// return a binding error
		logger.Error("failed to bind request data to model: ", err)
		return api.ErrorStr(c, data.BindErr)
	}
	if req.Email != "" && req.Password != "" {
		logger.Info("logging user with email: ", req.Email)
		item, readErr := ctl.GetKey([]byte(req.Email))
		if readErr == nil {
			dto := auth.NewEmptyAuthRequest()
			pErr := db.Unserialize(item, &dto)
			if pErr != nil {
				logger.Error("failed to unserialize data: ", pErr.Error())
				return api.ErrorStr(c, data.DatabaseError)
			} else {
				// check if email and password matches
				matches := req.Email == dto.Email && db.CompareHash(req.Password, dto.Password)
				if matches {
					// create authentication token
					token, err := ctl.createToken(dto.Uuid)
					if err != nil || token == "" {
						logger.Error("failed to create authentication token: ", err)
						return api.ErrorStr(c, data.InvalidLoginData)
					} else {
						return api.Success(c, data.UserLogin, auth.NewLoginResponse(token).Json())
					}
				} else {
					return api.ErrorStr(c, data.InvalidLoginData)
				}
			}
		} else {
			//db read error
			// this code is trigger each time user fails a login attempt
			return api.ErrorStr(c, data.FailedLoginVerification)
		}
	} else {
		return api.ErrorStr(c, data.MissingLoginFields)
	}
}

// registers new user data in the api sever
func (ctl UIAuthController) register(c *echo.Context) error {
	//new login request
	req := auth.NewEmptyAuthRequest()
	if err := c.Bind(&req); err != nil {
		// return a binding error
		logger.Error("failed to bind request data to model: ", err)
		return api.ErrorStr(c, data.BindErr)
	}
	if req.Email != "" && req.Password != "" && req.Username != "" {
		logger.Info("registering user with email: ", req.Email)
		// hash user password
		req.Password = db.Hash(req.Password)
		req.Uuid = id.GenerateIDString().UnsafeString()
		req.Role = constants.StandardUser
		saveErr := ctl.SetUniqueKey(str.UnsafeBytes(req.Email), db.Serialize(req))
		if saveErr != nil {
			logger.Error("failed to register new user due to: ", saveErr)
			return api.ErrorStr(c, data.UserRegisterFailed)
		} else {
			return api.Success(c, data.RegistrationSuccess, nil)
		}
	}
	return nil
}

// triggers user account recovery mecanisms
func (ctl UIAuthController) recover(c *echo.Context) error {
	//new recovery request
	req := auth.NewEmptyAuthRequest()
	if err := c.Bind(&req); err != nil {
		// return a binding error
		logger.Error("failed to bind request data to model: ", err)
		return api.ErrorStr(c, data.BindErr)
	}
	if req.Email != "" {
		logger.Info("recovering user with email: ", req.Email)
		return api.Success(c, []byte("account recovery in progress"), nil)
	}
	return nil
}

// validates recatpcha requests
// more info at: https://www.google.com/recaptcha/admin/site/346227166
func (ctl UIAuthController) recaptcha(c *echo.Context) error {
	return data.ErrNotImplemented
}

func (ctl UIAuthController) RegisterRouters(router *echo.Group) {
	logger.Debug("exposing ui controller methods")
	router.POST("/auth/login", ctl.login)
	router.POST("/auth/register", ctl.register)
	router.POST("/auth/recover", ctl.recover)
	router.POST("/auth/recaptcha", ctl.recaptcha)
}
func (ctl UIAuthController) createToken(userUuid string) (string, error) {
	type Claims struct {
		User string `json:"u"`
		jwt.StandardClaims
	}
	// Declare the expiration time of the token
	// here, we have kept it as 20 minutes
	expirationTime := time.Now().Add(20 * time.Minute)
	claims := &Claims{
		User: userUuid,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign and get the complete encoded token as a string using the secret
	return token.SignedString(authTokenSecret)
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
	decoded.Uuid = mapc["u"].(string)
	return decoded, nil
}
