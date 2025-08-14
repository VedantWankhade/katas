package udp

import (
	"log"
	"net"
)

func udpClient() {
	conn, err := net.Dial("udp", "127.0.0.1:12000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	_, err = conn.Write([]byte("hello, world"))
	if err != nil {
		log.Fatal(err)
	}
}
