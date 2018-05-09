package main

import (
	"fmt"
)

func algo() func() int {
	x := 1
	return func() int {
		x++
		return x
	}
}

func main() {
	y := algo()

	fmt.Println(y())
	fmt.Println(y())

}
