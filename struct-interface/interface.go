package main

import "fmt"

// Shap Interface
type Shape interface {
	Area() float64
	Perimeter() float64
}


func main(){

var s Shape
fmt.Println(s)
fmt.Printf("type of interface %T",s)

}