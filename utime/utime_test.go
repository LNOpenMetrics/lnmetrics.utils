package utime

import (
	"testing"
	"time"
)

func TestSameDayUnixOne(t *testing.T) {
	first := time.Now().Unix()
	second := time.Now().Unix()

	if !SameDayUnix(first, second) {
		t.Errorf("The two date are different, but them should be equal")
	}
}

func TestSameMonthUnixOne(t *testing.T) {
	first := time.Now().Unix()
	second := time.Now().Unix()

	if !SameMonthUnix(first, second) {
		t.Errorf("The two date are in different month, but them should be equal")
	}
}

func TestSameInRangeUnixOne(t *testing.T) {
	start := time.Now().Unix()
	target := time.Now().Add(23 * time.Hour).Unix()

	if !InRangeFromUnix(target, start, 24*time.Hour) {
		t.Errorf("The two dates are out of the 1 day range, them should be inside this range")
	}
}

func TestSameInRangeMonthsUnixOne(t *testing.T) {
	start := time.Now().Unix()
	target := time.Now().Add(1 * Month).Unix()

	if !InRangeMonthsUnix(target, start, 1) {
		t.Errorf("The two dates are out of the 1 month range, them should be inside this range")
	}
}

func TestSameInRangeMonthsUnixTwo(t *testing.T) {
	start := time.Now().Unix()
	target := time.Now().Add(2 * Month).Unix()

	if InRangeMonthsUnix(target, start, 1) {
		t.Errorf("The two dates are inside the range of 1 month, them should be outside of this range")
	}
}

func TestCountOccurenceInOneHour(t *testing.T) {
	start := time.Now().Unix()
	end := time.Now().Add(1 * time.Hour).Unix()

	occur := OccurenceInUnixRange(start, end, 30*time.Minute)

	if occur != 2 {
		t.Errorf("Expected 2 by received %d", occur)
	}
}

func TestCountOccurenceInOneDay(t *testing.T) {
	start := time.Now().Unix()
	end := time.Now().Add(24 * time.Hour).Unix()

	occur := OccurenceInUnixRange(start, end, 30*time.Minute)

	if occur != 2*24 {
		t.Errorf("Expected %d but received %d", 2*24, occur)
	}
}

func TestSubtractDay(t *testing.T) {
	target := time.Now().Unix()
	targetDay := time.Now().Day()

	yesterday := SubToTimestamp(target, 1*24*time.Hour)
	yesterdayDay := time.Unix(yesterday, 0).Day()
	if targetDay-1 != yesterdayDay {
		t.Errorf("The target day %d is different from the day calculated %d", targetDay, yesterdayDay)
	}
}

func TestSubtractMinute(t *testing.T) {
	target := time.Now().Unix()
	targetDay := time.Now().Day()

	yesterday := SubToTimestamp(target, 24*time.Minute)
	yesterdayDay := time.Unix(yesterday, 0).Day()
	if targetDay != yesterdayDay {
		t.Errorf("The target day %d is different from the day calculated %d", targetDay, yesterdayDay)
	}
}
