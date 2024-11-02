package hello

const (
	englishHelloPrefix    = "Hello, "
	spanishLanguagePrefix = "Hola, "
)

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	if language == "Spanish" {
		return spanishLanguagePrefix + name
	}

	return englishHelloPrefix + name
}
