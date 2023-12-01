package main

import "fmt"

/*
range iterates over elements in a variety of data structures.
Let’s see how to use range with some of the data structures we’ve already learned.
*/
func main() {

	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	for k := range kvs {
		fmt.Println("key:", k)
	}

	/*
		range on strings iterates over Unicode code points.
		The first value is the starting byte index of the rune and the second the rune itself.
		See Strings and Runes for more details.
	*/
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
