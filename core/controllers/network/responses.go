package network

import "math/big"

// save result in the cache
type BalanceResponse struct {
	Value *big.Int `json:"wei"`
	Raw   string   `json:"raw"`
	Eth   string   `json:"eth"`
}
