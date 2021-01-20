package main

import (
	"fmt"
	"github.com/pratap0007/go-exercise/00/file/filereader"
	"log"
	"sort"
	"strings"
)

func main() {
	filedata, err := filereader.Filereader("/usr/share/doc/coreutils-common/README")
	filedataWithoutspace := strings.Replace(filedata, " ", "\n", -1)
	if err != nil {
		log.Fatal(err)
	}

	filedataword := strings.Fields(filedataWithoutspace)

	m := map[string]int{}
	for _, word := range filedataword {
		m[word]++
	}

		var keys []string

	for mapkey := range m {
		keys = append(keys, mapkey)
	}
	sort.Strings(keys)
	for _,i := range keys{
		fmt.Println(i);
	}

	// type kv struct {
	// 	Key   string
	// 	Value int
	// }
	// var ss []kv

	// for _,k:= range keys{
	// 	ss = append(ss, kv{k, m[k]})
	// }

	// sort.Slice(ss, func(i, j int) bool {
	// 	return ss[i].Value < ss[j].Value
	// })
	// for k,_:= range ss{
	//  	fmt.Println(ss[k].Value, " ",ss[k].Key)
	// }








}
