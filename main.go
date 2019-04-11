package main

import (
	"./lib"
	"./lib/datastructures"
	"./lib/structures"
	"github.com/sivsivsree/practice/basics"
	"github.com/sivsivsree/practice/hello"
)

func main() {
	hello.World()
	basics.Init()
	// lib.Create()
	lib.JSONParse()
	structures.MakeStruct()
	datastructures.TestLikedList()

}
