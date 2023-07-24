package main

import "fmt"

const english = "English"
const spanish = "Spanish"
const franch = "Franch"

const englishPrefix = "Hello"
const spanishPrefix = "Hola"
const franchPrefix = "Bonjour"

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	if language == "" {
		language = "english"
	}

	switch language {
	case spanish:
		return fmt.Sprintf("%s, %s", spanishPrefix, name)
	case franch:
		return fmt.Sprintf("%s, %s", franchPrefix, name)
	default:
		return fmt.Sprintf("%s, %s", englishPrefix, name)
	}
}

func main() {
	fmt.Println(Hello("World", ""))
}
