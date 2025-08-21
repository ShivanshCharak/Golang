package v5

import "fmt"

const Spanish = "Spanish"
const English = "English"
const French = "French"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

func Hello(name string, language string) string {
	if name == "" {
		name = "world"
	}
	prefix := englishHelloPrefix
	switch language {
	case English:
		prefix = englishHelloPrefix
	case Spanish:
		prefix = spanishHelloPrefix
	case French:
		prefix = frenchHelloPrefix
	}
	return prefix + name
}

func main() {
	fmt.Println(Hello("world", ""))
}
