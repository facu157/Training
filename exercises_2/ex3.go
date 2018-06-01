package main

import (
	"fmt"
)

var grande int

func test(numeros ...int) int {

	for _, v := range numeros {
		if v > grande {
			grande = v
		}
	}

	return grande
}

func main() {
	// masgrande (sin importar el nombre que le ponga), es el return de la funcion anterior
	masgrande := test(1, 2, 3, 4, 5)
	fmt.Println(masgrande)
}
