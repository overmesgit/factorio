package main

import (
	"fmt"
)

func main() {

	c := make(chan int, 2)
	c <- 1
	c <- 1
	select {
	case c <- 1:
		fmt.Println("a")
	default:
		fmt.Println("b")
	}
}
