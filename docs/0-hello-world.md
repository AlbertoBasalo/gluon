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

On `gluon.go` file

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

