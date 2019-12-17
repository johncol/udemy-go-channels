package main

import (
	"fmt"
	"net/http"
)

var urls = []string{
	"http://google.com",
	"http://facebook.com",
	"http://golang.org",
	"http://faaaake.com",
}

func main() {
	for _, url := range urls {
		reportStatus(url)
	}
}

func reportStatus(url string)  {
	if websiteIsUp(url) {
		fmt.Println("OK", url, "is up and running!")
	} else {
		fmt.Println("WARN", url, "did not respond..")
	}
}

func websiteIsUp(url string) bool {
	_, err := http.Get(url)
	return err == nil
}
