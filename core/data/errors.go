// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package data

import (
	"errors"
)

const (
	NotAuthorized        = "not authorized"
	NotImplementedStr    = "not implemented"
	ErrUnknownModeStr    = "unknown mode selected. Allowed modes are: single, git, zip, targz"
	SolcVersionFailedStr = "failed to get solc version"
)

var (
	NotImplemented = []byte("not implemented")

	Symbol             = []byte("symbol")
	Name               = []byte("name")
	Allowance          = []byte("allowance")
	Decimals           = []byte("decimals")
	TotalSupply        = []byte("totalSupply")
	Transfer           = []byte("transfer")
	Deploy             = []byte("deploy")
	Summary            = []byte("summary")
	BalanceOf          = []byte("balanceof")
	Balance            = []byte("balance")
	BalanceAtBlock     = []byte("balance_at_block")
	EthInfo            = []byte("eth_info")
	NetVersion         = []byte("net_version")
	TransactionReceipt = []byte("transaction_receipt")
	ShhVersion         = []byte("shh_version")

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

	AddressNoSetupStr = "invalid ethereum address setup when creating connection profile. Please provide a alid address as 'from'"

	InvalidContractAddress = []byte("invalid contract address provided")
	InvalidAccountAddress  = []byte("invalid account address provided")
	InvalidMethodName      = []byte("invalid contract method name provided")
	InvalidEtherValue      = []byte("invalid ether amount value provided")
	InvalidTokenValue      = []byte("invalid token amount value provided")
	InvalidDstAddress      = []byte("invalid destination address provided")
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
	JwtErrorMessage         = []byte("please provide a connection profile token for this kind of call")

	ReadErr = []byte("there was an error during execution")
	BindErr = []byte("there was an error while processing your request information")

	UserLogin          = []byte("login")
	UserRegistered     = []byte("user registered")
	UserRegisterFailed = []byte("failed to register new user account")

	DatabaseError = []byte("Failed to process your login request at this moment. Please try it later")

	InvalidLoginData        = []byte("Invalid username or password provided")
	MissingLoginFields      = []byte("Invalid login data provided. please fill all required fields")
	FailedLoginVerification = []byte("Failed to verify your login information at this time. Please try it few minutes later.")
	RegistrationSuccess     = []byte("Registration successfully finished. Please verify your account with the message sent to your inbox.")

	MnemonicLanguageNotProvided = []byte("provided language is not supported")
	MnemonicSizeNotSupported    = []byte("provided mnemonic size is not supported. allowed sizes are: 128,160,192,224,256")
	InvalidEntropySource        = []byte("failed to get a full entropy source")
	MnemonicSuccess             = []byte("mnemonic successfully created")
	HDWalletSuccess             = []byte("hd wallet successfully created")
	EntropySizeNotSupported     = []byte("provided entropy size is not supported")
	EntropySuccess              = []byte("Entropy data generated")
	EthAccountSuccess           = []byte("ethereum account created")
	EthAccountFailed            = []byte("failed to generate ecdsa private key")
	EthAddressValidation        = []byte("address validation")

	ErrBlockTorConnection = []byte("invalid connection address")
)

var (
	ErrUnknownMode           = errors.New(ErrUnknownModeStr)
	ErrCannotReadSolcVersion = errors.New(SolcVersionFailedStr)
	ErrNotImplemented        = errors.New(NotImplementedStr)
	ErrNotAuthorized         = errors.New(NotAuthorized)
	ListingNotSupported      = errors.New("listing not supported")
)

// database related errors
var (
	DuplicateKeyErr = errors.New("duplicate key found on database. cannot store")
)

// project controller related errors
var (
	FailedToProcess     = []byte("failed to process current request")
	ProvideProjectId    = []byte("you have to provide a valid project id")
	ProvideId           = []byte("you have to provide a valid object id")
	SuccessfullyCreated = []byte("successfully created")
	SuccessfullyDeleted = []byte("successfully deleted")
	NotAllowedToList    = []byte("you are not allowed to list items")
)
