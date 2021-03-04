package main

import "fmt"


type Rectangle struct{
	height , width int
}

type Circle struct{
	radius int
}


func (t Circle) Area() int {
	return t.radius*t.radius*3
}

func(t Rectangle) Area() int {
	return t.height*t.width
}

func main(){

	c := Circle{5}
	r := Rectangle{4,5}

	fmt.Println("area of circle", c.Area())
	fmt.Println("area of rectangle", r.Area())


}