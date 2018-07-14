package main

import (
	"io/ioutil"
	"os"
)

func main() {
	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		if err !=nil
	}
}
