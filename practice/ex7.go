package main

import (
	"fmt"
)

func main() {
	colores := map[int]string{

		0: "dorado",
		1: "azul",
		2: "rojo",
		3: "amarillo",
		4: "verde",
	}
	for key, value := range colores {

		fmt.Println(key, " - ", value)
	}
	fmt.Println(colores)
}
