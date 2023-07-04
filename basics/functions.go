package main

// named returned values
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
func add(x, y int) *int {
	sumOfTwo := x + y
	return &sumOfTwo // 0xc000024088
}

// multiple result
func swap(x, y string) (string, string) {
	return y, x
}

func main() {

}
