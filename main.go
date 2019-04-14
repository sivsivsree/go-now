package main

import (
	"github.com/sivsivsree/go-now/basics"
	"github.com/sivsivsree/go-now/hello"
	"github.com/sivsivsree/go-now/lib"
	"github.com/sivsivsree/go-now/lib/datastructures"
	"github.com/sivsivsree/go-now/lib/structures"
	"github.com/sivsivsree/go-now/web"
)

func main() {
	hello.World()

	// gui.ScreenOne()

	basics.Init()

	// lib.CreateSomething()
	lib.JSONParse()

	structures.MakeStruct()

	defer datastructures.TestLikedList()

	web.HandleTCP()

}
