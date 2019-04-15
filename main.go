package main

import (
	"github.com/sivsivsree/go-now/hello"
	"github.com/sivsivsree/go-now/web"
)

func main() {
	hello.World()

	// gui.ScreenOne()

	/*
		basics.Init()


			lib.JSONParse()

			structures.MakeStruct()

			defer datastructures.TestLikedList()
	*/

	go web.HandleTCP()
	web.Server()

}
