package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	fmt.Println("hey shiva move fast othervise u can't completer your dream")
	filecontent, err := ioutil.ReadFile("/usr/share/doc/coreutils-common/README")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(filecontent))

}
