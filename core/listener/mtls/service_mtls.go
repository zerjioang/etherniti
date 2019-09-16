// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package mtls

import (
	"crypto/tls"
	"crypto/x509"
	"net/http"

	"github.com/zerjioang/etherniti/core/config"
	"github.com/zerjioang/etherniti/core/listener/common"
	base "github.com/zerjioang/etherniti/core/listener/http"
	"github.com/zerjioang/etherniti/core/listener/https"
	"github.com/zerjioang/etherniti/core/logger"
	"github.com/zerjioang/etherniti/shared/def/listener"
)

var (
	//default etherniti proxy configuration
	cfg = config.GetDefaultOpts()
	//store server mutual tls configuration parameters
	tlsConfig *tls.Config
)

// based on https.HttpsListener to implement mTLS
type MtlsListener struct {
	https.HttpsListener
}

func recoverFromPem() {
	if r := recover(); r != nil {
		logger.Error("recovered from pem", r)
	}
}

func init() {
	// On the Server, we create a similar CA pool and supply it
	// to the TLS config to serve as the authority to validate Client certificates.
	// We also use the same key pair for the Server certificate.
	// Create a CA certificate pool and add cert.pem to it
	logger.Info("loading mTLS configuration and internal verification pool")
	defer recoverFromPem()
	caCertBytes := cfg.GetCertPem()
	if caCertBytes == nil || len(caCertBytes) == 0 {
		logger.Error("failed to load mTLS server certificate content")
	}

	// create certification authority pool
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCertBytes)

	// Create the TLS Config with the CA pool and enable Client certificate validation
	// The first two items are the most imporant!
	// Without them there is a potential authentication bypass vulnerability.
	tlsConfig = &tls.Config{
		// TLS 1.2 because we can
		MinVersion:       tls.VersionTLS12,
		CurvePreferences: []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
		// PFS because we can but this will reject client with RSA certificates
		CipherSuites: []uint16{
			tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
			tls.TLS_RSA_WITH_AES_128_GCM_SHA256,
			tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_RSA_WITH_AES_256_CBC_SHA,
		},
		// Ensure that we only use our "CA" to validate certificates
		ClientCAs: caCertPool,
		// Reject any TLS certificate that cannot be validated
		ClientAuth: tls.RequireAndVerifyClientCert,
		// Force it server side
		PreferServerCipherSuites: true,
	}
	tlsConfig.BuildNameToCertificate()
}

// fetch specific server configuration
// in this case, we return basic http configuration + mtls configuration
func (l MtlsListener) ServerConfig() *http.Server {
	mtlsServer := common.DefaultHttpsServerConfig
	mtlsServer.TLSConfig = tlsConfig
	return mtlsServer
}

// create new deployer instance
func NewMtlsListener() listener.ListenerInterface {
	d := MtlsListener{}
	d.HttpListener = base.NewHttpListenerCustom()
	return d
}
