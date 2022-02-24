package main

import (
	"fmt"
	"net/http"
)

func main() {

	type Result struct {
		Error    error
		Response *http.Response
	}

	checkStatus := func(done <-chan struct{}, urls ...string) <-chan Result {
		responses := make(chan Result)
		go func() {
			defer close(responses)
			for _, url := range urls {
				var result Result
				result.Response, result.Error = http.Get(url)
				select {
				case <-done:
					return
				case responses <- result:
				}
			}
		}()
		return responses
	}

	done := make(chan struct{})
	defer close(done)

	urls := []string{"www.google.com", "https://www.yandex.ru"}
	for result := range checkStatus(done, urls...) {
		if result.Error != nil {
			fmt.Println("Error:", result.Error)
			continue
		}
		fmt.Println("Response:", result.Response.Status)
		//_ = response.Body.Close()
	}
}
