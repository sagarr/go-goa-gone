package main

import "fmt"

func Hello(name, langauge string) string {
	if name == "" {
		name = "World"
	}
	prefix := greetingPrefix(langauge)
	return prefix + ", " + name
}

func greetingPrefix(langauge string) (prefix string) {
	switch langauge {
		case "french":
			prefix = "Bojour"
		case "spanish":
			prefix = "Hola"
		default:
			prefix = "Hello"
	}
	return
}

func main() {           
	fmt.Println(Hello("World", "English"))
}
