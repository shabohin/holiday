package clock

import "time"

//nolint: lll
//go:generate mockgen -destination mock/clock_mock.go github.com/shabohin/holiday.git/pkg/clock Clock

type Clock interface {
	Now() time.Time
}

type RealClock struct {
}

func NewRealClock() Clock {
	return &RealClock{}
}

func (c RealClock) Now() time.Time {
	return time.Now()
}
