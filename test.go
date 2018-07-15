package main

import (
	"fmt"
	"strconv"
)
type student struct {
	id string
	name string
}

type class struct {
	students []student
}


func (l *class )add(s *student) {
	l.students = append(l.students, *s)
}


//func del(s []string, index int) string {
//
//	if index < len(s)-1 {
//		s = append(s[:index-1], s[index+1:]...)
//	}else{
//		fmt.Println("index out of range")
//
//
//	}
//	delSemnt := s[index-1]
//	return delSemnt
//}
////var a []string

func (c *class) del(id int) *student{
	if id <=0 || id > len(c.students){
		fmt.Println("index out of range")
		return nil
	}

	delval := &c.students[id-1]
	if id < len(c.students){
		fmt.Println("in id <len")
		c.students = append(c.students[:id-1], c.students[id:]...)
	} else if len(c.students) == 1{
		c.students = make([]student, 0)
	} else if id == len(c.students){
		fmt.Println("in id =len")
		c.students = c.students[:id-1]
	}else{
		fmt.Println("unknow err")
		return nil
	}

	return delval
}


var s1, s2,s3,s4,s5 student
var c class
var i =0
func main() {
	i++
	s1 = student{strconv.Itoa(i),"zhangxue"}
	//i++
	//s2 = student{strconv.Itoa(i),"zhanglei"}
	//i++
	//s3 = student{strconv.Itoa(i),"ersu"}
	//i++
	//s4 = student{strconv.Itoa(i),"wangyu"}
	//i++
	//s5 = student{strconv.Itoa(i),"hngd"}
	c.add(&s1)
	//c.add(&s2)
	//c.add(&s3)
	//c.add(&s4)
	//c.add(&s5)
	c.del(1)
	// fmt.Println(c.students[0:])
	//fmt.Println(c.students[:1])
	// c.students = append(c.students[:1], c.students[1:]...)
	fmt.Println(c.students)

}
