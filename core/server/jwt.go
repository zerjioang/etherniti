// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package server

import (
	"net/http"

	"github.com/labstack/echo"
)

const (
	jwtCookieName = "auth"
	jwtHeader     = "auth"
	jwtParam      = "auth"
	authScheme    = "sha256"
)

var (
	ErrJWTMissing = echo.NewHTTPError(http.StatusBadRequest, "missing or malformed jwt")
)

// jwtFromHeader returns a `jwtExtractor` that extracts token from the request header.
func jwtFromHeader(c echo.Context) (string, error) {
	auth := c.Request().Header.Get(jwtHeader)
	l := len(authScheme)
	if len(auth) > l+1 && auth[:l] == authScheme {
		return auth[l+1:], nil
	}
	return "", ErrJWTMissing
}

// jwtFromQuery returns a `jwtExtractor` that extracts token from the query string.
func jwtFromQuery(c echo.Context) (string, error) {
	token := c.QueryParam(jwtParam)
	if token == "" {
		return "", ErrJWTMissing
	}
	return token, nil
}

// jwtFromCookie returns a `jwtExtractor` that extracts token from the named cookie.
func jwtFromCookie(c echo.Context) (string, error) {
	cookie, err := c.Cookie(jwtCookieName)
	if err != nil {
		return "", ErrJWTMissing
	}
	return cookie.Value, nil
}
