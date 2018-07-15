package main

import (
	"os"
	"fmt"
	"net/http"
	"io/ioutil"
)

func main() {
	//params := os.Args[1:]
	println(os.Args[1:])
	println(len(os.Args[1:]))
	//
	//if len(os.Args[1:]) != 1 {
	//	fmt.Println("url is illegal")
	//	os.Exit(1)
	//}
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s", b)
	}
}
