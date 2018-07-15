package main

import (
	"os"
	"bufio"
	"fmt"
)

/*map 是一个由 make 函数创建的数据结构的引用。 map 作为为参数传递给某函数时，该函数
接收这个引用的一份拷贝（copy，或译为副本），被调用函数对 map 底层数据结构的任何修
改，调用者函数都可以通过持有的 map 引用看到。在我们的例子中， countLines 函数
向 counts 插入的值，也会被 main 函数看到。*/

func countLines(s *os.File, counts map[string]map[string]int, filename string) {
	input := bufio.NewScanner(s)
	for input.Scan() {
		if input.Text() == "q" {
			break
		}
		//tmp := counts[input.Text()]
		//tmp[filename]++
		////counts[input.Text()] = tmp
		counts[input.Text()]  [filename]++
	}
}

func main() {
	counts := make(map[string]map[string]int)
	files := os.Args[1:]
	fmt.Println(files)

	if len(files) == 0 {
		//countLines(os.Stdin, counts)
		fmt.Println("/t/t-----plase specify the  file name!")
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
			}
			countLines(f, counts, arg)
			f.Close()
		}
	}
	// {43:{kk:2}
	fmt.Println("---------------------------")
	for line, n := range counts {
		//fmt.Printf("%s\t%s\n", n[], line)
		fmt.Println(line, n)
		//for k, v := range n {
		//	fmt.Println(k,v)
		//}

	}

}

// 执行 go run dup2_readfile.go ./kk
