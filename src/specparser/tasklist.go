package specparser

import (
	"io/ioutil"
	"log"
	"time"
)

var Debug = log.New(ioutil.Discard, "", 0)

//var Error *log.Logger = log.New(os.Stderr,"ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
//var Debug *log.Logger = log.New(os.Stdout,"->\t", 0)
//var DevelOut *log.Logger = log.New(os.Stdout,"", 0)

type (
	Work     map[time.Time]*TaskSpec // Time indexed pointers to TaskSpec, TaskSpec.Command can then be accessed
	Schedule []time.Time             // Ordered slice of Work indexes, a queue interface would be nice (push, pop - ensure ptr's are cleared)
	TaskList struct {
		Work     Work
		Schedule Schedule
	}
)

func NewWork() Work {
	return make(map[time.Time]*TaskSpec)
}

func NewSchedule() Schedule {
	return make([]time.Time, 0)
}

func (t *TaskList) AddTask(time time.Time, spec *TaskSpec) *TaskList {
	if t.Schedule == nil {
		t.Schedule = NewSchedule()
	}

	if t.Work == nil {
		t.Work = NewWork()
	}

	t.Work[time] = spec
	t.Schedule = append(t.Schedule, time)
	return t
}

/**
  Initialize TaskList object, _build schedule_ -> extract this to own method
*/
func NewTaskList(spec TaskSpec, t time.Time, lookAheadMins int) (taskList TaskList, err error) {
	debugMsg := "Checking next %d slots\n from: %s\n   to: %s\n (inclusive)\n\n"
	Debug.Printf(debugMsg, lookAheadMins, t, t.Add(time.Minute*time.Duration(10)))

	var failMsg string

	for i := 0; i < lookAheadMins; i++ {
		var pass = false
		var minute = Minute(t.Minute())
		var hour = Hour(t.Hour())
		var day = Day(t.Day())
		var month = Month(t.Month())
		var dayOfWeek = DayOfWeek(t.Weekday())

		switch {
		case !spec.HasMonth(month):
			failMsg = "not in month"
			break
		case !spec.HasDayOfWeek(dayOfWeek):
			failMsg = "not in daysOfWeek"
			break
		case !spec.HasDay(day):
			failMsg = "not in days"
			break
		case !spec.HasHour(hour):
			failMsg = "not in hours"
			break
		case !spec.HasMinute(minute):
			failMsg = "not in minutes"
			break
		default:
			pass = true
		}

		Debug.Println("Pass: ", pass)

		if pass {
			taskList.AddTask(t, &spec)
			Debug.Println(spec.Expression, "matches")
			Debug.Printf("%4d-%02d-%02d %02d:%02d:00 +0000 (Day:%d)\n", 2017, month, day, hour, minute, dayOfWeek)
		} else {
			Debug.Println(failMsg)
			Debug.Printf("%4d-%02d-%02d %02d:%02d:00 +0000 (Day:%d)\n", 2017, month, day, hour, minute, dayOfWeek)
		}

		Debug.Println()

		t = t.Add(time.Minute)
	}

	return taskList, err
}
