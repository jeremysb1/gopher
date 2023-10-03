package main

import "fmt"

func main() {
	greeting := greet()
	fmt.Println(greeting)
}

// language represents the language's code
type language string

// greet says hello in the specified language

func greet(l language) string {
	switch l {
	case "en":
		return "hello world"
	case "fr":
		return "bonjour le monde"
	default:
		return ""
	}
}
