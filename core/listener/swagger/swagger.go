// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package swagger

import (
	"io/ioutil"
	"os"
	"strings"

	"github.com/zerjioang/etherniti/core/config"
	"github.com/zerjioang/etherniti/core/logger"
	"github.com/zerjioang/etherniti/shared/constants"
	"github.com/zerjioang/etherniti/thirdparty/gommon/log"
)

func ConfigureFromTemplate() {
	configureSwaggerJsonWithDir(config.ResourcesDir)
}

func configureSwaggerJsonWithDir(resources string) {
	//read template file
	log.Debug("reading swagger json file")
	raw, err := ioutil.ReadFile(resources + "/swagger/swagger-template.json")
	if err != nil {
		logger.Error("failed reading swagger template file", err)
		return
	}
	//replace hardcoded variables
	str := string(raw)
	str = strings.Replace(str, "$title", "Etherniti REST API Proxy", -1)
	str = strings.Replace(str, "$version", constants.Version, -1)
	str = strings.Replace(str, "$host", config.SwaggerAddress, -1)
	str = strings.Replace(str, "$basepath", "/v1", -1)
	str = strings.Replace(str, "$header-auth-key", constants.HttpProfileHeaderkey, -1)
	//write swagger.json file
	writeErr := ioutil.WriteFile(resources+"/swagger/swagger.json", []byte(str), os.ModePerm)
	if writeErr != nil {
		logger.Error("failed writing swagger.json file", writeErr)
		return
	}
}
