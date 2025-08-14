package udp

import (
	"fmt"
	"log"
	"net"
)

func udpServer() {
	conn, err := net.ListenPacket("udp", "127.0.0.1:12000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	var msg []byte = make([]byte, 1024)
	_, _, err = conn.ReadFrom(msg)
	fmt.Println(string(msg))
}
