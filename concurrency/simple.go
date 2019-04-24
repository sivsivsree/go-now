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

func UsingRange() {

	c := make(chan string)
	go count("With Range", c)

	for msg := range c {
		fmt.Println(msg)
	}
}

func BufferedChannel() {

	c := make(chan string, 2)

	c <- "One"
	msg := <-c
	fmt.Println(msg)

}

func UsingSelect() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			c1 <- "Every Millisecond"
			time.Sleep(time.Millisecond * 500)
		}

	}()

	go func() {

		for {
			c2 <- "Every 2 Seconds"
			time.Sleep(time.Second * 2)
		}

	}()

	for {
		select {
		case ms1 := <-c1:
			fmt.Println(ms1)
		case ms2 := <-c2:
			fmt.Println(ms2)

		}
	}

}

func WorkerPulls() {

	jobs := make(chan int, 100)
	results := make(chan int, 100)

	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)

	for i := 0; i < 100; i++ {
		jobs <- i
	}

	close(jobs)

	for i := 0; i < 100; i++ {
		fmt.Println(<-results)
	}
}

func worker(jobs <-chan int, results chan<- int) {

	for n := range jobs {
		results <- fib(n)
	}

}

func fib(n int) int {

	if n <= 1 {
		return n
	}

	return fib(n-1) + fib(n-2)
}

func count(thing string, c chan string) {
	for i := 1; i < 5; i++ {
		c <- thing
		time.Sleep(time.Millisecond * 500)
	}

	close(c)
}
