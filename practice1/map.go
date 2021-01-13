package main

import "fmt"

type Data struct {
	name, address string
}

var m map[int]Data

func main() {
	m = make(map[int]Data)
	m[0] = Data{
		"hello", "map",
	}
	fmt.Println(m[0])
}
