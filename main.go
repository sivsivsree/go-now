package main

import (
	"./lib"
	"./lib/datastructures"
	"./lib/structures"
	"github.com/sivsivsree/go-now/basics"
	"github.com/sivsivsree/go-now/hello"
)

func main() {
	hello.World()
	basics.Init()
	// lib.Create()
	lib.JSONParse()
	structures.MakeStruct()
	datastructures.TestLikedList()

}
