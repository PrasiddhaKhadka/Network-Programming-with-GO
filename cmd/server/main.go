package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	list, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	defer list.Close()

	for {
		conn, err := list.Accept()

		if err != nil {
			log.Println(err)
			continue
		}

		io.WriteString(conn, "Hello whats up!")
		fmt.Println(conn, "How is your day")
		fmt.Fprintf(conn, "%v", "Well! I hope")
		conn.Close()

	}

}
