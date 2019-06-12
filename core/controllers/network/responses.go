package network

import "math/big"

// save result in the cache
type BalanceResponse struct {
	Value *big.Int `json:"value"`
	Raw   string   `json:"raw"`
}
