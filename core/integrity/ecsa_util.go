package integrity

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"crypto/x509"
	"encoding/asn1"
	"encoding/hex"
	"encoding/pem"
	"math/big"

	"github.com/zerjioang/etherniti/core/config"
	"github.com/zerjioang/etherniti/core/util"
)

var (
	zero = big.NewInt(0)
	// sha256 hash
	h = sha256.New()
	//decode private key
	integrityPrivKey *ecdsa.PrivateKey
	//private integrity key bytes
	privateBytes = []byte(config.IntegrityPrivateKeyPem)
	publicBytes = []byte(config.IntegrityPublicKeyPem)
)

func init(){
	integrityPrivKey, _ = decode(privateBytes, publicBytes)
}

func SignMsgWithIntegrity(message string) (string, string) {
	//create test message
	str := util.Bytes(message)
	// hash test message
	h.Reset()
	h.Write(str)
	signhash := h.Sum(nil)
	hexhash := hex.EncodeToString(signhash)
	r, s, _ := ecdsaSign(signhash, integrityPrivKey)
	signature := PointsToDER(r, s)
	return hexhash, signature
}

// create a ecdsa signature
func ecdsaSign(message []byte, priv *ecdsa.PrivateKey) (r, s *big.Int, err error) {
	r = zero
	s = zero
	r, s, err = ecdsa.Sign(rand.Reader, priv, message)
	return r, s, err
}

// verify given ecdsa signature
func ecdsaVerify(hash []byte, r *big.Int, s *big.Int, pub *ecdsa.PublicKey) bool {
	return ecdsa.Verify(pub, hash, r, s)
}

// encode to pem both keys
func encode(privateKey *ecdsa.PrivateKey, publicKey *ecdsa.PublicKey) (string, string) {
	x509Encoded, _ := x509.MarshalECPrivateKey(privateKey)
	pemEncoded := pem.EncodeToMemory(&pem.Block{Type: "PRIVATE EC KEY", Bytes: x509Encoded})

	x509EncodedPub, _ := x509.MarshalPKIXPublicKey(publicKey)
	pemEncodedPub := pem.EncodeToMemory(&pem.Block{Type: "PUBLIC KEY", Bytes: x509EncodedPub})

	return util.ToString(pemEncoded), util.ToString(pemEncodedPub)
}

// decode private key and public from pem
func decode(pemEncoded []byte, pemEncodedPub []byte) (*ecdsa.PrivateKey, *ecdsa.PublicKey) {
	block, _ := pem.Decode(pemEncoded)
	x509Encoded := block.Bytes
	privateKey, _ := x509.ParseECPrivateKey(x509Encoded)

	blockPub, _ := pem.Decode(pemEncodedPub)
	x509EncodedPub := blockPub.Bytes
	genericPublicKey, _ := x509.ParsePKIXPublicKey(x509EncodedPub)
	publicKey := genericPublicKey.(*ecdsa.PublicKey)

	return privateKey, publicKey
}

// Convert an ECDSA signature (points R and S) to a byte array using ASN.1 DER encoding.
func PointsToDER(r, s *big.Int) string {
	type ecdsaWrapper struct {
		R, S *big.Int
	}
	sequence := ecdsaWrapper{r, s}
	encoding, _ := asn1.Marshal(sequence)
	return hex.EncodeToString(encoding)
}
