package main

import (
	"unicode"
	"fmt"
)

//4.3： 重写reverse函数，使用数组指针代替slice
func reverse(s *[5]string) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

//写一个函数在原地完成消除[]string中相邻重复的字符串的操作。
func eliminate(s []string) []string {
	count := 0
	for i := 1; i < len(s)-1; i++ {
		if s[i-1] == s[i] {
			s[i] = s[i+1]
			count ++
		}
	}
	return s[:len(s)-count+1]
}

//编写一个函数，原地将一个UTF-8编码的[]byte类型的slice中相邻的空格（参考
//unicode.IsSpace）替换成一个空格返回
func eliminateSpace(s []byte) []byte {
	// sum := len(s)
	// fmt.Println(sum)
	for i, _ := range s {
		fmt.Println("before",i)
		if i+1 > len(s) {
			fmt.Println(i)
			break
		}
		if unicode.IsSpace(rune(s[i])) && unicode.IsSpace(rune(s[i+1])) {
			copy(s[i:], s[i+1:])
			s = s[:len(s)-1]
			// i--
		}
		fmt.Println("affter",i)
		fmt.Println("-----------------")
	}
	fmt.Println(s)
	return s
}

func emptyString2(s []byte) []byte {
	index := 0
	num := len(s)
	for i := 0; i < num; i++ {
		index = i + 1
		num = len(s)
		if index >= num {
			break
		}
		if unicode.IsSpace(rune(s[i])) && unicode.IsSpace(rune(s[index])) {
			//结合remove函数
			copy(s[i:], s[index:])
			s = s[:len(s)-1]
			i--
		}
	}
	fmt.Println(s)
	return s
}

func main() {
	//a := [5]string{"a","b","b","d","e"}
	//b := []string{"a","b","b","d","e","e"}
	//reverse(&a)
	//fmt.Println(a)
	//tmp := eliminate(b)
	//fmt.Println(tmp)
	c := []byte{'a', ' ', ' ', 'c', ' ', ' ', 'd', ' ', ' ','f'}
	d := []byte{'a', ' ', ' ', 'c', ' ', ' ', 'd', ' ', ' ','f'}
	eliminateSpace(c)
	emptyString2(d)
}
