package main

import (
	"fmt"
	"strings"
)

// derived type
type MyString string

func (s MyString ) toUpperCase() string{
	normalString := string(s)
	return strings.ToUpper(normalString)
}

func main(){

	m := MyString("shiv verma")
	fmt.Println(m.toUpperCase())

}