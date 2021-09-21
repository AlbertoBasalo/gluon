package main

import "testing"

func AssertEqual(t testing.TB, actual, expected string) {
	t.Helper()
	if actual != expected {
		t.Errorf("❌ Actual: %q expected: %q", actual, expected)
	}
	// else {
	// 	t.Log("✅ ", t.Name())
	// }
}
