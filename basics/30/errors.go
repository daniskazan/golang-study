package main

import (
	"errors"
	"fmt"
)

func f1(arg int) (int, error) {
	if arg == 42 {

		return -1, errors.New("can't work with 42")

	}

	return arg + 3, nil
}

/*
It’s possible
to use custom types as errors by implementing the Error() method on them.
Here’s a variant on the example above that uses a custom type to explicitly represent an argument error.
*/
type argError struct {
	arg  int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
	if arg == 42 {
		// In this case we use &argError syntax to build a new struct, supplying values for the two fields arg and prob.
		return -1, &argError{arg, "can't work with it"}
	}
	return arg + 3, nil
}

func main() {

	for _, i := range []int{7, 42} {
		if res, err := f1(i); err != nil {
			fmt.Println("f1 failed:", err)
		} else {
			fmt.Println("f1 worked:", res)
		}
	}
	for _, i := range []int{7, 42} {
		if res, err := f2(i); err != nil {
			fmt.Println("f2 failed:", err)
		} else {
			fmt.Println("f2 worked:", res)
		}
	}

	_, e := f2(43)
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
}
