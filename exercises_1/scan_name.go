package main

import (
	"fmt"
)

func main() {
	var myName string
	fmt.Println("Your name is?")
	fmt.Scan(&myName)
	fmt.Println("Hello,", myName)
}
