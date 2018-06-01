package main

import (
	"fmt"
)

func main() {
	age := 20
	fmt.Println(age)
	fmt.Println(&age)

	//cambio directorio

	change(&age)
}

func change(z *int) {
	*z = 30
	fmt.Println(*z)
	fmt.Println(z)
}
