package main

import "fmt"

const (
	englishHelloPrefix = "Hello "
	spanishHelloPrefix = "Hola "
	frenchHelloPrefix  = "Bonjour "
	spanish            = "Spanish"
	french             = "French"
)

func getLanguagePrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func Hello(name, language string) string {
	if name == "" {
		return "Hello world!"
	}
	prefix := getLanguagePrefix(language)
	return prefix + name
}

func main() {
	fmt.Println(Hello("world", ""))
}
