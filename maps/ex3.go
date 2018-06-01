package main

import (
	"fmt"
)

func main() {
	var name string
	fmt.Println("Nombre: ")
	fmt.Scan(&name)

	var lastname string
	fmt.Println("Apellido: ")
	fmt.Scan(&lastname)

	map1 := make(map[string]string)

	map1["nombre"] = name

	map1["apellido"] = lastname

	v1 := map1["nombre"]

	v2 := map1["apellido"]

	fmt.Println(v1, v2)

}
