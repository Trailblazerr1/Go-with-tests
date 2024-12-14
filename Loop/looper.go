package main

import "fmt"

func Looper(s string, no int) (repeatedStr string) {
	for i := 0; i < no; i++ {
		repeatedStr = repeatedStr + s
	}
	return
}

func main() {
	result := Looper("Hello ", 10)
	fmt.Print(result)
}
