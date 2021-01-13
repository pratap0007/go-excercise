// slice does not modify any data only refences to array
//  Changing the elements of a slice modifies the corresponding elements of its underlying array.
package main

import "fmt"

func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	slice := primes[1:4]
	slice[1] = 6326
	fmt.Println(slice)
	fmt.Println(primes[2:])
}
