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
