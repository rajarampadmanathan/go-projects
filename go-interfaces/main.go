package main

import "fmt"

type bot interface {
	getGreeting() string
}
type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}
	// s := "test"
	// printBotGreeting(s)
	printBotGreeting(eb)
	printBotGreeting(sb)
}

func printBotGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func printGreeting(eb englishBot) {
	fmt.Println(eb.getGreeting())
}

func printGreetingsb(sb spanishBot) {
	fmt.Println(sb.getGreeting())
}

func (eb englishBot) getGreeting() string {
	return "Hello !!!"
}

func (sb spanishBot) getGreeting() string {
	return "Hola !!!"
}
