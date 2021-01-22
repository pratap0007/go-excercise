// go  does not have class it has method with type

package main

import (
	"fmt"
)

type Fint int

type Add struct {
	a, b float64
}

type Diff struct {
	a, b int64
}

func (x Add) Addition() float64 {
	return x.a + x.b
}

func (d Diff) Diifrence() int64 {
	return d.a - d.b

}
func (x Fint) Intx() int {
	return int(x)

}

func Next(s string) string {

	return s + "iiiii"

}

func main() {
	sum := Add{4, 5}
	d := Diff{9, 7}
	fmt.Println(sum.Addition())
	fmt.Println(d.Diifrence())
	fmt.Println(d.Diifrence())
	fmt.Println(Next("shhhhh"))

	t := Fint(3000)
	fmt.Println(t.Intx())
}
