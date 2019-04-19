package main

import (
	"github.com/joho/godotenv"
	"github.com/sivsivsree/go-now/hello"
	"github.com/sivsivsree/go-now/web"
	"log"
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

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	go web.Server()
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	// fmt.Println(<-pongs)

	web.HandleTCP()

}

func ping(pings chan<- string, msg string) {
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}
