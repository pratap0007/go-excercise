package main

import (
	"fmt"
)

func main() {

	m := make([]bool, 7)
	a := make([]int, 5)
	fmt.Println("a", a)

	b := make([]int, 0, 5)
	fmt.Println("b", b)
	fmt.Println(m)

}
