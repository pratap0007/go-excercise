package main

import "fmt"

// salary
type Salary struct{
	basic int
	allowance int
}

// Employee details
	//  create a nested struct
// type Employee struct {
// 	firstName string;
// 	lastName string;
// 	salary salary;
// 	fullTime bool

// }

// annonymous nested strcture

type Employee struct {
	firstName string;
	lastName string;
	fullTime bool;
	Salary // annonymous struct
}



	/// struct with annonymous fields
	// decalre struct with annonymous fields

	type Annonymous struct{
		string
		int
		bool
	}






func main() {

	y := Annonymous{ "ooo",700,false}

	// var x Employee
	// x.firstName ="abc"
	// x.lastName ="cde"
	// x.fullTime = true

	// initializing a struct

	// x := Employee{
	// 	firstName: "abc",
	// 	lastName: "cde",
	// 	salary: 30000,
	// 	fullTime: false,
	// }


// needd to provide all the fields in their orders
	// x := Employee {"abc","cde",3000,true}


	// pointer to struct
	x := &Employee{
		firstName: "abc",
		lastName: "cde",
		Salary: Salary{50000,5000},
		fullTime: false,

	}


	fmt.Println(x.Salary)
	fmt.Println(x.Salary.basic)
	fmt.Println(y)
	// fmt.Println(y.string)
	// fmt.Println(y.int)


}