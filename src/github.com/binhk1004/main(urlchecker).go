package main

import (
	"errors"
	"fmt"
	"net/http"
)

type requestResult struct {
	url string
	status string
}
var errRequestFailed = errors.New("요청이 실패 하였습니다.")

func main()  {
	results := make(map[string]string)
	c := make(chan requestResult)
	urls := []string{
		"https://www.airbnb.co.kr/",
		"https://www.google.co.kr/",
		"https://www.amazon.com/",
		"https://www.reddit.com/",
		"https://soundcloud.com/",
		"https://www.facebook.com/",
		"https://www.instagram.com/",
		"https://nomadcoders.co/",
	}
	for _, url := range urls {
		go hitURL(url, c)
	}
	for i := 0; i < len(urls); i++{
		result := <-c
		results[result.url] = result.status
	}
	for url, status := range results{
		fmt.Println(url, status)
	}
}

func hitURL(url string, c chan<- requestResult)  {
	resp, err := http.Get(url)
	status := "성공 했습니다."
	if err != nil || resp.StatusCode >= 400 {
		status = "실패 했습니다."
	}
	c <- requestResult{url: url, status: status}
}