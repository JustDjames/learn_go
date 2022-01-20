package main

import "fmt"

type englishBot struct{}
type spanishBot struct{}

// this interface type "bot" will get any other types with a function called "getGreeting" that returns a string and includes them in the bot type. those other types can then use the function "printGreeting" as it requires a input of type bot
type bot interface {
	getGreeting() string
}

func (englishBot) getGreeting() string {
	return "Hey David"
}

func (spanishBot) getGreeting() string {
	return "Hola David"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func main() {
	span := spanishBot{}
	eng := englishBot{}

	printGreeting(span)
	printGreeting(eng)
}
