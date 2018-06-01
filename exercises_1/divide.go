package main

import (
	"fmt"
)

func main() {
	var smallNumber int
	fmt.Println("Enter a small number")
	fmt.Scan(&smallNumber)
	var largeNumber int
	fmt.Println("Enter a large number")
	fmt.Scan(&largeNumber)
	fmt.Println(largeNumber, "Divide", smallNumber, "Remains :", largeNumber%smallNumber)
}
