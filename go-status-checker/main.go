package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {
	sites := []string{
		"https://google.com",
		"https://amazon.com",
		"https://golang.org",
		"https://golang.org/idjhfsdjklfdklfldsfjl",
		"https://stackoverflow.com",
		"https://microsoft.com"}
	c := make(chan string)
	for _, site := range sites {
		go checkSite(site, c)
	}
	for l := range c {
		go func(url string) {
			time.Sleep(5 * time.Second)
			checkSite(url, c)
		}(l)
	}
}

func checkSite(url string, c chan string) {

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	fmt.Println(time.Now(), "-", url, "-", resp.StatusCode, "-", resp.Status)
	c <- url
	//io.Copy(os.Stdout, resp.Body)
}
