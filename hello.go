package main

import "fmt"

const englishHelloPrefix = "Hello "
const spanishHelloPrefix = "Hola "
const frenchHelloPrefix = "Bonjour "
const spanish = "Spanish"
const french = "French"

func Hello(name, language string) string {
	if name == "" {
		return "Hello world!"
	}
	if language == spanish {
		return spanishHelloPrefix + name
	}
	if language == french {
		return frenchHelloPrefix + name
	}
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("world", ""))
}
