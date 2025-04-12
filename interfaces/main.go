package main

type englishBot struct{}
type spanishBot struct{}

type bot interface {
	getGreeting() string
}

func main() {
	eb := englishBot{} // Create an instance of englishBot
	sb := spanishBot{} // Create an instance of spanishBot

	printGreeting(eb) // Call the printGreeting function with eb
	printGreeting(sb) // Call the printGreeting function with sb
}

func printGreeting(b bot) {
	println(b.getGreeting())
}

func (eb englishBot) getGreeting() string {
	return "Hello!"
}

func (sb spanishBot) getGreeting() string {
	return "Hola!"
}