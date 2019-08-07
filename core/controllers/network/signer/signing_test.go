package signer

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zerjioang/etherniti/shared/protocol"
)

func TestSigning(t *testing.T) {
	test := protocol.EthSignRequest{
		From:       "",
		To:         "0x4592d8f8d7b001e72cb26a73e4fa1806a51ac79d",
		PrivateKey: "fad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19",
	}
	assert.NotNil(t, test)
}
