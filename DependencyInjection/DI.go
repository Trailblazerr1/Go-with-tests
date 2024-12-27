package main

import (
	"fmt"
	"io"
	"os"
)

// by default fmt.print writes at stdout, but if we check go code, printf calls a emthod fprintf which takes
// the medium where to output as a parameter, that medium must be of type io.Writer, so it could be buffer, file
// system or stdout
// var nameConst = "Hello"

func Greet(name string, writer io.Writer) {
	fmt.Fprintf(writer, "Hello %s", name)
}

// Here output is sent to stdout
func main() {
	Greet("duniya", os.Stdout)
}
