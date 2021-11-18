package main

import "fmt"

func main() {
	cards := newDeck()
	cards.shuffle()
	myhand, remainingCards := deal(cards, 5)

	fmt.Println("my hand:")
	myhand.print()

	myhand.saveToFile("my_hand")
	remainingCards.saveToFile("remaing_cards")
}
