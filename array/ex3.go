package main

import (
	"fmt"
)

func main() {
	var x [6]int
	fmt.Println(x)

	for i := 0; i <= 5; i++ {
		x[i] = i
	}
	fmt.Println(x)

	var nombre string
	fmt.Println("Como es tu nombre?")
	fmt.Scan(&nombre)
	fmt.Println("Es un gusto conocerte ", nombre)
}
