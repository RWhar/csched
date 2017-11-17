package specparser

import "time"

type Clock interface {
	Now() time.Time
	Until(d time.Duration) time.Time
	Wait(until time.Duration)
}

type ClockInterface struct{}

func (ClockInterface) Now() time.Time { return time.Now() }

func (ClockInterface) Until(t time.Time) time.Duration { return t.Sub(time.Now()) }

func (ClockInterface) Wait(until time.Duration) {
	timer := time.NewTimer(until)

	<-timer.C
	timer.Stop()
}
