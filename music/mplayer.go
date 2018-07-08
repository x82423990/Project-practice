package main

import (
	"gotest/music/library"
	"fmt"
	"strconv"
	"bufio"
	"os"
	"strings"
	"gotest/music/mp"
)

var lib *library.MusicManger
var ret *library.MusicEntry
var id int = 0
var ctrl, signal chan int

func handleLibCommands(tokens []string){
	switch tokens[1] {
	case "list":
		for i:=0;i<lib.Len();i++{
			e,_ :=  lib.Get(i)
			fmt.Println(e.Id, ":", e.Name, e.Artist, e.Source, e.Type, e.Genre)
		}
	case "add":
		fmt.Println("enter add()")
		if len(tokens)==7{
			id ++
			//lib add HugeStone MJ POP ~/MusicLib/hs.mp3 MP3
			lib.Add(&library.MusicEntry{strconv.Itoa(id), tokens[2], tokens[3],
			tokens[4],tokens[5],  tokens[6]})
			fmt.Printf("add %v success\n", tokens[2] )
		}else{
			fmt.Println("USAGE: lib add <name><artist><source><type>")
		}
	case "find":
		if len(tokens) ==3{
			ret := lib.Find(tokens[2])
			if ret !=nil{
				fmt.Println(ret.Id,ret.Name)
			}else{
				fmt.Println("not found")
			}
		}

	case "remove":
		if len(tokens) ==3{
			index := tokens[2]
			//lib.Remove(index)
			//if v, ok := interface{}(index).(int);ok{
			//	fmt.Println("this is number")
			//	ret = lib.Remove(v)
			//}else {
			tmp, _:=strconv.Atoi(lib.Find(index).Id)
			fmt.Println(tmp)
			ret = lib.Remove(tmp)

				 // fmt.Println()

			}
			fmt.Println(ret)
		}
	}

func handlePlayCommand(tokens []string){
	if len(tokens) !=2{
		fmt.Println("USAGE: play <name>")
		return
	}
	e := lib.Find(tokens[1])
	if e ==nil{
		fmt.Println("The music", tokens[1], "does not exist.")
		return
	}
	mp.Play(e.Source, e.Type)
}
func main(){
	fmt.Println(`
		Enter following commands to control the player:
		lib list -- View the existing music lib
		lib add <name><artist><source><type> -- Add a music to the music lib
		lib remove <name> -- Remove the specified music from the lib
		play <name> -- Play the specified music
	`)
	lib = library.NewMusicManager()
	r := bufio.NewReader(os.Stdin)
	for{
		fmt.Println("Enter Command->")
		rewLine,_,_:=r.ReadLine()
		line := string(rewLine)
		if line =="q" || line =="e"{
			break
		}
		tokens :=strings.Split(line, " ")
		if tokens[0] == "lib"{
			handleLibCommands(tokens)
		}else if tokens[0] == "play"{
			handlePlayCommand(tokens)
		}else{
			fmt.Println("Unrecognized command:", tokens[0])
		}
	}

}
