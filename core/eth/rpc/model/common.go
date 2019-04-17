package model

type BlockPeriod string
const (
	// The default block parameter
	//	When requests are made that act on the state of ethereum, the last default block parameter determines the height of the block.
	//
	//The following options are possible for the defaultBlock parameter:
	//
	//HEX String - an integer block number
	//String "earliest" for the earliest/genesis block
	//String "latest" - for the latest mined block
	//String "pending" - for the pending state/transactions

	NoPeriod   BlockPeriod = "none"
	LatestBlockNumber   BlockPeriod = "latest"
	EarliestBlockNumber  BlockPeriod = "earliest"
	pendingBlockNumber  BlockPeriod = "pending"
)
