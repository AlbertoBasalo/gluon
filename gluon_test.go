package main

import "testing"

func TestGetStartLine(t *testing.T) {
	t.Run("it should create the start message", func(t *testing.T) {
		actual := GetStartLine("test.go")
		expected := "🚀 test.go started"
		AssertEqual(t, expected, actual)
	})
}
