package main

import "fmt"

func main() {
	fmt.Println(Hello("World", ""))
}

const englishPrefix = "Hello "
const spanishPrefix = "Hola "
const spanishLanguage = "Spanish"
const defaultName = "world"
const frenchLanguage = "French"
const frenchPrefix = "Salut "

func Hello(name string, lang string) string {
	if name == "" {
		return englishPrefix + defaultName
	}

	return greetingPrefix(lang) + name
}

func greetingPrefix(lang string) (prefix string) {
	switch lang {
	case spanishLanguage:
		prefix = spanishPrefix
	case frenchLanguage:
		prefix = frenchPrefix
	default:
		prefix = englishPrefix
	}
	return
}
