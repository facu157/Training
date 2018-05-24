package main

import (
	"fmt"
)

func main() {
	greet("hola")
	greet("chau")
	greet("cualquier cosa")
}
func greet(name string) {
	fmt.Println(name)
}
