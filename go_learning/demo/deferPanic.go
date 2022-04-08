package main

import "fmt"

func main() {
    defer fmt.Println("Clear resources1")
	defer func() {
		fmt.Println("Clear resources2")
	}()

    panic("Fatal error")

}