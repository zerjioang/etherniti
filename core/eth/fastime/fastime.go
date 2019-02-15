package fastime

import "time"

type Duration int64
// Nanoseconds returns the duration as an integer nanosecond count.
func (d Duration) Nanoseconds() int64 { return int64(d) }

const (
	Nanosecond  Duration = 1
	Microsecond          = 1000 * Nanosecond
	Millisecond          = 1000 * Microsecond
	Second               = 1000 * Millisecond
	Minute               = 60 * Second
	Hour                 = 60 * Minute
)

type FastTime struct {
	nsec uint32
	sec  int64
}

func (fastTime FastTime) Unix() int64 {
	return fastTime.sec
}
func (fastTime FastTime) Add(duration Duration) FastTime {
	fastTime.nsec += uint32(duration.Nanoseconds())
	fastTime.sec += duration.Nanoseconds() / 1000000
	return fastTime
}

func Now() FastTime {
	t := time.Now()
	return FastTime{
		nsec:uint32(t.Nanosecond()),
		sec: t.Unix(),
	}
}