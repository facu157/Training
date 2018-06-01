package main

import (
	"fmt"
)

func main() {
	var name string
	fmt.Println("Hola como te llamas?")
	fmt.Scan(&name)

	nomb := name

	m := make(map[string]string)

	m["Nombre"] = nomb

	v1 := m["Nombre"]

	fmt.Println("Hola ", v1)

}
