package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for _, i := range []int{1, 2, 3} {
		fmt.Println(from, ": ", i)
	}
}

func main() {

	f("direct")

	go f("goroutine")

	go func(msg string) {
		nums := []int{1, 2, 3}
		for _, i := range nums {
			fmt.Println(msg, ": ", i)
		}
	}("going")

	time.Sleep(time.Second)
	fmt.Println("done")
}
