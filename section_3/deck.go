package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// Create a new type of 'deck'
// which is a slice of strings

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	// underscore means ignore the iterator variable, otherwise you would have to use it

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

// the (d deck) part is a receiver that ensures that this function can only be called by variables of the deck type
// the variable d is called so, as it's convention to use a abbrivation of the type used in the function

func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}

// first return gets the first values in the slice up tp the handsize var. second gets the values starting from the handsize var to the end
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// turns the deck type back into a slice of strings and then turns the slice into a single string joined by using a ','
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

// saves the string to a file. to do this it first needs to convert the string to a slice of bytes
func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bysli, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
