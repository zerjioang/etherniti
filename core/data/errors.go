// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package data

import (
	"errors"

	"github.com/zerjioang/etherniti/core/modules/stack"
)

const (
	FailedToBind             = "failed to bind request data to model: "
	NotAuthorized            = "not authorized"
	NotImplementedStr        = "not implemented"
	ErrUnknownModeStr        = "unknown mode selected. Allowed modes are: single, git, zip, targz"
	SolcVersionFailedStr     = "failed to get solc version"
	BindErrStr               = "there was an error while processing your request information"
	ProvideProjectIdStr      = "you have to provide a valid project data"
	ProvideValidDataStr      = "you have to provide valid data"
	OperationNotSupportedStr = "operation not supported"
	InvalidAddressStr        = "invalid address provided"
	InvalidPayloadStr        = "invalid payload provided"
)

var (
	NotImplemented = []byte("not implemented")

	Symbol                      = []byte("symbol")
	Name                        = []byte("name")
	Allowance                   = []byte("allowance")
	Decimals                    = []byte("decimals")
	TotalSupply                 = []byte("totalSupply")
	Transfer                    = []byte("transfer")
	Deploy                      = []byte("deploy")
	Summary                     = []byte("summary")
	BalanceOf                   = []byte("balanceof")
	Balance                     = []byte("balance")
	BalanceAtBlock              = []byte("balance_at_block")
	EthInfo                     = []byte("eth_info")
	NetVersion                  = []byte("net_version")
	TransactionReceipt          = []byte("transaction_receipt")
	TransactionCount            = []byte("transaction_count")
	TransactionCountHash        = []byte("transaction_count_hash")
	TransactionCountBlockNumber = []byte("transaction_count_block_number")
	CompileLLL                  = []byte("compile_lll")
	CompileSerpent              = []byte("compile_serpent")
	CompileSolidity             = []byte("compile_solidity")
	EstimateGas                 = []byte("estimate_gas")
	GetStorage                  = []byte("get_storage")
	GetCode                     = []byte("get_code")
	EthSign                     = []byte("eth_sign")
	EthSignatureParse           = []byte("eth_sign_parse")

	Sha3 = []byte("sha3")

	PutString = []byte("db_putString")
	GetString = []byte("db_getString")
	PutHex    = []byte("db_putHex")
	GetHex    = []byte("db_getHex")

	ChainId    = []byte("chain_id")
	ShhVersion = []byte("shh_version")

	InvalidUrlWeb3    = []byte("invalid url or web3 method provided")
	NetworkNoResponse = []byte("the network peer did not return any response")

	IsContract        = []byte("is_contract")
	IsGanache         = []byte("is_ganache")
	AccountsBalanced  = []byte("accounts_balanced")
	SolcVersion       = []byte("solc_version")
	SolcCompiled      = []byte("solc_compile")
	DataBindFailedStr = "failed to execute requested operation"
	DataBindFailed    = []byte(DataBindFailedStr)

	SolcVersionFailed = []byte(SolcVersionFailedStr)

	AddressNoSetupStr = "invalid ethereum address setup when creating connection profile. Please provide a valid address as 'from'"

	InvalidContractAddress = []byte("invalid contract address provided")
	InvalidAccountAddress  = []byte("invalid account address provided")
	InvalidMethodName      = []byte("invalid contract method name provided")
	InvalidEtherValue      = []byte("invalid ether amount value provided")
	InvalidTokenValue      = []byte("invalid token amount value provided")
	InvalidAddress         = []byte(InvalidAddressStr)
	InvalidSrcAddress      = []byte("invalid source address (from) provided")
	InvalidDstAddress      = []byte("invalid destination address (to) provided")
	InvalidReceiverAddress = []byte("invalid transfer receiver address provided")
	InvalidAccountOwner    = []byte("invalid account owner address provided")
	InvalidAccountSpender  = []byte("invalid account spender address provided")

	ProvideContractName    = []byte("provide a valid contract name")
	ProvideContractAddress = []byte("provide a valid contract address")
	NoResults              = []byte("no results found for given contract address")
	LinkSuccess            = []byte("abi successfully linked to contract")
	InvalidAbi             = []byte("invalid abi data provided on field 'abi'")

	MissingAddress = []byte("please, provide a valid ethereum or quorum address")

	ContractNameSpaceResolved     = []byte("contract information successfully resolved")
	ContractResolutionFailed      = []byte("failed to resolve given contract id")
	ContractNameSpaceRegistered   = []byte("contract successfully registered in naming service")
	ContractNameSpaceUnregistered = []byte("contract successfully unregistered from naming service")

	ProfileTokenSuccess = []byte("profile token successfully created")

	InfuraJwtErrorMessage   = []byte("please provide an Infura connection profile token including provided Infura endpoint URL (https://$NETWORK.infura.io/v3/$PROJECT_ID) for this kind of call.")
	QuiknodeJwtErrorMessage = []byte("please provide a QuikNode connection profile token including provided full peer endpoint URL")
	UserJwtErrorMessage     = []byte("please provide a valid account token")
	JwtErrorMessage         = []byte("please provide a private connection profile token for this kind of call")

	ReadErr = []byte("there was an error during execution")
	BindErr = []byte(BindErrStr)

	UserLogin          = []byte("login")
	UserRegistered     = []byte("user registered")
	UserRegisterFailed = []byte("failed to register new user account")

	DatabaseError = []byte("failed to process your login request at this moment. Please try it later")

	InvalidLoginData        = []byte("invalid username or password provided")
	MissingLoginFields      = []byte("invalid login data provided. please fill all required fields")
	FailedLoginVerification = []byte("failed to verify your login information at this time. Please try it few minutes later.")
	RegistrationSuccess     = []byte("registration successfully finished. Please verify your account with the message sent to your inbox.")

	MnemonicLanguageNotProvided = []byte("provided language is not supported")
	MnemonicSizeNotSupported    = []byte("provided mnemonic size is not supported. allowed sizes are: 128,160,192,224,256")
	InvalidEntropySource        = []byte("failed to get a full entropy source")
	MnemonicSuccess             = []byte("mnemonic successfully created")
	HDWalletSuccess             = []byte("hd wallet successfully created")
	EntropySizeNotSupported     = []byte("provided entropy size is not supported")
	EntropySuccess              = []byte("entropy data generated")
	EthAccountSuccess           = []byte("ethereum account created")
	EthAccountFailed            = []byte("failed to generate ecdsa private key")
	EthAddressValidation        = []byte("address validation")

	ErrBlockTorConnection = []byte("invalid connection address")
	UserTokenFailed       = []byte("failed to generate user token")
)

var (
	ErrInvalidAddress        = errors.New(InvalidAddressStr)
	ErrInvalidPayload        = errors.New(InvalidPayloadStr)
	ErrInvalidBlockHash      = errors.New("Invalid block hash provided")
	ErrUnknownMode           = errors.New(ErrUnknownModeStr)
	ErrCannotReadSolcVersion = errors.New(SolcVersionFailedStr)
	ErrNotImplemented        = errors.New(NotImplementedStr)
	ErrNotAuthorized         = errors.New(NotAuthorized)
	ErrListingNotSupported   = errors.New("listing not supported")
	ErrBind                  = stack.New(BindErrStr)
	ErrStackProject          = stack.New(ProvideProjectIdStr)
	ErrInvalidData           = stack.New(ProvideValidDataStr)
	ErrOperationNotSupported = stack.New(OperationNotSupportedStr)
	ErrInvalidBlockNumber    = errors.New("provided block number is not valid. remember allowed values are: an hex number, 'earliest', 'latest' or 'pending'")
)

// database related errors
var (
	ErrDuplicateKey = errors.New("duplicate key found on database. cannot store")
)

// project controller related data
var (
	OperationNotSupported = []byte(OperationNotSupportedStr)
	FailedToProcess       = []byte("failed to process current request")
	ProvideProjectId      = []byte(ProvideProjectIdStr)
	ProvideId             = []byte("you have to provide a valid object id")
	SuccessfullyCreated   = []byte("successfully created")
	SuccessfullyDeleted   = []byte("successfully deleted")
	NotAllowedToList      = []byte("you are not allowed to list items")
)

// external controller related data {
var (
	EthPrice  = []byte("ethereum-price")
	EthTicker = []byte("ethereum-ticker")
)

// profile token related errors
var (
	ErrTokenNoValid               = errors.New("provided token contains invalid or missing fields")
	ErrInvalidSigningMethod       = errors.New("unexpected signing method")
	ErrFailedToRead               = errors.New("failed to read token claims")
	ErrMissingAuthentication      = errors.New("authentication token was not found in request")
	ErrInvalidAuthenticationToken = []byte("invalid authentication token")
	ErrMissingAuthenticationToken = []byte("missing authentication token")
)
