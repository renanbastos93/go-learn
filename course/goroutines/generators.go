package main

import (
	"fmt"
	"net/http"
	"regexp"
	"io/ioutil"
)

func init() {
	// Google I/O 2012 - Go Concurrency Patterns
	// <- chan - channel just read
}

func titles(urls ...string) <-chan string {
	c := make(chan string)
	for _, url := range urls {
		go func(url string) {
			resp, _ := http.Get(url)
			html, _ := ioutil.ReadAll(resp.Body)
			r, _ := regexp.Compile("<title>(.*)?<\\/title>")
			c <- r.FindStringSubmatch(string(html))[1]
		}(url)
	}
	return c
}

func main() {
	t1 := titles("http://renanbastos.com.br", "https://google.com.br")
	t2 := titles("http://renanbastos.com.br", "https://youtube.com.br")
	fmt.Println("First:", <-t1, "|", <-t2)
	fmt.Println("Second:", <-t1, "|", <-t2)
}