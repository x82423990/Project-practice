package main

import (
	"gotest/music/library"
	"fmt"
)

var lib *library.MusicManger
var id int = 1
var ctrl, signal chan int

func handleLibCommands(token []string){
	switch token[1] {
	case "list":
		for i:=0;i<lib.Len();i++{
			e,_ :=  lib.Get(i)
			fmt.Println(i+1, ":", e.Name, e.Artist, e.Source, e.Type, e.Genre)
		}
	}
}