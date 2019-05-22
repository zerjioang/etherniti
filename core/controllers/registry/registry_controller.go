// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package registry

import (
	"github.com/zerjioang/etherniti/core/controllers/common"
	"github.com/zerjioang/etherniti/core/logger"
	"github.com/zerjioang/etherniti/core/model/registry"
)

type RegistryController struct {
	common.DatabaseController
}

// constructor like function
func NewRegistryController() RegistryController {
	pc := RegistryController{}
	var err error
	pc.DatabaseController, err = common.NewDatabaseController("registry", registry.NewDBRegistry)
	if err != nil {
		logger.Error("failed to create registry controller ", err)
	}
	return pc
}
