package gt

import (
	"time"
)

func GreaterThan[T orderable](actual T, expected T) bool {
	return actual > expected
}

func GreaterThanOrEq[T orderable](actual T, expected T) bool {
	return actual >= expected
}

func LessThan[T orderable](actual T, expected T) bool {
	return actual < expected
}

func LessThanOrEq[T orderable](actual T, expected T) bool {
	return actual <= expected
}

// minに指定した数値と、expectedに指定した数値は範囲に含まれます
func Between[T orderable](min T) func(T, T) bool {
	return func(actual, max T) bool {
		return GreaterThanOrEq(actual, min) && LessThanOrEq(actual, max)
	}
}

func After(actual time.Time, expected time.Time) bool {
	return actual.After(expected)
}

func AfterOrEq(actual time.Time, expected time.Time) bool {
	return actual.After(expected) || actual.Equal(expected)
}

func Before(actual time.Time, expected time.Time) bool {
	return actual.Before(expected)
}

func BeforeOrEq(actual time.Time, expected time.Time) bool {
	return actual.Before(expected) || actual.Equal(expected)
}

// minに指定した数値と、expectedに指定した数値は範囲に含まれます
func TimeBetween[T time.Time](from time.Time) func(time.Time, time.Time) bool {
	return func(actual, to time.Time) bool {
		return AfterOrEq(actual, from) && BeforeOrEq(actual, to)
	}
}
