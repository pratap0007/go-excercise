package main

import "fmt"



func dummy(c chan string){
	fmt.Println(" hello " + <-c + "!!!")
	close(c)
}

func main(){
	fmt.Println("main started")
	c:= make(chan string)
	// go dummy(c)
	c<- "pratap007"

	fmt.Println("main stpped")
}