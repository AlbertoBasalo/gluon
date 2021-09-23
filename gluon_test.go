package main

import "testing" // imports testing core utilities

/*
	Tests are funcions with an special argument "t"
	that is used to run and report test failures.
	Test belong to the same package as the code they test.
*/
func TestGetStartLine(t *testing.T) {
	// runs the test
	t.Run("it should create the start message", func(t *testing.T) {
		// arrange
		input := "test.go"
		// act
		actual := GetStartLine(input)
		// assert
		expected := "🚀 test.go started"
		AssertEqual(t, expected, actual) // uses a custom function
	})
}
