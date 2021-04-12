// Go program to illustrate how
// to close a channel using for
// range loop and close function
package main

import "fmt"

// Function
func myfun(mychnl chan string) {

	for v := 0; v < 10; v++ {
		mychnl <- "GeeksforGeeks"
	}
	// close(mychnl)
}

func dff(x chan string){
	fmt.Print("3266326776")
	x <- "xxx"
}

// Main function
func main() {

	// Creating a channel
	c := make(chan string)

	// calling Goroutine
	go myfun(c)

	go dff(c)

	// When the value of ok is
	// set to true means the
	// channel is open and it
	// can send or receive data
	// When the value of ok is set to
	// false means the channel is closed
	for {
		res, ok := <-c
		if ok == false {
			fmt.Println("Channel Close ", ok)
			break
		}
		fmt.Println("Channel Open ", res, ok)
	}
}
