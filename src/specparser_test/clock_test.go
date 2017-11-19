package specparser

import (
	"specparser"
	"testing"
	"time"
)

func TestMyClock_Now(t *testing.T) {
	clock := specparser.ClockInterface{}
	expected := clock.Now().Format("1970-12-23 14:05:04")
	actual := time.Now().Format("1970-12-23 14:05:04")

	if expected != actual {
		t.Error("time should match", expected, actual)
	}
}

func TestMyClock_Until(t *testing.T) {
	clock := specparser.ClockInterface{}
	later := time.Now().Add(5 * time.Minute)

	if clock.Until(later) != (5 * time.Minute) {
		t.Error("duration should match")
	}
}
