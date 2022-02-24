package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	checkStatus := func(done <-chan struct{}, urls ...string) <-chan *http.Response {
		responses := make(chan *http.Response)
		go func() {
			defer close(responses)
			for _, url := range urls {
				resp, err := http.Get(url)
				if err != nil {
					log.Println("Error received: ", err)
					continue
				}
				select {
				case <-done:
					return
				case responses <- resp:
				}
			}
		}()
		return responses
	}

	done := make(chan struct{})
	defer close(done)

	urls := []string{"www.google.com", "https://www.yandex.ru"}
	for response := range checkStatus(done, urls...) {
		fmt.Println("Response:", response.Status)
		_ = response.Body.Close()
	}
}
