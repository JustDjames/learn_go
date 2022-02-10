// This program runs through a list of websites to check if they are up
package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	// dontwork.com isn't a website so it fails
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.co.uk",
		"http://dontwork.com",
	}
	// defining the string channel
	c := make(chan string)

	for _, link := range links {
		//  the `go` keyword runs the function in a seperate go routine. this allows each of the function calls to run concurrently. however if we use just this, the main.go routine will complete without waiting for the responses from the child go routine, so we won't know if the links are up

		//  the channel is used to communicate with the child go routines so that the main.go routine can receive the responses
		go checkLink(link, c)
	}

	// We are waiting for the channel to return a value (in this case a link) from the child go routines, assigning that value to `l` and then sending it into our function literal which sleeps and calls the checkLink function. once the channel receives another message, it will do the same thing again and again, effectively acting as a infinite loop. if it wasn't in a for loop, it would send the first link it receives to the function literal and would finish without waiting for the messages from the other child go routines as there isn't any more code to run

	for l := range c {

		// We pass `l` as a argument so that the value used by the function is stored in a different memory location to the original `l`. otherwise, the function's go routine will be trying to read a value that is constantly being updated by the main go routine. it will fail at this and continually send the same value to the checkLink method
		go func(link string) {
			// we add a sleep so that we don't send too many requests to the links.
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	//  we don't care about the response only the error which is why we use `_`
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, "might be down!")
		// here we are sending a message to the main function through the channel. in this case the message is the link
		c <- link
		fmt.Println("got error:", err)
	} else {
		fmt.Println(link, "is up!")
		c <- link
	}
}
