package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	ch := make(chan int)
	wg.Add(2)
	go func() {
		i := <-ch
		fmt.Println(i)
		ch <- 25
		wg.Done()
	}()
	go func() {
		ch <- 43
		fmt.Println(<-ch)
		wg.Done()
	}()
	wg.Wait()
}
