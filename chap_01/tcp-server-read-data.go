package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "8080"
	CONN_TYPE = "tcp"
)

func handleRequest(conn net.Conn) {
	message, err := bufio.NewReader(conn).ReadString('\n')

	if err != nil {
		log.Fatal("Error reading:", err.Error())
	}
	fmt.Println("Message Received from the client: ", string(message))
	conn.Close()
}

func main() {
	listener, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
	if err != nil {
		log.Fatal("Error starting tcp server : ", err)
	}
	defer listener.Close()
	log.Println("Listening on ", CONN_HOST+":"+CONN_PORT)
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Error Accepting: ", err.Error())
		}
		go handleRequest(conn)
	}
}
