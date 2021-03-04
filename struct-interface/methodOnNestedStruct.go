// method on nested struct

package main

import (
	"fmt"

	"github.com/pratap0007/go-exercise/pkg"
)


type Salary struct {
basic,hra int

}

type Employee struct{
	name string
	dep string
	Salary
}

func ( e *Employee) changeSalary( newBasic int) {
	e.basic = newBasic

}

func ( s *Salary) changeSalary( newBasic int) {
	s.basic = newBasic

}


func main(){

	x := pkg.Employee{
		FirstName : "shiv",
		LastName : "verma",
	}



	e := Employee{ "shiv","tekton",Salary{ 300000, 100000}}

	// s := Salary{50000,20000}

	fmt.Println(e)

	e.changeSalary(50000)
	fmt.Println(e)

	fmt.Println(x)

}