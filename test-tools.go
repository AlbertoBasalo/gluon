package main

import (
	"testing"
)

/*
 Utility function to check if the two values are equal
 Sends an erro if not, makeing the test fail
*/
func AssertEqual(t testing.TB, actual, expected string) {
	t.Helper() // marks this funcion as a helper to skip it from the stack trace
	if actual != expected {
		// %q format with quoted strings
		t.Errorf("❌ Actual: %q expected: %q", actual, expected)
	}
}
