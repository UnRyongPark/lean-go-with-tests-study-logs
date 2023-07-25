package hello

import "fmt"

const (
	english = "English"
	spanish = "Spanish"
	franch  = "Franch"

	englishPrefix = "Hello"
	spanishPrefix = "Hola"
	franchPrefix  = "Bonjour"
)

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	if language == "" {
		language = "english"
	}

	return fmt.Sprintf("%s, %s", greetingPefix(language), name)
}

func greetingPefix(language string) (prefix string) {
	switch language {
	case spanish:
		return spanishPrefix
	case franch:
		return franchPrefix
	default:
		return englishPrefix
	}
}

func main() {
	fmt.Println(Hello("World", ""))
}
