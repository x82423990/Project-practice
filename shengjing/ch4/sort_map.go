package main

import (
	"sort"
	"fmt"
)

// var args = make(map[string]int)
func main() {
	args := map[string]int{
		"xie":   30,
		"liu":   29,
		"zhang": 29,
		"ahi":   81,
	}
	named := make([]string, 0, len(args))
	for names := range args {
		named = append(named, names)
	}
	sort.Strings(named)
	for _,j := range named{
		fmt.Println(j,args[j])
	}
	// nil map
	//var test map[string]int
	//fmt.Println(test == nil)
	//test["bob"] = 1
	//fmt.Println(test)
}
