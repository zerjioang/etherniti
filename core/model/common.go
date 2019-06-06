package model

import "github.com/zerjioang/etherniti/core/modules/stack"

var (
	UnsupportedDataErr = stack.New("unsupported data")
)

// helper methods
func ConditionalStringUpdate(newValue string, lastValue string, defaultValue string) string {
	if newValue != lastValue && newValue != defaultValue {
		return newValue
	}
	return lastValue
}
