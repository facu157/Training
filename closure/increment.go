package main

import (
	"fmt"
)

var x = 0

func increment() int {
	x++
	return x

}
func main() {
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())

	var y = "Facu el mejor"
	fmt.Println(y)

}
