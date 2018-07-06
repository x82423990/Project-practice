package main

import "fmt"

func printAll(vals []interface{}) { //1
	for _, val := range vals {
		fmt.Println(val)
	}
}
type I interface {
	speak(s string)
}

type animal struct {
	name string
	have_leg bool
}

type dog struct {
	animal
	have_mouse bool
}
func ( a *animal) speak(s string)  {
	fmt.Println(s)
}

func main(){
	var i I
	i = &animal{"xiaogang", true}
	i.speak("wangwang")
	var d I
	d = &dog{animal{name:"wangcai", have_leg:true}, true}
	d.speak("waiwcai")
}