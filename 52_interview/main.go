package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(2)
	chA := make(chan int, 1)
	chB := make(chan int)

	chA <- 1

	go func() {
		defer wg.Done()
		for i := 0; i < 4; i++ {
			<-chA
			fmt.Printf("%c", 'a'+i)
			chB <- 1
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 4; i++ {
			<-chB
			fmt.Printf("%d", i+1)
			chA <- 1
		}
	}()

	wg.Wait()
}
