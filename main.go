package main

import (
	"github.com/sivsivsree/go-now/basics"
	"github.com/sivsivsree/go-now/hello"
	"github.com/sivsivsree/go-now/lib"
	"github.com/sivsivsree/go-now/lib/datastructures"
	"github.com/sivsivsree/go-now/lib/structures"
	"github.com/sivsivsree/go-now/lib/gui"
)

func main() {
	hello.World()
	basics.Init()
	lib.JSONParse()
	structures.MakeStruct()
	datastructures.TestLikedList()

	gui.ScreenOne()

}
