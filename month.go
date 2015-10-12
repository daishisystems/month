// Package month provides functionality, pertaining to months of the year, that
// is not available in standard Go libraries. Examples of this include the
// ability to return the last day of any given month.
package month

import "time"

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

// String returns an English language based representation of m. E.g., January,
// February, etc.
func (m Month) String() string {

	return monthNames[m-1]
}

// LastDay returns the last day of m, based on year. The function takes into
// account leap years.
func (m Month) LastDay(year uint16) int {
	ts := time.Date(int(year), time.Month(m)+1, 0, 0, 0, 0, 0, time.UTC)
	return ts.Day()
}
