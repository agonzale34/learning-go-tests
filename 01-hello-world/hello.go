package main

import "fmt"

const frenchLng = "French"
const spanishLng = "Spanish"
const portugueseLng = "Portuguese"

const frenchHelloPrefix = "Bonjour, "
const spanishHelloPrefix = "Hola, "
const englishHelloPrefix = "Hello, "
const portugueseHelloPrefix = "Olá, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case frenchLng:
		prefix = frenchHelloPrefix
	case spanishLng:
		prefix = spanishHelloPrefix
	case portugueseLng:
		prefix = portugueseHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("world", ""))
}
