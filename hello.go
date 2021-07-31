package main

import "fmt"

func main() {
	fmt.Println(hello("World"))
}

const englishHelloPrefix = "Hello "

func hello(name string) string {
	if name == "" {
		return englishHelloPrefix + "world"
	}
	return englishHelloPrefix + name
}
