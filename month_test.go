package month

import "testing"

func TestString(t *testing.T) {

	m := Month(January)
	monthName := m.String()

	if monthName != "January" {
		t.Fatal(
			"For", "1",
			"expected", "January",
			"got", monthName)
	}
}

func TestLastDay(t *testing.T) {

	lastDays := [12]int{
		31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31,
	}

	// Normal years

	for i := 1; i <= 12; i++ {

		m := Month(i)
		actual := m.LastDay(2015)
		expected := lastDays[i-1]

		if actual != expected {
			t.Error(
				"For", m,
				"expected", expected,
				"got", actual)
		}
	}

	// Leap year

	actual := Month(February).LastDay(2012)

	if actual != 29 {
		t.Error(
			"For leap year",
			"expected", 29,
			"got", actual)
	}
}
