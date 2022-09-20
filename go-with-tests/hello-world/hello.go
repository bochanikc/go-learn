package helloworld

import "fmt"

const (
	englishHelloPrefix = "Hello"
	spanishHelloPrefix = "Hola"
	frenchhHelloPrefix = "Bonjour"
)

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}
	return fmt.Sprintf("%s, %s!", getPrefix(language), name)
}

func getPrefix(language string) (prefix string) {
	switch language {
	case "Spanish":
		prefix = spanishHelloPrefix
	case "French":
		prefix = frenchhHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return prefix
}

func main() {
	fmt.Println(Hello("Chris", ""))
}
