// Utils Time package, contains a sequence of time check
// useful to compare two dates.
package utime

import (
	"time"
)

const (
	Month = 730 * time.Hour
	Year  = 12 * Month
)

// Return true if the first and second unix timestamp defined in seconds
// are inside the same day
func SameDayUnix(first int64, second int64) bool {
	firstTime := time.Unix(int64(first), 0)
	secondTime := time.Unix(int64(second), 0)
	return firstTime.Day() == secondTime.Day()
}

// Return true if the first and second unix timestamp definied in seconds
// are inside the same month
func SameMonthUnix(first int64, second int64) bool {
	firstTime := time.Unix(first, 0)
	secondTime := time.Unix(second, 0)
	return firstTime.Month() == secondTime.Month()
}

// Return true if the first and second unix timestamp defined in seconds
// are inside the same year.
func SameYearUnix(first int64, second int64) bool {
	firstTime := time.Unix(first, 0)
	secondTime := time.Unix(second, 0)
	return firstTime.Year() == secondTime.Year()
}

// Return true if the unix timestamp it is inside a range of
// [second, range] or [range, second]
func InRangeFromUnix(target int64, start int64, offset time.Duration) bool {
	end := time.Unix(start, 0).Add(offset)
	return target >= start && target <= end.Unix()
}

func InRangeMonthsUnix(target int64, start int64, months int64) bool {
	end := time.Unix(start, 0).Add(time.Duration(months) * Month)
	return target >= start && target <= end.Unix()
}
