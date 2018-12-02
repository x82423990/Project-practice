package main

import "fmt"

var graph = make(map[string]map[string]bool)

func  addEdge(from, to string)  {
	edge :=graph[from]
	if edge==nil{
		edge = make(map[string]bool)
		graph[from] = edge
	}
	edge[to] = true
}

func hasEdge(from, to string)bool {
	return graph[from][to]
}

func main()  {
	addEdge("a","b")
	fmt.Println(hasEdge("a","c"))
}