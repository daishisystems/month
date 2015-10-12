// Package month provides functionality, pertaining to months of the year, that
// is not available in standard Go libraries. Examples of this include the
// ability to return the last day of any given month.
package month

// Month represents a standard calendar month, applicable to integer format.
type Month int

// Standard calendar months.
const (
	January Month = 1 + iota
	February
	March
	April
	May
	June
	July
	August
	September
	October
	November
	December
)

var monthNames = [12]string{
	"January",
	"February",
	"March",
	"April",
	"May",
	"June",
	"July",
	"August",
	"September",
	"October",
	"November",
	"December",
}

var monthDays = map[int]int{
	1:  31,
	3:  31,
	4:  30,
	5:  31,
	6:  30,
	7:  31,
	8:  31,
	9:  30,
	10: 31,
	11: 30,
	12: 31,
}

// String returns an English language based representation of m. E.g., January,
// February, etc.
func (m Month) String() string {

	return monthNames[m-1]
}

// LastDay returns the last day of m, based on year. The function takes into
// account leap years.
func (m Month) LastDay(year uint16) int {

	if m == 2 {

		isLeapYear := year%4 == 0 &&
			(year%100 != 0 || year%400 == 0)

		if isLeapYear {
			return 29
		}
		return 28
	}
	return monthDays[int(m)]
}
