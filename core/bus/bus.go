package bus

import (
	"github.com/zerjioang/etherniti/core/logger"
	"github.com/zerjioang/go-bus/mutex"
)

const (
	PowerOffEvent = "poweroff"
)

var (
	// define global shared bus, thread safe
	sb *mutex.Bus
)

func init() {
	logger.Info("creating internal event bus")
	sb = mutex.NewBusPtr()
}
func SharedBus() *mutex.Bus {
	return sb
}
