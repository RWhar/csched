package specparser

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

type (
	ValueExpression string
	TimeUnitType    int
	ValueSet        []TimeUnit
)

const (
	TimeUnitMinutes TimeUnitType = iota
	TimeUnitHours
	TimeUnitDays
	TimeUnitMonths
	TimeUnitDaysOfWeek
)

func (v *ValueExpression) IsWildCard() bool {
	return *v == "*"
}

func (v *ValueExpression) ToString() string {
	return string(*v)
}

func (v *ValueExpression) IsSimple() bool {
	return regexp.MustCompile(`^[0-9]{1,2}$`).MatchString(v.ToString())
}

func (v *ValueExpression) IsList() bool {
	pattern := `^([0-9]{1,2}(-\s*[0-9]{1,2})?)((,\s*[0-9]{1,2}(-\s*[0-9]{1,2})?)+)$`
	return regexp.MustCompile(pattern).MatchString(v.ToString())
}

func (v *ValueExpression) IsRange() bool {
	return regexp.MustCompile(`^([0-9]{1,2}-\s*[0-9]{1,2})$`).MatchString(v.ToString())
}

func (v *ValueExpression) IsInterval() bool {
	return regexp.MustCompile(`^([0-9]{1,2}|\*)(/[0-9]{1,2})$`).MatchString(v.ToString())
}

func seq(first int, last int, timeUnitType TimeUnitType) (a []TimeUnit) {
	for i := first; i <= last; i++ {
		switch timeUnitType {
		case TimeUnitMinutes:
			a = append(a, Minute(i))
			break
		case TimeUnitHours:
			a = append(a, Hour(i))
			break
		case TimeUnitDays:
			a = append(a, Day(i))
			break
		case TimeUnitMonths:
			a = append(a, Month(i))
			break
		case TimeUnitDaysOfWeek:
			a = append(a, DayOfWeek(i))
			break
		default:
			panic(0)
		}
	}

	return a
}

var Minutes = seq(0, 59, TimeUnitMinutes)
var Hours = seq(0, 23, TimeUnitHours)
var Days = seq(1, 31, TimeUnitDays)
var Months = seq(1, 12, TimeUnitMonths)
var DaysOfWeek = seq(1, 7, TimeUnitDaysOfWeek)

func (e *ValueSet) appendAll(timeUnitType TimeUnitType) (err error) {
	switch timeUnitType {
	case TimeUnitMinutes:
		*e = append(*e, Minutes...)
		break
	case TimeUnitHours:
		*e = append(*e, Hours...)
		break
	case TimeUnitDays:
		*e = append(*e, Days...)
		break
	case TimeUnitMonths:
		*e = append(*e, Months...)
		break
	case TimeUnitDaysOfWeek:
		*e = append(*e, DaysOfWeek...)
		break
	default:
		return errors.New("unknown unit type")
	}

	return
}

func (e *ValueSet) appendSimple(valueExpression ValueExpression, timeUnitType TimeUnitType) (err error) {
	intVal, err := strconv.Atoi(valueExpression.ToString())

	switch timeUnitType {
	case TimeUnitMinutes:
		*e = append(*e, Minute(intVal))
		break
	case TimeUnitHours:
		*e = append(*e, Hour(intVal))
		break
	case TimeUnitDays:
		*e = append(*e, Day(intVal))
		break
	case TimeUnitMonths:
		*e = append(*e, Month(intVal))
		break
	case TimeUnitDaysOfWeek:
		*e = append(*e, DayOfWeek(intVal))
		break
	default:
		return errors.New("unknown unit type")
	}

	return
}

func (e *ValueSet) appendRange(valueExpression ValueExpression, timeUnitType TimeUnitType) (err error) {
	rangeBoundaries := strings.Split(valueExpression.ToString(), "-")

	min, err := strconv.Atoi(rangeBoundaries[0])

	if err != nil {
		return errors.New("invalid value at list start")
	}

	max, err := strconv.Atoi(rangeBoundaries[1])

	if err != nil {
		return errors.New("invalid value at list end")
	}

	a := make([]int, max-min+1)

	switch timeUnitType {
	case TimeUnitMinutes:
		for i := range a {
			*e = append(*e, Minute(min+i))
		}
		break
	case TimeUnitHours:
		for i := range a {
			*e = append(*e, Hour(min+i))
		}
		break
	case TimeUnitDays:
		for i := range a {
			*e = append(*e, Day(min+i))
		}
		break
	case TimeUnitMonths:
		for i := range a {
			*e = append(*e, Month(min+i))
		}
		break
	case TimeUnitDaysOfWeek:
		for i := range a {
			*e = append(*e, DayOfWeek(min+i))
		}
		break
	default:
		return errors.New("unknown unit type")
	}

	return
}

func (e *ValueSet) appendInterval(valueExpression ValueExpression, timeUnitType TimeUnitType) (err error) {
	operands := strings.Split(valueExpression.ToString(), "/")

	var offset = 0

	if len(operands) < 2 {
		operands = append([]string{"0"}, operands...)
	}

	if operands[0] == "*" {
		offset = 0
	} else {
		offset, err = strconv.Atoi(operands[0])
	}

	if err != nil {
		return errors.New("invalid value at list start")
	}

	interval, err := strconv.Atoi(operands[1])

	if err != nil {
		return errors.New("invalid value at list end")
	}

	switch timeUnitType {
	case TimeUnitMinutes:
		for i := Minutes[0].ToInt(); (i + offset) < Minutes[len(Minutes)-1].ToInt(); i += interval {
			*e = append(*e, Minute(i+offset))
		}
		break
	case TimeUnitHours:
		for i := Hours[0].ToInt(); (i + offset) < Hours[len(Hours)-1].ToInt(); i += interval {
			*e = append(*e, Hour(i+offset))
		}
		break
	case TimeUnitDays:
		for i := Days[0].ToInt(); (i + offset) < Days[len(Days)-1].ToInt(); i += interval {
			*e = append(*e, Day(i+offset))
		}
		break
	case TimeUnitMonths:
		for i := Months[0].ToInt(); (i + offset) < Months[len(Months)-1].ToInt(); i += interval {
			*e = append(*e, Month(i+offset))
		}
		break
	case TimeUnitDaysOfWeek:
		for i := DaysOfWeek[0].ToInt(); (i + offset) <= DaysOfWeek[len(DaysOfWeek)-1].ToInt(); i += interval {
			*e = append(*e, DayOfWeek(i+offset))
		}
	default:
		err = errors.New("unknown unit type")
	}

	return
}

func (e *ValueSet) appendList(timeSpecPart ValueExpression, timeUnitType TimeUnitType) (err error) {
	items := strings.Split(timeSpecPart.ToString(), ",")

	for i := range items {
		values, err := ValueExpression(items[i]).Expand(timeUnitType)
		*e = append(*e, values...)

		if err != nil {
			return err
		}
	}

	return
}

type ExpandValues func(timeSpecPart ValueExpression, timeUnitType TimeUnitType) (values ValueSet, err error)

func (v ValueExpression) Expand(timeUnitType TimeUnitType) (values ValueSet, err error) {
	switch {
	case v.IsWildCard():
		err = values.appendAll(timeUnitType)
		break
	case v.IsSimple():
		err = values.appendSimple(v, timeUnitType)
		break
	case v.IsRange():
		err = values.appendRange(v, timeUnitType)
		break
	case v.IsInterval():
		err = values.appendInterval(v, timeUnitType)
		break
	case v.IsList():
		err = values.appendList(v, timeUnitType)
		break
	default:
		err = errors.New("invalid v " + v.ToString())
	}

	return values, err
}
