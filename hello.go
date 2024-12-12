// hello
package main

import (
	"fmt"
)

const prefix = "Hello "

func Hello(s string) string {
	if s == "" {
		s = "world"
	}
	return prefix + s
}

func main() {
	fmt.Println(Hello(""))
}
