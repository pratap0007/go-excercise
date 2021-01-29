// pointer reciever

package main

import (

	"fmt"
)

//Index ...
type Index struct{
	x string
}

//List ...
func (t *Index)  List(m string)  string{
return t.x + " " + m
}


func main() {

	 g := Index{"hello"}
	 fmt.Println(g.List("world"))
}