package config

import "os"

var (
	//cert content as bytes
	certPemBytes []byte
	//key content as bytes
	keyPemBytes []byte
	//system gopath
	gopath = os.Getenv("GOPATH")
	// current executable process id
	Pid = os.Getpid()
	//resources base dir
	ResourcesDir = gopath + "/src/github.com/zerjioang/etherniti/testdata"
	// define internal folders
	ResourcesDirInternal          = ResourcesDir + "/internal"
	ResourcesDirInternalSecurity  = ResourcesDirInternal + "/security"
	ResourcesDirInternalBots      = ResourcesDirInternal + "/bots"
	ResourcesDirInternalPwd       = ResourcesDirInternal + "/pwd"
	ResourcesDirInternalTor       = ResourcesDirInternal + "/tor"
	ResourcesDirInternalTokens    = ResourcesDirInternal + "/tokens"
	ResourcesDirInternalBadIps    = ResourcesDirInternal + "/badips"
	ResourcesDirInternalCors      = ResourcesDirInternal + "/cors"
	ResourcesDirInternalHosts     = ResourcesDirInternal + "/hosts"
	ResourcesDirInternalTemplates = ResourcesDirInternal + "/templates"
	ResourcesDirInternalEmail     = ResourcesDirInternal + "/templates/mail"
	ResourcesDirRoot              = ResourcesDir + "/root"
	ResourcesIndexHtml            = ResourcesDirRoot + "/index.html"
	ResourcesDirSwagger           = ResourcesDirRoot + "/swagger"
	// define internal files
	ResourcesDirPHP         = ResourcesDirRoot + "/phpinfo.php"
	BlacklistedDomainFile   = ResourcesDirInternalSecurity + "/domains.json"
	TokenListFile           = ResourcesDirInternalTokens + "/list.json"
	PhishingDomainFile      = ResourcesDirInternalSecurity + "/phishing.json"
	AntiBotsFile            = ResourcesDirInternalBots + "/bots"
	TorExitFile             = ResourcesDirInternalTor + "/tor_exit"
	TorAllFile              = ResourcesDirInternalTor + "/tor_all"
	BadIpsFile              = ResourcesDirInternalBadIps + "/list_any_5"
	CorsFile                = ResourcesDirInternalCors + "/cors"
	HostsFile               = ResourcesDirInternalHosts + "/hosts"
	BlacklistedPasswordFile = ResourcesDirInternalPwd + "/blacklist"
)
