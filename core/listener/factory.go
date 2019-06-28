// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package listener

import (
	"github.com/zerjioang/etherniti/core/listener/http"
	"github.com/zerjioang/etherniti/core/listener/https"
	"github.com/zerjioang/etherniti/core/listener/mtls"
	"github.com/zerjioang/etherniti/core/listener/socket"
	"github.com/zerjioang/etherniti/shared/def/listener"
)

func FactoryListener(typeof listener.ServiceType) listener.ListenerInterface {
	switch typeof {
	case listener.HttpMode:
		return http.NewHttpListener()
	case listener.HttpsMode:
		return https.NewHttpsListener()
	case listener.MTLSMode:
		return mtls.NewMtlsListener()
	case listener.UnixMode:
		return socket.NewSocketListener()
	default:
		return socket.NewSocketListener()
	}
}
