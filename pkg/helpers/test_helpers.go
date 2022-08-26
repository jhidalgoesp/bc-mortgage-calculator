package helpers

import (
	"errors"
	"testing"
)

// AssertSameFloat asserts two int are equal
func AssertSameFloat(t testing.TB, got, want float64) {
	t.Helper()

	if got != want {
		t.Errorf("got %f want %f", got, want)
	}
}

// AssertSameInt asserts two int are equal
func AssertSameInt(t testing.TB, got, want int) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

// AssertNilError asserts an error is nil
func AssertNilError(t testing.TB, val interface{}) {
	t.Helper()

	if val != nil {
		t.Errorf("Expected nil, got: %s.", val)
	}
}

// AssertEqualErrors asserts two errors are equal
func AssertEqualErrors(t testing.TB, got, want error) {
	if !errors.Is(got, want) {
		t.Errorf("Expected %v, got %v", got, want)
	}
}
