package main

import (
	"fmt"
)

func test(x int) (float64, bool) {
	return float64(x) / 2, x%2 == 0
}

func main() {
	res, tf := test(5)
	fmt.Println(res, tf)
}
