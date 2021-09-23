# 0 - Hello world

## 0.1 Install tools

- [Official download site](https://golang.org/doc/install)

- [VS code Extension](https://marketplace.visualstudio.com/items?itemName=golang.go)

- [More VS Code info](https://code.visualstudio.com/docs/languages/go)


## 0.2 Hello world

```bash
mkdir gluon
cd gluon
 # generate a module for your program
go mod init atomic/gluon
code .
```

### Show me the code

On `gluon.go` file (View code)[https://github.com/AtomicBuilders/gluon/blob/main/gluon.go]

```go
package main // name of this package

import "fmt" // importing external packages

var program string // declare a variable with type string

func main() {
	program = "gluon"             // assign a value to the variable
	line := GetStartLine(program) // initialize and assing a value
	fmt.Println(line)             // print the value of the variable
}

/*
	GetStartLine returns the starting line of the program
	needs input and output types to be declared
*/
func GetStartLine(program string) string {
	return "🚀 " + program + " started"
}
```

### Build and run

```bash
go run gluon.go # builds and runs on memory
go build gluon.go # generates an executable
gluon.exe # you can run it anyware
```

## 0.3 Testing

### Tests have their files *_test.go

On `gluon_test.go` file (view code)[https://github.com/AtomicBuilders/gluon/blob/main/gluon_test.go]

```go
package main // name of this package

import "fmt" // importing external packages

var program string // declare a variable with type string

func main() {
	program = "gluon"             // assign a value to the variable
	line := GetStartLine(program) // initialize and assing a value
	fmt.Println(line)             // print the value of the variable
}

/*
	GetStartLine returns the starting line of the program
	needs input and output types to be declared
*/
func GetStartLine(program string) string {
	return "🚀 " + program + " started"
}
```

The language itself does not help a lot. So you need an utility file. `test-tools.go` [View code](https://github.com/AtomicBuilders/gluon/blob/main/test-tools.go)

```go
package main

import "testing"

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

```

### Run tests

```bash
go run test # builds and runs on memory
```

### [Back to index](https://github.com/AtomicBuilders/gluon/blob/main/docs/0-hello-world.md)