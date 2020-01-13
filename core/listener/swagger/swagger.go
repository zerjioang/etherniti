// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package swagger

import (
	"io/ioutil"
	"os"
	"strings"

	"github.com/zerjioang/etherniti/util/banner"

	"github.com/zerjioang/etherniti/core/config"
	"github.com/zerjioang/etherniti/core/logger"
	"github.com/zerjioang/etherniti/shared/constants"
	"github.com/zerjioang/go-hpc/thirdparty/gommon/log"
)

var (
	cfg = config.GetDefaultOpts()
)

func ConfigureFromTemplate() {
	configureSwaggerJsonWithDir(config.ResourcesDirSwagger)
}

func configureSwaggerJsonWithDir(resources string) {
	//read template file
	log.Debug("reading swagger json file")
	raw, err := ioutil.ReadFile(resources + "/swagger-template.json")
	if err != nil {
		logger.Error("failed reading swagger template file", err)
		return
	}
	//replace hardcoded variables
	str := string(raw)
	str = strings.Replace(str, "$title", "Etherniti: High Performance Web3 REST Proxy", -1)
	str = strings.Replace(str, "$version", banner.Version, -1)
	str = strings.Replace(str, "$host", config.GetSwaggerAddressWithPort(cfg), -1)
	str = strings.Replace(str, "$basepath", "/v1", -1)
	str = strings.Replace(str, "$header-auth-key", constants.HttpProfileHeaderkey, -1)
	//write swagger.json files
	writeErr := ioutil.WriteFile(resources+"/swagger.json", []byte(str), os.ModePerm)
	if writeErr != nil {
		logger.Error("failed writing swagger.json file", writeErr)
		return
	}
}
