package main

import "fmt"

func main() {
	greeting := greet("en")
	fmt.Println(greeting)
}

// language represents the language's code
type language string

// phrasebook holds greeting for each supported language
var phrasebook = map[language]string{
	"el": "Χαίρετε Κόσμε",     // Greek
    "en": "Hello world",       // English
    "fr": "Bonjour le monde",  // French
    "he": "שלום עולם",         // Hebrew
    "ur": "ہیلو",         // Urdu
    "vi": "Xin chào Thế Giới", // Vietnamese
}

// greet says hello in the specified language

func greet(l language) string {
	greeting, ok := phrasebook[l]
	if !ok {
		return fmt.Sprintf("unsupported language: %q", l)
	}
	return greeting
}
