package concurrency

import (
	"fmt"
	"time"
)

func Simple() {

	c := make(chan string)
	go count("Sheep", c)

	for {
		msg, open := <-c

		if !open {
			break
		}
		fmt.Println(msg)
	}
}

func count(thing string, c chan string) {
	for i := 1; i < 5; i++ {
		c <- thing
		time.Sleep(time.Millisecond * 500)
	}

	close(c)
}
