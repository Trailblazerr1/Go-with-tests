package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

// by default fmt.print writes at stdout, but if we check go code, printf calls a emthod fprintf which takes
// the medium where to output as a parameter, that medium must be of type io.Writer, so it could be buffer, file
// system or stdout
// var nameConst = "Hello"

func Greet(name string, writer io.Writer) {
	fmt.Fprintf(writer, "Hello %s", name)
}

// io.Writer also takes writer on a server via http , http.ResponseWriter is used for this
func GreetOnServer(writer http.ResponseWriter, request *http.Request) {
	Greet("Server", writer)
}

// Here output is sent to stdout
func main() {
	//Greet("duniya", os.Stdout)
	//This method calls GreetOnServer function and returns two params a writer and the request which is used
	//as params in GreeOnServer
	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(GreetOnServer)))
}
