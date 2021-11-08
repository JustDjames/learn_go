package main

import "fmt"

func main() {
	cards := newDeck()
	fmt.Println("Before:\n")
	cards.print()
	fmt.Println("\n")
	cards.shuffle()
	fmt.Println("After:\n")
	cards.print()
}
