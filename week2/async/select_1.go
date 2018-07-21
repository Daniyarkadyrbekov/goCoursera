package main

import (
	"fmt"
)

func main() {
	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)
	ch1 <- 1

	go func(ch2 chan int) {
		newch := <-ch2
		fmt.Printf("read newch %v\n", newch)
	}(ch2)

	select {
	//case val := <-ch1:
	//	fmt.Println("ch1 val", val)
	case ch2 <- 1:
		fmt.Println("put val to ch2")
	default:
		fmt.Println("default case")
	}

	fmt.Scanln()

	//	go func(newch chan int) {
	//		newch := <- ch2
	//		fmt.Printf("read newch %v\n", newch)
	//	}(ch2)
}
