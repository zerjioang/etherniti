package fastime

import "time"

type FastTime struct {
	//nanoseconds
	nsec uint32
	//unix time
	sec  int64
}

// Now returns the current local time.
func Now() FastTime {
	//get current time
	currentTime := time.Now()
	return FastTime{
		nsec: uint32(currentTime.Nanosecond()),
		sec:  currentTime.Unix(),
	}
}

// get unix time function
func (t FastTime) Unix() int64{
	return t.sec
}

// add time to current time
func (t FastTime) Add(value time.Duration) FastTime{
	t.nsec += uint32(value.Nanoseconds())
	t.sec += value.Nanoseconds() / 1000000000
	return t
}