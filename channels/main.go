package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	sites := []string{
		"http://facebook.com",
		"http://golang.org",
		"http://amazon.com",
		"http://stackoverflow.com",
		"http://google.com",
	}

	c := make(chan string)

	for _, site := range sites {
		go checkSite(site, c)
	}

	for s := range c {
		go func() {
			time.Sleep(5 * time.Second)
			checkSite(s, c)
		}()
	}

}

func checkSite(site string, c chan string) {
	_, err := http.Get(site)
	if err != nil {
		fmt.Println(site, "might be down")
		c <- site
		return
	}

	fmt.Println(site, "is up")
	c <- site
}
