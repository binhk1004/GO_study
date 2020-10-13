package main

import (
	"errors"
	"fmt"
	"net/http"
	"time"
)

var errRequestFailed = errors.New("요청이 실패 하였습니다.")

func main()  {
	var results = make(map[string]string)
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
		result := "OK"
		err := hitURL(url)
		if err != nil {
			result = "FAILED"
		}
		results[url] = result
	}
	for url, result := range results{
		fmt.Println(url, result)
	}
	c := make(chan bool)
	people := [2]string{"bin", "jina"}
	for _, person := range people{
		go isSexy(person, c)
	}
	fmt.Println(<- c)
	fmt.Println(<- c)
}

func hitURL(url string) error {
	fmt.Println("확인중: ",url)
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode >= 400 {
		fmt.Println(err, resp.StatusCode)
		return errRequestFailed
	}
	return nil
}

func isSexy(person string, c chan bool) {
	time.Sleep(time.Second * 5)
	fmt.Println(person)
	c <- true
}