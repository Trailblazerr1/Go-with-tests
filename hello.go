// hello
package main

import (
	"fmt"
)

// group for readablity
const (
	prefixEng   = "Hello "
	prefixHindi = "Namaste"
	prefixUrdu  = "Salaam"

	Hindi = "Hindi"
	Urdu  = "Urdu"
)

func Hello(s, lang string) string {
	if s == "" {
		s = "world"
	}
	return returnGreeting(lang) + s
}

// we can choose not to write whole prefix string here, can write just the type
// but it will help in doc generation, and less return msgs. This is same as declartion
// No need to declare again
func returnGreeting(lang string) (prefix string) {
	switch lang {
	case Hindi:
		prefix = prefixHindi
	case Urdu:
		prefix = prefixUrdu
	default:
		prefix = prefixEng
	}
	return //automatically return prefix
}

func main() {
	fmt.Println(Hello("", "Hindi"))
}
