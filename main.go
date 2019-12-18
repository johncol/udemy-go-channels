package main

import (
	"fmt"
	"net/http"
	"time"
)

// URL x
type URL = string

var urls = []URL{
	"http://google.com",
	"http://facebook.com",
	"http://golang.org",
	"http://faaaake.com",
}

func main() {
	channel := make(chan URL)

	for _, url := range urls {
		go reportStatus(url, channel)
	}

	for url := range channel {
		go func(url URL) {
			time.Sleep(5*time.Second)
			reportStatus(url, channel)
		}(url)
	}
}

func reportStatus(url URL, channel chan URL)  {
	if websiteIsUp(url) {
		fmt.Println(" --- OK", url, "is up and running!")
	} else {
		fmt.Println(" --- WARN", url, "did not respond..")
	}
	channel <- url
}

func websiteIsUp(url URL) bool {
	_, err := http.Get(url)
	return err == nil
}
