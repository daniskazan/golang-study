package main

import "fmt"

func main() {

	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	// While True:
	for {
		fmt.Println("loop")
		break
	}

	for n := 0; n <= 6; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
