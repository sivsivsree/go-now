package practices

import (
	"bufio"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func input(x *string) *string {

	reader := bufio.NewReader(os.Stdin)
	*x, _ = reader.ReadString('\n')

	return x
}

func main() {
	World()

	c := "yes?"
	input(&c)

	fmt.Println(c)

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

	Example()

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
