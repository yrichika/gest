package gt

import (
	"time"
)

func GreaterThan[T orderable](expected T) func(T) bool {
	return func(actual T) bool {
		return actual > expected
	}
}

func GreaterThanOrEq[T orderable](expected T) func(T) bool {
	return func(actual T) bool {
		return actual >= expected
	}
}
func LessThan[T orderable](expected T) func(T) bool {
	return func(actual T) bool {
		return actual < expected
	}
}

func LessThanOrEq[T orderable](expected T) func(T) bool {
	return func(actual T) bool {
		return actual <= expected
	}
}

// minに指定した数値と、expectedに指定した数値は範囲に含まれます
func Between[T orderable](min T, max T) func(T) bool {
	return func(actual T) bool {
		return GreaterThanOrEq(min)(actual) && LessThanOrEq(max)(actual)
	}
}

func After(expected time.Time) func(time.Time) bool {
	return func(actual time.Time) bool {
		return actual.After(expected)
	}
}

func AfterOrEq(expected time.Time) func(time.Time) bool {
	return func(actual time.Time) bool {
		return actual.After(expected) || actual.Equal(expected)
	}
}

func Before(expected time.Time) func(time.Time) bool {
	return func(actual time.Time) bool {
		return actual.Before(expected)
	}
}

func BeforeOrEq(expected time.Time) func(time.Time) bool {
	return func(actual time.Time) bool {
		return actual.Before(expected) || actual.Equal(expected)
	}
}

// minに指定した数値と、expectedに指定した数値は範囲に含まれます
func TimeBetween[T time.Time](from time.Time, to time.Time) func(time.Time) bool {
	return func(actual time.Time) bool {
		return AfterOrEq(from)(actual) && BeforeOrEq(to)(actual)
	}
}
