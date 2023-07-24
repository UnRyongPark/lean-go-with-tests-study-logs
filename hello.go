package main

import "fmt"

const english = "English"
const spanish = "Spanish"

const englishPrefix = "Hello"
const spanishPrefix = "Hola"

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	if language == "" {
		language = "english"
	}

	switch language {
	case "Spanish":
		return fmt.Sprintf("%s, %s", spanishPrefix, name)
	default:
		return fmt.Sprintf("%s, %s", englishPrefix, name)
	}
}

func main() {
	fmt.Println(Hello("World", ""))
}
