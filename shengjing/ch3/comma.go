package main

import (
	"bytes"
	"strings"
	"fmt"
)

func AddPoint(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	//12345
	return AddPoint(s[:n-3]) + "," + s[n-3:]
}

func AddPointNon(s string, rever bool) string {
	var test bytes.Buffer
	lens := len(s)
	for i := lens; i > 0; i-- {
		if (i)%3 == 0 && i != lens {
			test.WriteString(",")
		}
		if rever {
			test.WriteByte(s[lens-i])
		} else {
			test.WriteByte(s[i-1])
		}
	}
	return test.String()
}

func AddPlus(s string) string {
	var test bytes.Buffer
	slash := strings.HasPrefix(s, "-")
	if slash {
		s = s[1:]
		test.WriteString("-")
	}
	ret := strings.Contains(s, ".")
	if ret {
		n := strings.Split(s, ".")
		n1 := AddPointNon(n[0], true)
		n2 := AddPointNon(n[1], false)
		test.WriteString(n1)
		test.WriteString(".")
		test.WriteString(n2)
		return test.String()
	} else {
		test.WriteString(s)
		return test.String()
	}

}

func IsSame(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	//for _,v := range []byte(s1){
	//	if ! bytes.Contains([]byte(s2),[]byte{v}){
	//		return false
	//	}
	//}
	for i := 0; i < len(s1); i++ {
		if ! strings.Contains(s2, string(s1[i])) {
			return false
		}
	}
	return true
}

func main() {
	s := "-1234.3311"
	//fmt.Println(AddPoint(s))
	// AddPointNon(s)
	fmt.Println(AddPlus(s))
	// fmt.Println(tmp)
	/// s = 123456
	//fmt.Println(s[:4])
	//fmt.Println(s[4:])
	a1 := "123abc"
	a2 := "abc123"
	fmt.Println(IsSame(a1, a2))
	bu := "123"
	a3 := bu
	a3="bibi"
	fmt.Println(&a3)
	fmt.Println(&bu)
	bu = "234"
	fmt.Println(&bu, bu,a3)
	// strings 提供的6个函数
	//a := "xieyifiiiiian"
	//b :="yifan"
	//c:="xie"
	//fmt.Println(strings.Contains(a, b))
	//fmt.Println(strings.HasPrefix(a,c))
	//fmt.Println(strings.HasSuffix(a,b))
	//fmt.Println(strings.Index(a,c))
	//fmt.Println(strings.Fields(a))
	//fmt.Println(strings.Count(a,"i"))

}
