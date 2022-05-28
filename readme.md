# ptr.Of anything!
> Check out [my value package](https://github.com/mrkagelui/value) to do value.Of()!

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/mrkagelui/ptr)
![GitHub Workflow Status](https://img.shields.io/github/workflow/status/mrkagelui/ptr/testing)
[![codecov](https://codecov.io/gh/mrkagelui/ptr/branch/master/graph/badge.svg?token=UK7JNV1P7S)](https://codecov.io/gh/mrkagelui/ptr)
[![Go Report Card](https://goreportcard.com/badge/github.com/mrkagelui/ptr)](https://goreportcard.com/report/github.com/mrkagelui/ptr)
[![GitHub license](https://badgen.net/github/license/mrkagelui/ptr)](https://github.com/mrkagelui/ptr/blob/master/LICENSE)
[![made-with-Go](https://img.shields.io/badge/Made%20with-Go-1f425f.svg)](https://go.dev/)
![goptr](goptr.png "goptr")

## Why this utility? 

In many occasions, we need to take the address of a constant, e.g., an interface requires a nil-able integer, 
without this, we need to do
```go
func thatFunc(i *int) {
	// does something useful
}

func main() {
	myImportantNum := 5
	thatFunc(&myImportantNum)
}
```
This utility helps you do
```go
func thatFunc(i *int) {
	// does something useful
}

func main() {
	thatFunc(ptr.Of(5))
}
```

## Why doesn't Go allow me to do `&5`?

Because `5` is a constant. As [this answer](https://stackoverflow.com/a/35146856) points out, 
if you could take the address of a constant, you could change the value of that constant and 
jeopardize the entire world!

However, often we just want a pointer whose pointed value is given. This library provides a 
syntactic sugar by assigning that constant to a variable and returning the address of that 
variable. Thanks to generics introduced in Go 1.18, this is greatly simplified!
