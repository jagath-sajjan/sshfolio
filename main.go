package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	port:= os.Getenv("PORT")
	if port == "" {
			port = "22097"
	}

	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Listening on port", port)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		go func(c net.Conn) {
			defer c.Close()
			c.Write([]byte("Holaa from jogostation\n"))
		}(conn)
	}
}
