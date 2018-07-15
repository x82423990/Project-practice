package main

import (
	"flag"
	"fmt"
)

func feb(n int) int {
	x, y :=0,1
	for i:=0;i<n;i++{
		x,y = y, x+y
		fmt.Printf("-------------%d----------------------\n",i )
		println(x, y)
	}
	return x
}
var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", "", "separator")
func main()  {
	//flag.Parse()
	//// println(flag.Args())
	//fmt.Println(strings.Join(flag.Args(), *sep))
	//if !*n{
	//	fmt.Println("-----------")
	//}
	//fmt.Println(feb(24))
	feb(8)



}
