package main

import (
	"fmt"
)

func main() {

	test := func(x int) (float64, bool) {
		return float64(x) / 2, x%2 == 0
	}
	res, tf := test(5)
	fmt.Println(res, tf)
}
