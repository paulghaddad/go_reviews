package hello2

const (
	english = "English"
	spanish = "Spanish"
	french  = "French"
)

func Hello(name, language string) string {

	if name == "" {
		name = "World"
	}

	return greeting(language) + name + "!"
}

func greeting(language string) (salutation string) {
	switch language {
	case spanish:
		salutation = "Hola, "
	case french:
		salutation = "Bonjour, "
	default:
		salutation = "Hello, "
	}

	return salutation
}
