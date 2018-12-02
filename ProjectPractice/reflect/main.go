package main

import (
	"fmt"
	"reflect"
)

type NotKnowType struct {
	Name string `json:"student_name"`
	Age  int
}

func (p NotKnowType) GetName() {
	fmt.Printf("Hi, I'm is %s", p.Name)
}

func (p NotKnowType) SetAge(age int) (int) {
	p.Age = age
	return p.Age
}

func tesetStruct(a interface{}) {
	value := reflect.ValueOf(a)
	typ := reflect.TypeOf(a)
	fmt.Println(value)
	kd := value.Kind()
	if kd != reflect.Ptr && value.Elem().Kind() == reflect.Struct {
		fmt.Println("expect struct")
		return
	}
	//fmt.Printf("the has %d method\n", value.Elem().NumField())
	//fmt.Printf("type is %s\n", reflect.TypeOf(value))
	//fmt.Printf("kind is %s\n", value.Elem().Kind())
	num := value.Elem().NumField()
	//value.Elem().Field(0).SetString("yifanfenshun")
	value.Elem().Field(0).SetString("stu1000")
	// tag := typ.Elem().Field(0).Tag.Get("json")
	tag := typ.Elem().Field(0).Tag.Get("json")

	fmt.Printf("tag=%s\n", tag)
	for i := 0; i < num; i++ {
		fmt.Printf("filed %d is %v\n", i, value.Elem().Field(i))
	}
	var parmrs []reflect.Value
	//parmrs = append(parmrs, 30)
	fmt.Println(value.Elem().Method(0))
	value.Elem().Method(0).Call(parmrs)
	//fmt.Println(results)
	fmt.Println(a)
}
func main() {
	var a NotKnowType = NotKnowType{Name: "yifan", Age: 20}
	tesetStruct(&a)
}
