package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {

	fmt.Printf("pid: %d\n", os.Getpid())

	li, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Println(err)
	}

	defer li.Close()

	for {
		conn, err := li.Accept()
		fmt.Printf("pid: %d \n", os.Getpid())
		fmt.Printf("Serving %s \n", conn.RemoteAddr().String())

		if err != nil {
			log.Println(err)
		}

		go handleAllConnection(conn)

		defer conn.Close()

	}

}

func handleAllConnection(conn net.Conn) {
	for {
		io.WriteString(conn, "Hello from TCP server..")
	}
}
