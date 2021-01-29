package main

import (
	"fmt"
)

func TypeList(i interface{}) {

	switch i.(type) {
	case int:
		fmt.Println("int")
	case float64:
		fmt.Println("float")

	case string:
		fmt.Println("string")
	default:
		fmt.Println("no type")
	}

}

func main() {

	TypeList("aaa")
	TypeList(32.437)
	TypeList(3234)
}
