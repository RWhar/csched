package specparser

type TimeUnit interface {
	ToInt() int
}

type TimeExpression struct {
	Minute    ValueExpression
	Hour      ValueExpression
	Day       ValueExpression
	Month     ValueExpression
	DayOfWeek ValueExpression
}

type TimeSpecExtended struct {
	Minutes    []TimeUnit
	Hours      []TimeUnit
	Days       []TimeUnit
	Months     []TimeUnit
	DaysOfWeek []TimeUnit
}

type (
	Minute    int
	Hour      int
	Day       int
	Month     int
	DayOfWeek int
)

func (t TimeExpression) New(minute string, hour string, day string, month string, dayOfWeek string) *TimeExpression {
	t.Minute = ValueExpression(minute)
	t.Hour = ValueExpression(hour)
	t.Day = ValueExpression(day)
	t.Month = ValueExpression(month)
	t.DayOfWeek = ValueExpression(dayOfWeek)

	return &t
}

func (m Minute) ToInt() int {
	return int(m)
}

func (m Hour) ToInt() int {
	return int(m)
}

func (m Day) ToInt() int {
	return int(m)
}

func (m Month) ToInt() int {
	return int(m)
}

func (m DayOfWeek) ToInt() int {
	return int(m)
}

func (t TimeExpression) Explode() (timeSpecExtended TimeSpecExtended, err error) {
	if timeSpecExtended.Minutes, err = t.Minute.Expand(TimeUnitMinutes); err != nil {
		return timeSpecExtended, err
	}

	if timeSpecExtended.Hours, err = t.Hour.Expand(TimeUnitHours); err != nil {
		return timeSpecExtended, err
	}

	if timeSpecExtended.Days, err = t.Day.Expand(TimeUnitDays); err != nil {
		return timeSpecExtended, err
	}

	if timeSpecExtended.Months, err = t.Month.Expand(TimeUnitMonths); err != nil {
		return timeSpecExtended, err
	}

	if timeSpecExtended.DaysOfWeek, err = t.DayOfWeek.Expand(TimeUnitDaysOfWeek); err != nil {
		return timeSpecExtended, err
	}

	return timeSpecExtended, err
}
