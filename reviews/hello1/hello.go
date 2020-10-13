package hello1

const (
	english = "English"
	french  = "French"
	spanish = "Spanish"
)

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name + "!"
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = "Hola, "
	case french:
		prefix = "Bonjour, "
	default:
		prefix = "Hello, "
	}

	return
}
