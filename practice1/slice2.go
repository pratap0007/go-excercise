package main

import "fmt"

func main() {
	a := []int{3, 4, 5, 6}

	b := []string{"aa", "bb", "cc"}
	list := []struct {
		i int
		b bool
	}{
		{3, true},
		{7, false},
	}
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(list)
	fmt.Printf("len=%d cap=%d \n", len(a), cap(a[1:]))

}
