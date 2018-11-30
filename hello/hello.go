package hello

import "fmt"

const englishGreetingPrefix = "Hello, "
const spanishGreetingPrefix = "Hola, "
const frenchGreetingPrefix = "Bonjour, "
const italianGreetingPrefix = "Ciao, "
const spanish = "Spanish"
const french = "French"
const italian = "Italian"

func greetingPrefix(lang string) string {
	switch lang {
	case spanish:
		return spanishGreetingPrefix
	case french:
		return frenchGreetingPrefix
	case italian:
		return italianGreetingPrefix
	default:
		return englishGreetingPrefix
	}
}

func Hello(name string, lang string) string {
	if name == "" {
		name = "world"
	}
	return greetingPrefix(lang) + name
}

func main() {
	fmt.Println(Hello("", ""))
}
