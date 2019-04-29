package data

var(
	Symbol = []byte("symbol")
	Name = []byte("name")
	Allowance = []byte("allowance")
	Decimals = []byte("decimals")
	TotalSupply = []byte("totalSupply")
	Transfer = []byte("transfer")
	Deploy = []byte("deploy")
	Summary = []byte("summary")
	BalanceOf = []byte("balanceof")
	Balance = []byte("balance")
	TransactionReceipt =[]byte("transaction_receipt")

	IsContract = []byte("is_contract")
	IsGanache = []byte("is_ganache")
	AccountsBalanced = []byte("accounts_balanced")

	DataBindFailedStr = "failed to execute requested operation"
	DataBindFailed = []byte(DataBindFailedStr)

	AddressNoSetupStr = "invalid ethereum address setup when creating connection profile. Please provide a alid address as 'from'"

	InvalidContractAddress = []byte("invalid contract address provided")
	InvalidAccountAddress = []byte("invalid account address provided")
	InvalidMethodName = []byte("invalid contract method name provided")
	InvalidEtherValue = []byte("invalid ether amount value provided")
	InvalidTokenValue = []byte("invalid token amount value provided")
	InvalidDstAddress = []byte("invalid destination address provided")
	InvalidReceiverAddress = []byte("invalid transfer receiver address provided")
	InvalidAccountOwner = []byte("invalid account owner address provided")
	InvalidAccountSpender = []byte("invalid account spender address provided")

	ProvideContractName = []byte("provide a valid contract name")
	ProvideContractAddress = []byte("provide a valid contract address")
	NoResults = []byte("no results found for given contract address")
	LinkSuccess = []byte("abi successfully linked to contract")
	InvalidAbi = []byte("invalid abi data provided on field 'abi'")

	MissingAddress = []byte("please, provide a valid ethereum or quorum address")
)
