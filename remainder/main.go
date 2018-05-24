package main

import (
	"fmt"
)

func main() {

	x := 2

	if x == 2 {
		fmt.Println("Par")
	} else {
		fmt.Println("Impar")
	}

	for i := 1; i < 100; i++ {
		if i%2 == 1 {
			fmt.Println("Impar")
		} else {
			fmt.Println("Par")
		}
	}

}
