package constants

type RequestScheme uint8

const (
	Http RequestScheme = iota
	Https
	Unix
	Websocket
	Other
)
