package constants

type UserRole uint8

const (
	AdminUser UserRole = iota
	StandardUser
	PremiumUserUser
)
