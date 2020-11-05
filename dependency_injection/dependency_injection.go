package main

import (
	"fmt"
	"io"
	"net/http"
)

func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

func MyGreetHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "Paul!")
}

func main() {
	// Greet(os.Stdout, "Hello!")
	http.ListenAndServe(":5000", http.HandlerFunc(MyGreetHandler))
}
