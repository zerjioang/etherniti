package client

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
	"time"

	"github.com/valyala/fasthttp"
	"github.com/zerjioang/etherniti/core/logger"
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
	client *fasthttp.Client
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
		logger.Error(err)
		return
	}
	// Here, we read the cert.pem file and supply it as the root CA when creating the Client.
	// Running the Client should now successfully display the following.
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)
	// On the Client, read and supply the key pair as the client certificate.
	client = &fasthttp.Client{
		ReadTimeout:     time.Second * 3,
		WriteTimeout:    time.Second * 3,
		WriteBufferSize: 2048,
		ReadBufferSize:  2048,
		TLSConfig: &tls.Config{
			Certificates: []tls.Certificate{cert},
			RootCAs:      caCertPool,
		},
	}
}

// executes a mutual tls configured https requests
func mTLSRequest(client *fasthttp.Client, url string, content []byte) ([]byte, error) {
	req := fasthttp.AcquireRequest()
	req.SetRequestURI(url)                    //set URL
	req.Header.SetMethodBytes([]byte("POST")) //set method mode
	req.SetBody(content)                      //set body

	res := fasthttp.AcquireResponse()
	// Request /hello via the created HTTPS client over port 8443 via GET
	err := client.Do(req, res)
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	return req.Body(), nil
}
