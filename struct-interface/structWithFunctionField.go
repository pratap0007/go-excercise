package main

import "fmt"


type FullNameType func (string,string) string


// employee struct
type Employee1 struct{
	firstName,lastName string
	fullName FullNameType
}


func main(){

	e := Employee1{
		firstName: "shiv",
		lastName: "verma",
		fullName:  func(f string, l string) string{
			return f + " " + l
		},
	}

fmt.Println("ful name of employee==",e.fullName(e.firstName,e.lastName))

}