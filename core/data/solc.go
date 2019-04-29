package data

type CompilerMode string

const (
	SingleRawFile    = "single-raw"
	SingleBase64File = "single-base64"
	GitMode          = "git"
	ZipMode          = "zip"
	TargzMode        = "targz"
)

