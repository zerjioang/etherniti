// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package listener

import "net/http"

type ListenerInterface interface {
	Listen(notifier chan error)
	ServerConfig() *http.Server
}

type ServiceType uint8

const (
	HttpMode ServiceType = iota
	HttpsMode
	MTLSMode
	UnixMode
	WebsocketMode
	SecureWebsocketMode
	UnknownMode
)
