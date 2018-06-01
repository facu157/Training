package main

import (
	"fmt"
)

func main() {
	slice := make([]string, 4, 5)

	slice[0] = "0"
	fmt.Println(slice[0])
	slice[1] = "1"
	fmt.Println(slice[1])
	slice[2] = "2"
	fmt.Println(slice[2])
	slice[3] = "3"
	fmt.Println(slice[3])

	slice = append(slice, "4")
	slice = append(slice, "5")
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))
}
