package main

import "fmt"
import "math/rand"

func main() {
	var (
		first  int
		second int
	)
	first = rand.Intn(100)
	second = rand.Intn(100)
	if first%second == 0 {
		fmt.Println("first is divisible by second")
	} else {
		fmt.Println("first is divisible by second")
	}

	if third := rand.Intn(100); third < 0 {
		fmt.Println(first, "is negative")
	} else if first > 100 {
		fmt.Println(first, "greater than 100")
	} else {
		fmt.Println(first, "default")
	}
}
