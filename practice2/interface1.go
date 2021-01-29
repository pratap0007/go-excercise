// a type implemens an interface by implicity it's method

package main

import (
	"fmt"
)

type X interface {
	dummy()
}

func (v m) dummy() {
	fmt.Println(v.s)

}

type m struct {
	s string
}

type F float64

func (f F) dummy() {
	fmt.Println(f)
}

func main() {

	var t X
	t = m{"hello !"}
	fmt.Printf("%V %T", t, t)

	t = F(55.55)
	t.dummy()

}
