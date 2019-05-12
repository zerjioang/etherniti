package config

import "os"

// database config
var (
	Home            = os.Getenv("HOME")
	DatabaseRootPath = Home + "/.etherniti/"
)
