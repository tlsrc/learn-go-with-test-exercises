package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main2() {
	Greet(os.Stdout, "you!")
}

func main() {
	log.Fatal(http.ListenAndServe(":5000", http.HandlerFunc(GreeterHandler)))
}

func GreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}

func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello %s", name)
}
