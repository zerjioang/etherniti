package integrity

import (
	"crypto/rand"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"testing"
)

func TestIntegrity(t *testing.T) {

	//read private key
	block, _ := pem.Decode([]byte(integrityPrivateKey))
	key, _ := x509.ParsePKCS1PrivateKey(block.Bytes)
	fmt.Println(key.N)

	message := []byte("the code must be like a piece of music")
	label := []byte("")
	hash := sha256.New()
	signed, err := key.Sign(rand.Reader, []byte(message), opt)
	if err != nil {
		fmt.Errorf("could not sign request: %v", err)
	}
	sig := base64.StdEncoding.EncodeToString(signed)
	fmt.Printf("Signature: %v\n", sig)

	parser, perr := loadPublicKey("public.pem")
	if perr != nil {
		fmt.Errorf("could not sign request: %v", err)
	}
}
