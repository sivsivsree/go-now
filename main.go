package main

import (
<<<<<<< HEAD
<<<<<<< HEAD
=======
	"fmt"
	"github.com/sivsivsree/practice/basics"
>>>>>>> 384719b81fc20aa1502044ee0e7796e589c47936
	"github.com/sivsivsree/practice/hello"
=======
	"./lib"
	"./lib/datastructures"
	"./lib/structures"
	"github.com/sivsivsree/go-now/basics"
	"github.com/sivsivsree/go-now/hello"
>>>>>>> 9377ecb1d8477c0a0057613d4b673acbf6933fa8
)

func main() {
	hello.World()
	basics.Init()
	// lib.Create()
	lib.JSONParse()
	structures.MakeStruct()
	datastructures.TestLikedList()

}
