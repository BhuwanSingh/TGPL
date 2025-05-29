package main

import "fmt"

func appendInt(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1
	if zlen <= cap(x) {
		z = x[:zlen]
	} else {
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
	}
	z[len(x)] = y
	return z
}

func main() {
	a := []int{1, 34, 35, 34, 5346, 4, 57, 565, 8754, 63, 4, 534, 53, 46, 56}
	b := appendInt(a, 8)
	fmt.Println(b)
}
