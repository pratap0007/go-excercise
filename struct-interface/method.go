package main

import "fmt"

//kekwqkijk
type Empoyee struct{
	firstName,lastName  string
}


// func fullName( firstName string, lastName string) (fullname string){

// 	return firstName + " " + lastName
// }

func (t Empoyee) fullName() string{
	return t.firstName + " " + t.lastName
}

func main(){

	e := Empoyee{ "pratap" ,"0007"}
	fmt.Println("full name of the employe",e.fullName())
// fmt.Println("fullName==", fullName(e.firstName,e.lastName))

}