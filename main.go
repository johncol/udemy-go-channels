package main

import (
	"fmt"
	"net/http"
)

// Url x
type Url = string

var urls = []Url{
	"http://google.com",
	"http://facebook.com",
	"http://golang.org",
	"http://faaaake.com",
}

func main() {
	channel := make(chan Url)

	for _, url := range urls {
		go reportStatus(url, channel)
	}

	
	for url := range channel {
		go reportStatus(url, channel)
	}
}

func reportStatus(url Url, channel chan Url)  {
	if websiteIsUp(url) {
		fmt.Println(" --- OK", url, "is up and running!")
	} else {
		fmt.Println(" --- WARN", url, "did not respond..")
	}
	channel <- url
}

func websiteIsUp(url Url) bool {
	_, err := http.Get(url)
	return err == nil
}
