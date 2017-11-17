package specparser

import (
	"specparser"
	"testing"
)

func TestMinute_ToInt(t *testing.T) {
	var sample = specparser.Minute(23)

	if sample.ToInt() != 23 {
		t.Error("Error converting minute to int")
	}
}

func TestHour_ToInt(t *testing.T) {
	var sample = specparser.Hour(23)

	if sample.ToInt() != 23 {
		t.Error("Error converting hour to int")
	}
}

func TestDay_ToInt(t *testing.T) {
	var sample = specparser.Day(23)

	if sample.ToInt() != 23 {
		t.Error("Error converting day to int")
	}
}

func TestMonth_ToInt(t *testing.T) {
	var sample = specparser.Month(23)

	if sample.ToInt() != 23 {
		t.Error("Error converting month to int")
	}
}

func TestDayOfWeek_ToInt(t *testing.T) {
	var sample = specparser.DayOfWeek(23)

	if sample.ToInt() != 23 {
		t.Error("Error converting dayOfWeek to int")
	}
}

//func TestExplode()  {
//	var sample = specparser.TimeExpression{
//		specparser.ValueExpression("1-3"),
//		specparser.ValueExpression("4,6,8"),
//		specparser.ValueExpression("*/2"),
//		specparser.ValueExpression("4-11"),
//		specparser.ValueExpression("*"),
//	}
//}

func TestTimeExpression_New(t *testing.T) {
	var expectedMinute = "1-5"
	var expectedHour = "2-4"
	var expectedDay = "1,3,5"
	var expectedMonth = "*/3"
	var expectedDayOfWeek = "*"

	sample := specparser.TimeExpression{}.New(
		expectedMinute,
		expectedHour,
		expectedDay,
		expectedMonth,
		expectedDayOfWeek,
	)

	if sample.Minute.ToString() != expectedMinute {
		t.Fail()
	}

	if sample.Hour.ToString() != expectedHour {
		t.Fail()
	}

	if sample.Day.ToString() != expectedDay {
		t.Fail()
	}

	if sample.Month.ToString() != expectedMonth {
		t.Fail()
	}

	if sample.DayOfWeek.ToString() != expectedDayOfWeek {
		t.Fail()
	}
}

func TestTimeExpression_Explode(t *testing.T) {
	var expectedMinute = "1-5"
	var expectedHour = "2-4"
	var expectedDay = "1,3,5"
	var expectedMonth = "*/3"
	var expectedDayOfWeek = "*"

	sample, err := specparser.TimeExpression{
		Minute:    specparser.ValueExpression(expectedMinute),
		Hour:      specparser.ValueExpression(expectedHour),
		Day:       specparser.ValueExpression(expectedDay),
		Month:     specparser.ValueExpression(expectedMonth),
		DayOfWeek: specparser.ValueExpression(expectedDayOfWeek),
	}.Explode()

	if err != nil {
		t.Error(err)
	}

	switch {
	case len(sample.Minutes) != 5:
	case sample.Minutes[0] != specparser.Minute(1):
	case sample.Minutes[1] != specparser.Minute(2):
	case sample.Minutes[2] != specparser.Minute(3):
	case sample.Minutes[3] != specparser.Minute(4):
	case sample.Minutes[4] != specparser.Minute(5):
	case len(sample.Hours) != 3:
	case sample.Hours[0] != specparser.Hour(1):
	case sample.Hours[1] != specparser.Hour(3):
	case sample.Hours[2] != specparser.Hour(4):
	case len(sample.Days) != 3:
	case sample.Days[0] != specparser.Day(1):
	case sample.Days[1] != specparser.Day(3):
	case sample.Days[2] != specparser.Day(5):
	case len(sample.Months) != 4:
	case sample.Months[0] != specparser.Month(1):
	case sample.Months[1] != specparser.Month(4):
	case sample.Months[2] != specparser.Month(7):
	case sample.Months[3] != specparser.Month(10):
	case len(sample.DaysOfWeek) != 7:
	case sample.DaysOfWeek[0] != specparser.DayOfWeek(1):
	case sample.DaysOfWeek[1] != specparser.DayOfWeek(2):
	case sample.DaysOfWeek[2] != specparser.DayOfWeek(3):
	case sample.DaysOfWeek[3] != specparser.DayOfWeek(4):
	case sample.DaysOfWeek[4] != specparser.DayOfWeek(5):
	case sample.DaysOfWeek[5] != specparser.DayOfWeek(6):
	case sample.DaysOfWeek[6] != specparser.DayOfWeek(7):
		t.Error(err)
		break
	}

}
