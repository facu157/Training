package main

import (
	"fmt"
)

func main() {
	fmt.Println(greet("Facundo ", "Moraleda"))
}

func greet(name, lastname string) string {
	return fmt.Sprint(name, lastname)
}
