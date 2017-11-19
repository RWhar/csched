package specparser

import (
	"specparser"
	"testing"
)

func TestTaskSpec_HasMinute(t *testing.T) {
	taskSpec := specparser.TaskSpec{
		Schedule: specparser.TimeSpecExtended{
			Minutes: []specparser.TimeUnit{
				specparser.Minute(12),
			},
		},
	}

	if !taskSpec.HasMinute(12) {
		t.Error("should match when item is in extended time spec")
	}

	if taskSpec.HasMinute(1) {
		t.Error("should not match when item is not in time spec")
	}
}

func TestTaskSpec_HasHour(t *testing.T) {
	taskSpec := specparser.TaskSpec{
		Schedule: specparser.TimeSpecExtended{
			Hours: []specparser.TimeUnit{
				specparser.Hour(14),
			},
		},
	}

	if !taskSpec.HasHour(14) {
		t.Error("should match when item is in extended time spec")
	}

	if taskSpec.HasHour(1) {
		t.Error("should not match when item is not in time spec")
	}

}

func TestTaskSpec_HasDay(t *testing.T) {
	taskSpec := specparser.TaskSpec{
		Schedule: specparser.TimeSpecExtended{
			Days: []specparser.TimeUnit{
				specparser.Day(28),
			},
		},
	}

	if !taskSpec.HasDay(28) {
		t.Error("should match when item is in extended time spec")
	}

	if taskSpec.HasDay(1) {
		t.Error("should not match when item is not in time spec")
	}

}

func TestTaskSpec_HasMonth(t *testing.T) {
	taskSpec := specparser.TaskSpec{
		Schedule: specparser.TimeSpecExtended{
			Months: []specparser.TimeUnit{
				specparser.Month(9),
			},
		},
	}

	if !taskSpec.HasMonth(9) {
		t.Error("should match when item is in extended time spec")
	}

	if taskSpec.HasMonth(1) {
		t.Error("should not match when item is not in time spec")
	}

}

func TestTaskSpec_HasDayOfWeek(t *testing.T) {
	taskSpec := specparser.TaskSpec{
		Schedule: specparser.TimeSpecExtended{
			DaysOfWeek: []specparser.TimeUnit{
				specparser.DayOfWeek(3),
			},
		},
	}

	if !taskSpec.HasDayOfWeek(3) {
		t.Error("should match when item is in extended time spec")
	}

	if taskSpec.HasDayOfWeek(1) {
		t.Error("should not match when item is not in time spec")
	}

}

func TestTaskSpec_HasDayOfWeekConvertsZero(t *testing.T) {
	taskSpec := specparser.TaskSpec{
		Schedule: specparser.TimeSpecExtended{
			DaysOfWeek: []specparser.TimeUnit{
				specparser.DayOfWeek(7),
			},
		},
	}

	if !taskSpec.HasDayOfWeek(0) {
		t.Error("should match as DoW 0 should be interpreted as day 7 - Sunday")
	}
}

func TestTaskSpec_New(t *testing.T) {
	_, err := specparser.NewTaskSpec("* * * * * command")

	if err != nil {
		t.Error("Basic init failed")
	}

	_, err = specparser.NewTaskSpec("* * * * * ")

	if err == nil {
		t.Error("Missing command did not generate error")
	}

	_, err = specparser.NewTaskSpec("* * * * command ")

	if err == nil {
		t.Error("Missing TimeSpecPart did not generate error")
	}

	_, err = specparser.NewTaskSpec("s * * * * command ")

	if err == nil {
		t.Error("Invalid value for minutes did not generate error")
	}
}
