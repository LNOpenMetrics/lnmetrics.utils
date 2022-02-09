// Utils Time package, contains a sequence of time check
// useful to compare two dates.
package utime

import (
	"math"
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

// Return true if the UNIX timestamp it is inside a range [start, start + duration]
// target: The timestamp that where we want make the check.
// start: The starting point in UNIX time where we want perform the check.
// offset: The time duration of the period expressed in time.Duration type.
func InRangeFromUnix(target int64, start int64, offset time.Duration) bool {
	end := time.Unix(start, 0).Add(offset)
	return target >= start && target <= end.Unix()
}

// Return true if the UNIX timestamp it is inside a range [start, start + months]
// target: The timestamp that where we want make the check.
// start: The starting point in UNIX time where we want perform the check.
// offset: The time duration of the period expressed in time.Duration type.
func InRangeMonthsUnix(target int64, start int64, months int64) bool {
	end := time.Unix(start, 0).Add(time.Duration(months) * Month)
	return target >= start && target <= end.Unix()
}

// Convert decimal unix timestamp in a int64 unix timestamp
func FromDecimalUnix(decimalUnix float64) int64 {
	// Source: https://stackoverflow.com/a/37628278/10854225
	sec, dec := math.Modf(decimalUnix)
	return time.Unix(int64(sec), int64(dec*(1e9))).Unix()
}

// Count the number of occurrence that are inside the following
func OccurenceInUnixRange(from int64, to int64, step time.Duration) uint64 {
	count := uint64(0)
	actual := from
	for actual < to {
		actual = increaseTimestampByDuration(actual, step)
		count++
	}
	return count
}

// Add a specified duration to the timestamp provided.
func AddToTimestamp(timestamp int64, duration time.Duration) int64 {
	return increaseTimestampByDuration(timestamp, duration)
}

// Subtract a specified duration to the timestamp provided.
func SubToTimestamp(timestamp int64, duration time.Duration) int64 {
	return decreaseTimestampByDuration(timestamp, duration)
}

// Increase the timestamp received by a specified duration.
func increaseTimestampByDuration(timestamp int64, duration time.Duration) int64 {
	actualTime := time.Unix(timestamp, 0).Add(duration)
	return actualTime.Unix()
}

// Decrease the timestamp provided by the specified duration.
func decreaseTimestampByDuration(timestamp int64, duration time.Duration) int64 {
	actualTime := time.Unix(timestamp, 0).Add(time.Duration(-duration))
	return actualTime.Unix()
}
