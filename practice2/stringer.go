// interfaces is Stringer defined by the fmt package.
package main

import "fmt"

type Address struct {
	Name    string
	Pincode int
}

func (p Address) String() string {
	return fmt.Sprintf("%v (%v - pincode)", p.Name, p.Pincode)
}

func main() {
	a := Address{"ALD", 2662362}
	z := Address{"KNP", 9001}
	fmt.Println(a, z)
}
