package main

import (
	"fmt"
)

func main() {
	// || funciona como "or", si uno de los dos es verdadero, funciona.
	if true || false {
		fmt.Println("This run")
	}
	// ! es el contrario del bool que pongo
	if !false {
		fmt.Println("This doesnt run")
	}
	// && significa Y, ambos tienen que ser true para andar
	if true && false {
		fmt.Println("This doesnt run")
	}

	if true && true && true {
		fmt.Println("This run")
	}

}
