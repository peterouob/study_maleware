package main

import (
	"bufio"
	"log"
	"net"
)

func echo(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)
	s, err := reader.ReadString('\n')
	if err != nil {
		log.Println("Unable to read data")
	}
	log.Printf("Read %d bytes : %s", len(s), s)
	log.Println("Writing Data")
	writer := bufio.NewWriter(conn)
	if _, err := writer.WriteString(s); err != nil {
		log.Println("Unable to write data")
	}
	writer.Flush()
}

func main() {
	listener, err := net.Listen("tcp", ":20080")
	if err != nil {
		log.Println("Unable to listen tcp port")
	}
	log.Println("Listen on 0.0.0.0:20080")
	for {
		conn, err := listener.Accept()
		log.Println("Received connect")
		if err != nil {
			log.Println("Unable to accept server")
		}
		go echo(conn)
	}
}
