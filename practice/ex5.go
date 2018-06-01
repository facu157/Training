package main

import (
	"fmt"
	"time"
)

func main() {

	currentDate := time.Now()
	fmt.Println(currentDate)

	currentYear := currentDate.Year()

	fmt.Println(currentYear)

	var ano int
	fmt.Println("En que año naciste?")
	fmt.Scan(&ano)

	edad := currentYear - ano

	fmt.Println("Tenes", edad, "años")

}
