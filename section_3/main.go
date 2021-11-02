package main

import "fmt"

func main() {
	cards := newDeck()

	hand, remainingCards := deal(cards, 5)

	fmt.Println("my hand:\n")
	hand.print()
	fmt.Println("\nremaining cards:\n")
	remainingCards.print()
}
