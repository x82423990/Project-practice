package main

import (
	"bytes"
	"fmt"
	"reflect"
)

func intToString(value []int) string{
	var buf bytes.Buffer
	buf.WriteByte('[')
	for k,v :=range value{
		if k>0{
			buf.WriteString(",")
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte('\n')
	// buf.WriteByte(']')
	buf.WriteRune(']')
	return buf.String()
}

func main()  {
	tmp := intToString([]int{1,2,3})
	fmt.Println(tmp)
	fmt.Println(reflect.TypeOf(tmp))
	
}
