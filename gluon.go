package main

import "fmt"

func main() {
	program := "gluon"
	fmt.Println(GetStartLine(program))
}

func GetStartLine(program string) string {
	return "🚀 " + program + " started"
}
