// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package interval

import (
	"sync/atomic"
	"time"

	"github.com/zerjioang/etherniti/thirdparty/echo"
)

type TaskMode uint8
type OnExpired func(ctx *echo.Context) []byte

const (
	Once TaskMode = iota
	Loop
)

type IntervalTask struct {
	ticker       *time.Ticker
	timer        *time.Timer
	name         string
	mode         TaskMode
	onExpired    OnExpired
	atomicResult atomic.Value
}

func NewTask(name string, d time.Duration, mode TaskMode, doOnce bool, onExpired OnExpired) *IntervalTask {
	t := new(IntervalTask)
	t.name = name
	t.mode = mode
	t.onExpired = onExpired
	if mode == Once {
		//use a timer
		t.timer = time.NewTimer(d)
	} else if mode == Loop {
		//use a ticker
		t.ticker = time.NewTicker(d)
	}
	if doOnce {
		// execute task at least once before time expires
		// useful for initial data populations
		t.triggerExpirationRoutine()
	}
	return t
}

func (task *IntervalTask) Do() *IntervalTask {
	if task.onExpired != nil {
		if task.mode == Once {
			go func() {
				<-task.timer.C
				//timer expired, execute requested action
				task.triggerExpirationRoutine()
			}()
		} else if task.mode == Loop {
			go func() {
				for range task.ticker.C {
					task.triggerExpirationRoutine()
				}
			}()
		}
	}
	return task
}

func (task *IntervalTask) triggerExpirationRoutine() {
	// atomic/thread-safe
	result := task.onExpired(nil)
	task.atomicResult.Store(result)
}

func (task *IntervalTask) Bytes() []byte {
	return task.atomicResult.Load().([]byte)
}

func (task *IntervalTask) Stop() bool {
	if task.mode == Once {
		return task.timer.Stop()
	} else if task.mode == Loop {
		task.ticker.Stop()
	}
	return true
}
