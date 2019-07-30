package bus

import (
	"github.com/zerjioang/etherniti/core/logger"
	"github.com/zerjioang/go-bus"
	"github.com/zerjioang/go-bus/mutex"
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

func Subscribe(topic string, listener gobus.EventListener) {
	sb.Subscribe(topic, listener)
}

func Emit(topic string) {
	sb.Emit(topic)
}
