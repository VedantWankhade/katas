package tcp

import (
	"fmt"
	"log"
	"net"
	"strings"
)

func tcpServer() {
	lis, err := net.Listen("tcp", "127.0.0.1:8090")
	if err != nil {
		log.Fatal(err)
	}
	defer lis.Close()
	conn, err := lis.Accept()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	msg := make([]byte, 1024)
	_, err = conn.Read(msg)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(msg))
	fmt.Println("received from", conn.RemoteAddr())

	msgString := strings.ToUpper(string(msg))

	// respond
	_, err = conn.Write([]byte(msgString))
	if err != nil {
		log.Fatal(err)
	}
}
