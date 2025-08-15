package udp

import (
	"fmt"
	"log"
	"net"
)

func udpClient() {
	conn, err := net.Dial("udp", "127.0.0.1:12000")
	if err != nil {
		log.Fatal(err)
	}
	_, err = conn.Write([]byte("hello, world"))
	if err != nil {
		log.Fatal(err)
	}

	conn.Close()
	lis, err := net.ListenPacket("udp", conn.LocalAddr().String())

	if err != nil {
		log.Fatal(err)
	}
	msg := make([]byte, 1024)
	_, _, err = lis.ReadFrom(msg)
	fmt.Println(string(msg))
}
