package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/sivsivsree/go-now/hello"
	"github.com/sivsivsree/go-now/web"
	"log"
	"sync"
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

	//for i := 0; i < 1000; i++ {
	//	 generateMessage("Message String I " + fmt.Sprintf("%v", i))
	//	 generateMessage("Message String II " + fmt.Sprintf("%v", i))
	//	 generateMessage("Message String III " + fmt.Sprintf("%v", i))
	//}

	go web.Server()

	// fmt.Println(<-pongs)

	web.HandleTCP()

}

var mutex = &sync.Mutex{}
var wg sync.WaitGroup

func generateMessage(message string) {
	//...

	wg.Add(1)
	go func() {
		defer wg.Done()
		printMessage(message)
	}()

	wg.Wait()
}

func printMessage(resultMessage string) {
	mutex.Lock()
	defer mutex.Unlock()

	// Internal Logic : Ignore
	//if IsTimeEnabled {
	//	log.Println(resultMessage)
	//	return
	//}
	fmt.Println(resultMessage)
}
