package specparser

import (
	"specparser"
	"testing"
	"time"
)

func TestMyClock_Now(t *testing.T) {
	clock := specparser.ClockInterface{}

	if clock.Now() != time.Now() {
		t.Error("time should match")
	}
}

func TestMyClock_Until(t *testing.T) {
	clock := specparser.ClockInterface{}
	later := time.Now().Add(5 * time.Minute)

	if clock.Until(later) != (5 * time.Minute) {
		t.Error("duration should match")
	}
}
