package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://amazon.com",
		"http://golang.org",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://kcbgroup.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}
	// for i := 0; i < len(links); i++ {
	// 	fmt.Println(<-c)
	// }
	// for {
	// 	go checkLink(<-c, c)
	// }
	// for l := range c { // loop thru items in the channel
	// 	go checkLink(l, c)
	// }
	// Sleep
	for l := range c {
		go func(link string) { //function literal
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		// c <- "Might be down for real!"
		c <- link
		return
	}
	fmt.Println(link, "is up!")
	// c <- "The link is up"
	c <- link
}
