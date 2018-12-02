package main

import "time"

type Car struct {
	name string
	age  int
}

type Train struct {
	Car
	int
	start time.Time
}

func main() {
	var t Train
	t.name = "Train" 
	t.age = 100
	t.int = 200
}
