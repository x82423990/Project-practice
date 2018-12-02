package main

import "fmt"

func rev(s []string){
	for i,j :=0,len(s)-1;i<j;i, j= i+1,j-1{
		s[i], s[j] = s[j],s[i]
	}
}
func main()  {
	s :=[]string{"a","b","c","d","e"}
	rev(s)
	fmt.Println(s)
}
