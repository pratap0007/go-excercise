package main

import (
	"fmt"
)
type Vertex struct {
	X int
	Y int
	Z string
}

func  main()  {

	fmt.Println(Vertex{1, 2,"string"})
	c := Vertex{1, 2,"string"}
	fmt.Println(c.X)
	fmt.Println(&c.Z)



}