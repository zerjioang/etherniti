package constants

type UserRole uint8

const (
	UndefinedUser UserRole = iota
	AdminUser
	StandardUser
	PremiumUserUser
)
