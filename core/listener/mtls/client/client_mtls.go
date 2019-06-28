package mtlsclient

import (
	"crypto/tls"
	"crypto/x509"
	"github.com/zerjioang/etherniti/core/logger"
	"io/ioutil"
	"log"
	"net/http"
)

/*

How to generate cert.pem and key.pem (self-signed) for development purposed
openssl req -newkey rsa:2048 \
  -new -nodes -x509 \
  -days 3650 \
  -out cert.pem \
  -keyout key.pem \
  -subj "/C=US/ST=California/L=Mountain View/O=Your Organization/OU=Your Unit/CN=localhost"

*/

var (
	client *http.Client
)

func init() {
	// Load our TLS key pair to use for authentication
	logger.Info("loading mTLS key pair to use for authentication purposes")
	cert, err := tls.LoadX509KeyPair("../cert.pem", "../key.pem")
	if err != nil {
		log.Fatalln("Unable to load cert", err)
	}
	logger.Info("loading CA certificate pool")
	// Load our CA certificate cert.pem
	caCert, err := ioutil.ReadFile("cert.pem")
	if err != nil {
		logger.Error(err); return
	}
	// Here, we read the cert.pem file and supply it as the root CA when creating the Client.
	// Running the Client should now successfully display the following.
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)
	// On the Client, read and supply the key pair as the client certificate.
	// Create a HTTPS client and supply the created CA pool
	client = &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				Certificates: []tls.Certificate{cert},
				RootCAs: caCertPool,
			},
		},
	}
}

// executes a mutual tls configured https requests
func MakeRequest(url string) ([]byte, error) {
	var body []byte
	var err error
	// Request /hello via the created HTTPS client over port 8443 via GET
	r, err := client.Get(url)
	if err != nil {
		logger.Error(err)
		return body, err
	}
	// Read the response body
	body, err = ioutil.ReadAll(r.Body)
	if err != nil {
		logger.Error(err)
		return body, err
	}
	_ = r.Body.Close()
	return body, nil
}