// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package fastime_test

import "github.com/zerjioang/etherniti/core/eth/fastime"

func ExampleFastTime() {
	tm2 := fastime.Now()
	u := tm2.Unix()
	if u > 0 {

	}
}
