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

	dataqueue := queue.New()

	go func() {
		fmt.Println("Dequeue Total Size: ", dataqueue.Size())

		for dataqueue.Size() > 0 {
			ele := dataqueue.Dequeue()
			fmt.Println(ele.Value, dataqueue.Size())

		}
	}()

	i := 0

	for i < 200 {
		i++
		dataqueue.Enqueue("Enqueue data" + strconv.Itoa(dataqueue.Size()))
		fmt.Println("Total Size: ", dataqueue.Size())

	}

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
