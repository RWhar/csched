package specparser

import (
	"errors"
	"fmt"
	"strings"
)

const CommandMaxStringLength = 999

type TaskSpec struct {
	Expression string
	Schedule   TimeSpecExtended
	Command    string
}

func NewTaskSpec(spec string) (taskSpec TaskSpec, err error) {
	parts := strings.Fields(strings.Trim(spec, " "))

	if len(parts) < 6 {
		err = errors.New(fmt.Sprintf("%s %d", "invalid spec only has", len(parts)))
		return
	}

	extendedTimeSpec, err := TimeExpression{}.New(parts[0], parts[1], parts[2], parts[3], parts[4]).Explode()
	//extendedTimeSpec, err := TimeExpression{
	//		Minute:    ValueExpression(parts[0]),
	//		Hour:      ValueExpression(parts[1]),
	//		Day:       ValueExpression(parts[2]),
	//		Month:     ValueExpression(parts[3]),
	//		DayOfWeek: ValueExpression(parts[4]),
	//	}.Explode()

	taskSpec = TaskSpec{
		Expression: strings.Join(parts[0:5], " "),
		Schedule:   extendedTimeSpec,
		Command:    strings.Join(parts[5:], " "),
	}

	if len(taskSpec.Command) > CommandMaxStringLength {
		err = errors.New("command exceeds maximum length of " + string(CommandMaxStringLength) + " chars")
	}

	return taskSpec, err
}

func (s *TaskSpec) HasMinute(minute Minute) bool {
	for i := range s.Schedule.Minutes {
		if s.Schedule.Minutes[i] == minute {
			return true
		}
	}

	return false
}

func (s *TaskSpec) HasHour(hour Hour) bool {
	for i := range s.Schedule.Hours {
		if s.Schedule.Hours[i] == hour {
			return true
		}
	}

	return false
}

func (s *TaskSpec) HasDay(day Day) bool {
	for i := range s.Schedule.Days {
		if s.Schedule.Days[i] == day {
			return true
		}
	}

	return false
}

func (s *TaskSpec) HasMonth(month Month) bool {
	for i := range s.Schedule.Months {
		if s.Schedule.Months[i] == month {
			return true
		}
	}

	return false
}

func (s *TaskSpec) HasDayOfWeek(dayOfWeek DayOfWeek) bool {
	// Handle zero as sunday
	if dayOfWeek == 0 {
		dayOfWeek = 7
	}

	for i := range s.Schedule.DaysOfWeek {
		if s.Schedule.DaysOfWeek[i] == dayOfWeek {
			return true
		}
	}

	return false
}
