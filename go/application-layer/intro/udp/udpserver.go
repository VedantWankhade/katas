package udp

import (
	"fmt"
	"log"
	"net"
	"strings"
)

func udpServer() {
	conn, err := net.ListenPacket("udp", "127.0.0.1:12000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	var msg []byte = make([]byte, 1024)
	_, retAddr, err := conn.ReadFrom(msg)
	msgString := string(msg)
	fmt.Println(msgString)
	fmt.Println("got from", retAddr.String())

	retConn, err := net.Dial(retAddr.Network(), retAddr.String())
	if err != nil {
		log.Fatal(err)
	}

	defer retConn.Close()
	ret := []byte(strings.ToUpper(msgString))
	retConn.Write(ret)
}
