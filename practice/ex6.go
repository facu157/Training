package main

import (
	"fmt"
)

func main() {
	sexo := map[int]string{
		1: "Hombre",
		2: "Mujer",
	}

	var se int
	fmt.Println("Que sexo sos?")
	fmt.Scan(&se)

	if se == 1 {
		fmt.Println(sexo[1])

	} else {
		fmt.Println(sexo[2])
	}
}
