package main

import (
	"fmt"
)

func main() {
	m := make(map[string]string)

	m["Nombre"] = "Facundo"

	v1 := m["Nombre"]

	fmt.Println(v1)

}
