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
	channel := make(chan bool)

	for _, url := range urls {
		go reportStatus(url, channel)
	}

	
	for i := 1; i <= len(urls); i++ {
		fmt.Println(<- channel)
	}
}

func reportStatus(url string, channel chan bool)  {
	isUp := websiteIsUp(url)
	if isUp {
		fmt.Println(" --- OK", url, "is up and running!")
	} else {
		fmt.Println(" --- WARN", url, "did not respond..")
	}
	channel <- isUp
}

func websiteIsUp(url string) bool {
	_, err := http.Get(url)
	return err == nil
}
