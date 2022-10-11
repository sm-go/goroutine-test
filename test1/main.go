package main

import (
	"fmt"
	"time"
)

func main() {
	go sayHello()
	fmt.Println("good bye!")
	time.Sleep(2 * time.Second)
}

func sayHello() {
	fmt.Println("hello function...")
}
