package tcp

import (
	"fmt"
	"log"
	"net"
)

func tcpClient() {
	conn, err := net.Dial("tcp", "127.0.0.1:8090")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	_, err = conn.Write([]byte("hello, world!"))
	if err != nil {
		log.Fatal(err)
	}
	res := make([]byte, 1024)
	_, err = conn.Read(res)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(res))
}
