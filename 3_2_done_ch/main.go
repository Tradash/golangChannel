package main

import (
	"fmt"
	"time"
)

func main() {
	doWork := func(done <-chan struct{}, strings <-chan string) <-chan struct{} {
		terminated := make(chan struct{})
		go func() {
			defer func() {
				fmt.Println("DoWork exited")
				close(terminated)
			}()
			for {
				select {
				case s := <-strings:
					fmt.Println(s)
				case <-done:
					return
				}
			}
		}()
		return terminated
	}

	done := make(chan struct{})
	ch := make(chan string)
	terminated := doWork(done, ch)
	doWork(done, ch)

	go func() {
		time.Sleep(time.Second)
		fmt.Println("Canceling doWork goroutine...")
		close(done)
	}()
	fmt.Println("Before Done")
	<-terminated
	fmt.Println("Done")
}
