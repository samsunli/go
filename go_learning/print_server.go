package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string)
	go print(c)
	time.Sleep(1 * time.Second)
	fmt.Println("main function: start writing msg")
	c <- "hello"

	var input string
	fmt.Scanln(&input)
}

func print(c <-chan string) {
	fmt.Println("print function: start reading")
	fmt.Println("print function: reading: " + <-c)
	time.Sleep(1 * time.Second)
}