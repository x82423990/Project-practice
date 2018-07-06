package main

import "fmt"

type PersonInfo struct {
	ID      string
	Name    string
	Address string
}

func main() {
	personDB := make(map[string]PersonInfo)
	personDB["12345"] = PersonInfo{"123", "liling", "YangZiShan"}
	personDB["1"] = PersonInfo{"1", "Jack", "Room 101,..."}
	person, ok := personDB["12345"]
	if ok {
		fmt.Println("FOUND Persion", person)
	}else {
		fmt.Println("not find person with id 1234")
	}
	// 删除
	delete(personDB, "1")
	fmt.Println(personDB)
}
