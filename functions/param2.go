package main

import (
	"fmt"
)

func main() {
	greet("Facundo", "Moraleda")
}
func greet(name, lastname string) {
	fmt.Println(name, lastname)
}
