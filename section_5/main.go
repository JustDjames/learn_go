package main

import (
	"fmt"
)

func main() {
	//  can define a map in a number of ways

	// var colours map[string]string

	// colours := make(map[string]string)
	// colours["white"] = "#ffffff"

	colours := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}

	// used to delete keys in a map
	// delete(colours, "white")
	printMap(colours)
}

func printMap(c map[string]string) {
	for colour, hex := range c {
		fmt.Println("hex code for", colour, "is", hex)
	}
}
