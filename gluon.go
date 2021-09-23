package main // name of this package

import "fmt" // importing external packages

var program string // declare a variable with type string

func main() {
	program = "gluon"             // assign a value to the variable
	line := GetStartLine(program) // initialize and assing a value
	fmt.Println(line)             // print the value of the variable
}

/*
   Function that returns a formated string
   needs input and output types to be declared
*/
func GetStartLine(program string) string {
	return "🚀 " + program + " started"
}
