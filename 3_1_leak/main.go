package main

import (
	"fmt"
	"time"
)

func main() {
	doWork := func(strings <-chan string) <-chan struct{} {
		completed := make(chan struct{})
		go func() {
			defer func() {
				fmt.Println("DoWork exited")
				close(completed)
			}()
			for s := range strings {
				fmt.Println(s)
			}
		}()
		return completed
	}

	ch := make(chan string)
	doWork(ch)

	time.Sleep(time.Second * 5)
	fmt.Println("Done")
}
