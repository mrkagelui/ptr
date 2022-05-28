# ptr.Of anything!
[![GitHub go.mod Go version of a Go module](https://img.shields.io/github/go-mod/go-version/gomods/athens.svg)](https://github.com/gomods/athens)
[![GoReportCard](https://goreportcard.com/badge/github.com/mrkagelui/ptr)](https://goreportcard.com/report/github.com/mrkagelui/ptr)
[![GitHub license](https://badgen.net/github/license/mrkagelui/ptr)](https://github.com/mrkagelui/ptr/blob/master/LICENSE)
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

## Why doesn't the standard library allow me to do `&5`?

Because `5` is a constant. As [this answer](https://stackoverflow.com/a/35146856) points out, 
if you could take the address of a constant, you could change the value of that constant and 
jeopardize the entire world!

However, often we just want a pointer whose pointed value is given. This library provides a 
syntax sugar by assigning that constant to a variable and returning the address of that 
variable. Thanks to generics introduced in Go 1.18, this is greatly simplified!
