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

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

// the same logic
// func printGreeting(eb englishBot) {
// 	fmt.Println(eb.getGreeting())
// }

// the same logic so we use interface to solve this problem
// func printGreeting(sb spanishBot) {
// 	fmt.Println(sb.spanishBot())
// }

func (eb englishBot) getGreeting() string {
	// Very custom logic for generation an english greeting
	return "Hi There!"
}
func (sb spanishBot) getGreeting() string {
	return "Hola!"
}
