// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package gmail

import (
	"testing"
)

func TestNewDefaultMailServerInstantiation(t *testing.T) {
	mailCfg := GetMailServerConfigInstance()
	if mailCfg == nil {
		t.Fatal("DefaultMailServerConfig could not be instantiated")
	}
}
