package main

import (
	"fmt"
)

func main() {
	var x [123]string
	for i := 65; i <= 122; i++ {
		x[i] = string(i)
	}
	fmt.Println(x)
	fmt.Println(x[70])
	fmt.Println(x[65])
	fmt.Println(x[67])
	fmt.Println(x[85])

}
