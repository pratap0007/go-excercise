package main

import (
	"fmt"
)

func numbers(done chan bool) {
    for i := 1; i <= 5; i++ {
        //  time.Sleep(250 * time.Millisecond)
        fmt.Printf("%d ", i)


    }
		done<-true
}
func alphabets(done chan bool) {
    for i := 'a'; i <= 'e'; i++ {
        //  time.Sleep(400 * time.Millisecond)
        fmt.Printf("%c ", i)

    }
		done<-true
}
func main() {
	done := make(chan bool)
    go numbers(done)
		<-done
    go alphabets(done)
		<-done
    // time.Sleep(3000 * time.Millisecond)
    fmt.Println("main terminated")
}