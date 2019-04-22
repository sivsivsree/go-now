package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/sivsivsree/go-now/hello"
	"github.com/sivsivsree/go-now/lib/datastructures/queue"
	"log"
	"strconv"
)

func main() {
	hello.World()

	// gui.ScreenOne()

	/*
		basics.Init()


			lib.JSONParse()

			structures.MakeStruct()

			defer datastructures.TestLikedList()

		go web.Server()



		web.HandleTCP()
	*/

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	messages := make(chan string, 1)
	dataqueue := queue.New()

	println(dataqueue.Size())

	go func() {
		for {
			msg := <-messages
			fmt.Println("Recived := ", msg)
			if dataqueue.Size() > 0 {
				ele := dataqueue.Dequeue()
				fmt.Println("Dequeue Total Size: ", ele.Value, dataqueue.Size())
			}

		}
	}()

	go func() {
		i := 0

		for {
			i++
			messages <- "<--- Send for enqueue" + strconv.Itoa(i)
			dataqueue.Enqueue("randomdata")
			fmt.Println("Enqueue Total Size: ", dataqueue.Size())

		}
	}()

	var input string
	fmt.Scanf("%s", &input)
	// when this is ≠ that

}

//var mutex = &sync.Mutex{}
//var wg sync.WaitGroup
//
//func generateMessage(message string) {
//	//...
//
//	wg.Add(1)
//	go func() {
//		defer wg.Done()
//		printMessage(message)
//	}()
//
//	wg.Wait()
//}
//
//func printMessage(resultMessage string) {
//	mutex.Lock()
//	defer mutex.Unlock()
//
//	// Internal Logic : Ignore
//	//if IsTimeEnabled {
//	//	log.Println(resultMessage)
//	//	return
//	//}
//	fmt.Println(resultMessage)
//}
