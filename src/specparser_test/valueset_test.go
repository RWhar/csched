package specparser_test

import (
	"errors"
	"specparser"
	"strconv"
	"testing"
)

func TestValueExpression_IsInterval(t *testing.T) {
	var sample = specparser.ValueExpression("*/3")

	if !sample.IsInterval() {
		t.Error("Error identifying interval")
	}
}

func TestValueExpression_IsWildCard(t *testing.T) {
	var sample = specparser.ValueExpression("*")

	if !sample.IsWildCard() {
		t.Error("Error identifying wildcard")
	}
}

func TestValueExpression_IsRange(t *testing.T) {
	var sample = specparser.ValueExpression("1-10")

	if !sample.IsRange() {
		t.Error("Error identifying range")
	}
}

func TestValueExpression_IsList(t *testing.T) {
	var sample = specparser.ValueExpression("7,11,22")

	if !sample.IsList() {
		t.Error("Error identifying interval")
	}
}

func TestValueExpression_IsSimple(t *testing.T) {
	var sample = specparser.ValueExpression("3")

	if !sample.IsSimple() {
		t.Error("Error identifying interval")
	}
}

func TestValueExpandWildcardMinutes(t *testing.T) {
	testValueExpandAll(t, specparser.TimeUnitMinutes, 0, 59)
}

func TestValueExpandWildcardHours(t *testing.T) {
	testValueExpandAll(t, specparser.TimeUnitHours, 0, 23)
}

func TestValueExpandWildcardDays(t *testing.T) {
	testValueExpandAll(t, specparser.TimeUnitDays, 1, 31)
}

func TestValueExpandWildcardMonths(t *testing.T) {
	testValueExpandAll(t, specparser.TimeUnitMonths, 1, 12)
}

func TestValueExpandWildcardDaysOfWeek(t *testing.T) {
	testValueExpandAll(t, specparser.TimeUnitDaysOfWeek, 1, 7)
}

func TestValueExpandWildcardUnknownType(t *testing.T) {
	testValueExpandAll(t, 22, 1, 7)
}

func testValueExpandAll(t *testing.T, unitType specparser.TimeUnitType, first int, last int) {
	valueExpression := specparser.ValueExpression("*")

	var expectError bool
	var expectedValue specparser.TimeUnit

	_, err := typeCast(1, unitType) // Set expectError if required

	if err != nil {
		expectError = true
	}

	for i := 0; i <= (last - first); i++ {
		expectedValue, _ = typeCast(i+first, unitType)

		values, err := valueExpression.Expand(unitType)

		if err != nil {
			if !expectError {
				t.Error("Error while expanding wildcard", err)
			}
			return
		} else if expectError {
			t.Error("Expecting error but none found")
			return
		}

		if len(values) < i || values[i] != expectedValue {
			t.Error("value collection missing expected value:", expectedValue, "actual value:", values[i])
		}
	}
}

func TestValueExpandRangeMinutes(t *testing.T) {
	testExpandRange(t, specparser.TimeUnitMinutes, "2", "23", false)
}

func TestValueExpandRangeHours(t *testing.T) {
	testExpandRange(t, specparser.TimeUnitHours, "9", "17", false)
}

func TestValueExpandRangeDays(t *testing.T) {
	testExpandRange(t, specparser.TimeUnitDays, "8", "29", false)
}

func TestValueExpandRangeMonths(t *testing.T) {
	testExpandRange(t, specparser.TimeUnitMonths, "6", "12", false)
}

func TestValueExpandRangeDaysOfWeek(t *testing.T) {
	testExpandRange(t, specparser.TimeUnitDaysOfWeek, "1", "5", false)
}

func TestValueExpandRangeUnknownType(t *testing.T) {
	testExpandRange(t, 77, "1", "5", true)
}

func TestValueExpandRangeUnknownListStart(t *testing.T) {
	values, err := specparser.ValueExpression("a-9").Expand(specparser.TimeUnitMinutes)

	if err == nil {
		t.Error("expecting error non found", values)
	}
}

func TestValueExpandRangeUnknownListEnd(t *testing.T) {
	values, err := specparser.ValueExpression("1-z").Expand(specparser.TimeUnitMinutes)

	if err == nil {
		t.Error("expecting error non found", values)
	}
}

func testExpandRange(t *testing.T, unitType specparser.TimeUnitType, first string, last string, expectError bool) {
	valueExpression := specparser.ValueExpression(first + "-" + last)

	var expectedValue specparser.TimeUnit

	_, err := typeCast(1, unitType) // Set expectError if required

	if err != nil {
		expectError = true
	}

	firstInt, err := strconv.Atoi(first)
	lastInt, err := strconv.Atoi(last)

	if err != nil {
		t.Error("error setting up test", err)
		return
	}

	for i := 0; i <= (lastInt - firstInt); i++ {
		expectedValue, _ = typeCast(i+firstInt, unitType)

		values, err := valueExpression.Expand(unitType)

		if err != nil {
			if !expectError {
				t.Error("Error while expanding range", valueExpression)
			}
			return
		} else if expectError {
			t.Error("Expecting error but none found")
			return
		}

		if values[i] != expectedValue {
			t.Error("value collection missing expected value:", expectedValue, "actual value:", values[i])
		}
	}
}

func TestValueExpandListMinutes(t *testing.T) {
	valueExpression := specparser.ValueExpression("3,8,12,14")

	var expectedValues []int
	expectedValues = append(expectedValues, 3)
	expectedValues = append(expectedValues, 8)
	expectedValues = append(expectedValues, 12)
	expectedValues = append(expectedValues, 14)

	testValueExpandList(t, valueExpression, specparser.TimeUnitMinutes, expectedValues)
}

func TestValueExpandListHours(t *testing.T) {
	valueExpression := specparser.ValueExpression("3,8,12,14")

	var expectedValues []int
	expectedValues = append(expectedValues, 3)
	expectedValues = append(expectedValues, 8)
	expectedValues = append(expectedValues, 12)
	expectedValues = append(expectedValues, 14)

	testValueExpandList(t, valueExpression, specparser.TimeUnitHours, expectedValues)
}

func TestValueExpandListDays(t *testing.T) {
	valueExpression := specparser.ValueExpression("3,8,12,14")
	var expectedValues []int

	expectedValues = append(expectedValues, 3)
	expectedValues = append(expectedValues, 8)
	expectedValues = append(expectedValues, 12)
	expectedValues = append(expectedValues, 14)

	testValueExpandList(t, valueExpression, specparser.TimeUnitDays, expectedValues)
}

func TestValueExpandListMonths(t *testing.T) {
	valueExpression := specparser.ValueExpression("3,8,12,14")
	var expectedValues []int

	expectedValues = append(expectedValues, 3)
	expectedValues = append(expectedValues, 8)
	expectedValues = append(expectedValues, 12)
	expectedValues = append(expectedValues, 14)

	testValueExpandList(t, valueExpression, specparser.TimeUnitMonths, expectedValues)
}

func TestValueExpandListDaysOfWeek(t *testing.T) {
	valueExpression := specparser.ValueExpression("3,8,12,14")

	var expectedValues []int
	expectedValues = append(expectedValues, 3)
	expectedValues = append(expectedValues, 8)
	expectedValues = append(expectedValues, 12)
	expectedValues = append(expectedValues, 14)

	testValueExpandList(t, valueExpression, specparser.TimeUnitDaysOfWeek, expectedValues)
}

func TestValueExpandListUnknownType(t *testing.T) {
	valueExpression := specparser.ValueExpression("3,8,12,14")

	var expectedValues []int
	expectedValues = append(expectedValues, 3)
	expectedValues = append(expectedValues, 8)
	expectedValues = append(expectedValues, 12)
	expectedValues = append(expectedValues, 14)

	testValueExpandList(t, valueExpression, 77, expectedValues)
}

func testValueExpandList(t *testing.T, valueExpression specparser.ValueExpression, unitType specparser.TimeUnitType, expectedValues []int) {
	var expectError bool
	var expectedValue specparser.TimeUnit

	_, err := typeCast(1, unitType) // Set expectError if required

	if err != nil {
		expectError = true
	}

	values, err := valueExpression.Expand(unitType)

	if err != nil {
		if !expectError {
			t.Error("Error while expanding list", valueExpression)
		}
		return
	} else if expectError {
		t.Error("expected error but none found")
		return
	}

	for i := 0; i < len(expectedValues); i++ {
		expectedValue, _ = typeCast(expectedValues[i], unitType)

		if values[i] != expectedValue {
			t.Error("value collection missing expected value:", expectedValue, "actual value:", values[i])
		}
	}
}

func TestValueExpandSimpleMinute(t *testing.T) {
	testValueExpandSimple(t, specparser.TimeUnitMinutes, 53)
}

func TestValueExpandSimpleHour(t *testing.T) {
	testValueExpandSimple(t, specparser.TimeUnitHours, 17)
}

func TestValueExpandSimpleDay(t *testing.T) {
	testValueExpandSimple(t, specparser.TimeUnitDays, 29)
}

func TestValueExpandSimpleMonth(t *testing.T) {
	testValueExpandSimple(t, specparser.TimeUnitMonths, 11)
}

func TestValueExpandSimpleDayOfWeek(t *testing.T) {
	testValueExpandSimple(t, specparser.TimeUnitDaysOfWeek, 5)
}

func TestValueExpandSimpleUnknownType(t *testing.T) {
	testValueExpandSimple(t, 77, 5)
}

func testValueExpandSimple(t *testing.T, unitType specparser.TimeUnitType, value int) {
	valueExpression := specparser.ValueExpression(strconv.Itoa(value))

	var expectError bool
	var expectedValue specparser.TimeUnit

	expectedValue, err := typeCast(value, unitType)

	if err != nil {
		expectError = true
	}

	values, err := valueExpression.Expand(unitType)

	if err != nil {
		if !expectError {
			t.Error("Error while expanding list", valueExpression)
		}
		return
	} else if expectError {
		t.Error("expected error none found")
		return
	}

	if values[0] != expectedValue {
		t.Error("value collection missing expected value:", expectedValue, "actual value:", values[0])
	}
}

func TestValueExpandIntervalMinutes(t *testing.T) {
	valueExpression := specparser.ValueExpression("1/5")
	var expectedValues []int

	expectedValues = append(expectedValues, 1)
	expectedValues = append(expectedValues, 6)
	expectedValues = append(expectedValues, 11)
	expectedValues = append(expectedValues, 16)
	expectedValues = append(expectedValues, 21)
	expectedValues = append(expectedValues, 26)
	expectedValues = append(expectedValues, 31)
	expectedValues = append(expectedValues, 36)
	expectedValues = append(expectedValues, 41)
	expectedValues = append(expectedValues, 46)
	expectedValues = append(expectedValues, 51)
	expectedValues = append(expectedValues, 56)

	testValueExpandInterval(t, valueExpression, specparser.TimeUnitMinutes, expectedValues)
}

func TestValueExpandIntervalHours(t *testing.T) {
	valueExpression := specparser.ValueExpression("*/2")
	var expectedValues []int

	expectedValues = append(expectedValues, 0)
	expectedValues = append(expectedValues, 2)
	expectedValues = append(expectedValues, 4)
	expectedValues = append(expectedValues, 6)
	expectedValues = append(expectedValues, 8)
	expectedValues = append(expectedValues, 10)
	expectedValues = append(expectedValues, 12)
	expectedValues = append(expectedValues, 14)
	expectedValues = append(expectedValues, 16)
	expectedValues = append(expectedValues, 18)
	expectedValues = append(expectedValues, 20)
	expectedValues = append(expectedValues, 22)

	testValueExpandInterval(t, valueExpression, specparser.TimeUnitHours, expectedValues)
}

func TestValueExpandIntervalDays(t *testing.T) {
	valueExpression := specparser.ValueExpression("*/14")
	var expectedValues []int

	expectedValues = append(expectedValues, 1)
	expectedValues = append(expectedValues, 15)
	expectedValues = append(expectedValues, 29)

	testValueExpandInterval(t, valueExpression, specparser.TimeUnitDays, expectedValues)
}

func TestValueExpandIntervalMonths(t *testing.T) {
	valueExpression := specparser.ValueExpression("*/3")
	var expectedValues []int

	expectedValues = append(expectedValues, 1)
	expectedValues = append(expectedValues, 4)
	expectedValues = append(expectedValues, 7)
	expectedValues = append(expectedValues, 10)

	testValueExpandInterval(t, valueExpression, specparser.TimeUnitMonths, expectedValues)
}

func TestValueExpandIntervalDaysOfWeek(t *testing.T) {
	valueExpression := specparser.ValueExpression("*/2")
	var expectedValues []int

	expectedValues = append(expectedValues, 1)
	expectedValues = append(expectedValues, 3)
	expectedValues = append(expectedValues, 5)
	expectedValues = append(expectedValues, 7)

	testValueExpandInterval(t, valueExpression, specparser.TimeUnitDaysOfWeek, expectedValues)
}

func TestValueExpandIntervalUnknownType(t *testing.T) {
	valueExpression := specparser.ValueExpression("*/3")
	var expectedValues []int

	expectedValues = append(expectedValues, 1)
	expectedValues = append(expectedValues, 4)
	expectedValues = append(expectedValues, 7)
	expectedValues = append(expectedValues, 10)

	testValueExpandInterval(t, valueExpression, 77, expectedValues)
}

func testValueExpandInterval(t *testing.T, valueExpression specparser.ValueExpression, unitType specparser.TimeUnitType, expectedValues []int) {
	var expectError bool
	var expectedValue specparser.TimeUnit

	_, err := typeCast(1, unitType) // Set expectError if required

	if err != nil {
		expectError = true
	}

	values, err := valueExpression.Expand(unitType)

	if err != nil {
		if !expectError {
			t.Error("Error while expanding interval", valueExpression, err)
		}
		return
	} else if expectError {
		t.Error("expected error non found")
		return
	}

	if len(expectedValues) != len(values) {
		t.Error("expected value count does not match actual value count", expectedValues, values)
		return
	}

	for i := 0; i < len(expectedValues); i++ {
		expectedValue, _ = typeCast(expectedValues[i], unitType)

		if values[i] != expectedValue {
			t.Error("value collection missing expected value:", expectedValue, "actual value:", values[i])
		}
	}
}

func typeCast(input int, unitType specparser.TimeUnitType) (value specparser.TimeUnit, err error) {
	switch unitType {
	case specparser.TimeUnitMinutes:
		value = specparser.Minute(input)
		break
	case specparser.TimeUnitHours:
		value = specparser.Hour(input)
		break
	case specparser.TimeUnitDays:
		value = specparser.Day(input)
		break
	case specparser.TimeUnitMonths:
		value = specparser.Month(input)
		break
	case specparser.TimeUnitDaysOfWeek:
		value = specparser.DayOfWeek(input)
		break
	default:
		err = errors.New("unknown unit type")
	}

	return
}
