package main

import (
	"fmt"
	"net/http"
	"time"
)

// main function is the entry point of the program.
// It initializes a list of URLs to be checked concurrently using goroutines and channels.
func main() {
	// List of websites to monitor.
	links := []string{
		"https://google.com",
		"https://facebook.com",
		"https://stackoverflow.com",
		"https://golang.org",
		"https://amazon.com",
	}

	// Create a channel to receive status messages from goroutines.
	c := make(chan string)

	// Launch a goroutine for each link to check its status.
	for _, link := range links {
		go checkLink(link, c)
	}

	// Continuously listen for values from the channel.
	// When a message is received, launch a new goroutine that waits 5 seconds before re-checking the link.
	for l := range c {
		// Capturing the value of 'l' to avoid closure pitfalls in goroutines.
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}
}

// checkLink performs a simple HTTP GET request to verify if the link is reachable.
// It sends the link back through the channel regardless of the result,
// allowing the link to be re-checked periodically.
func checkLink(link string, c chan string) {
	// Perform HTTP GET request
	_, err := http.Get(link)

	if err != nil {
		// If there was an error (e.g., the site is down or unreachable), log it.
		fmt.Println(link, "might be down!")
	} else {
		// If the GET request was successful, the link is up.
		fmt.Println(link, "is up!")
	}

	// Send the link back into the channel for the next cycle.
	c <- link
}
