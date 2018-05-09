package main

import (
	"fmt"
)

const metersToYards float64 = 1.09361

func main() {
	var meters float64
	fmt.Println("Cuantos metros nadas?")
	fmt.Scan(&meters)
	yards := meters * metersToYards
	fmt.Println(meters, "meters is", yards, "yards.")
	var respuesta float64
	fmt.Println("Queres saber cuantos metros son tus yardas?")
	fmt.Scan(&respuesta)
	resp := respuesta
	fmt.Println(resp, "no te calienta)
}
