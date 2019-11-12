package main

import (
	"fmt"
	"net/http"
	"regexp"
	"io/ioutil"
	"time"
)

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

func moreFast(url1, url2, url3 string) string {
	c1 := titles(url1)
	c2 := titles(url2)
	c3 := titles(url3)

	select {
	case t1 := <-c1:
		return t1
	case t2 := <-c2:
		return t2
	case t3 := <-c3:
		return t3
	case <-time.After(1000 * time.Millisecond):
		return "All lost"
	}
}

func main() {
	champion := moreFast("http://renanbastos.com.br", "https://google.com.br", "https://youtube.com.br")
	fmt.Println(champion)
}