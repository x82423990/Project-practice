package main

import "fmt"

func main() {
	mySlice := make([]int, 5, 10)
	fmt.Println("len(mySlice):", len(mySlice))
	fmt.Println("cap(mySlice):", cap(mySlice))
	mySlice = append(mySlice, 1, 2, 34, 5,5,5,5,5)
	fmt.Println("len(mySlice):", len(mySlice))
	fmt.Println("cap(mySlice):", cap(mySlice))
	// append outer array
	mySlice2 := []int{8, 9, 10,5,5,5,5,5}
	mySlice = append(mySlice, mySlice2...)
	fmt.Println("len(mySlice):", len(mySlice))
	fmt.Println("cap(mySlice):", cap(mySlice))

	// cap()
	oldSlice := []int{1,2,3,4,5}
	newSlice := oldSlice[:5]
	fmt.Println("cap oldSilce: ", cap(oldSlice))
	fmt.Println(newSlice)

}
