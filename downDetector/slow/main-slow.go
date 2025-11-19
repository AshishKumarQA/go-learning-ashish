package main

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"
)

func main() {
	urls := []string{
		"https://gdg.community.dev/gdg-cochin/",
		"https://golang.org",
		"https://httpstat.us/500",
		"https://www.google.com/",
		"https://www.facebook.com/",
		"https://www.twitter.com/",
		"https://www.instagram.com/",
		"https://site-not-present.io",
		"https://www.youtube.com/",
		"https://www.linkedin.com/",
		"https://www.github.com/",
		"https://www.stackoverflow.com/",
		"https://www.reddit.com/",
	}
	var wg sync.WaitGroup
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	ch := make(chan result)
	//defer close(ch)
	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			status, err := checkSite(url)
			if err != nil {
				fmt.Printf("Something wrong in %s, error: %v\n", url, err)
			}
			ch <- result{url, status}

		}(url)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("context timeout")
			return
		case result, ok := <-ch:
			if !ok {
				return
			}
			printer(result.url, result.status)
		}
	}

}

func checkSite(url string) (string, error) {

	res, err := http.Get(url)
	if err != nil {
		return "something went wrong", err
	}
	defer res.Body.Close()
	if res.StatusCode == 200 {
		return "UP", nil
	} else {
		return "DOWN", nil
	}
}

func printer(url string, status string) {
	fmt.Printf("URL: %s is %s\n", url, status)
}

type result struct {
	url    string
	status string
}
