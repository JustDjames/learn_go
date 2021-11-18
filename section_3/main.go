package main

import "fmt"

func main() {
	cards := newDeck()
	fmt.Println("Before:")
	cards.print()
	fmt.Println("")
	cards.shuffle()
	fmt.Println("After:")
	cards.print()
}
