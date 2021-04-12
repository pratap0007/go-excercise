package main

import (
	"fmt"
)

func main() {

    // create a buffered channel
    // with a capacity of 2.
    ch := make(chan string)
    ch <- "abc"
    ch <- "def"
		ch <- "ghi"
    fmt.Println(<-ch)
    fmt.Println(<-ch)
		fmt.Println(<-ch)

		fmt.Println("stopped main")
}