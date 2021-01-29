// interface is defined as set of method signature
package main

import "fmt"

//List ...
type List interface {
	Abs() float64
}

type Vertex struct {
	x, y float64
}

func main() {

	c := Vertex{3.5, 6.5}

	fmt.Println(c.Abs())

}

func (i Vertex) Abs() float64 {
	return i.x + i.y

}
