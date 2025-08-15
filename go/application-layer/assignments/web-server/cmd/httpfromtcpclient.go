package main

import (
	"bytes"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 4 {
		printUsaeg()
		os.Exit(1)
	}

	// skipping argument verification for now
	host := os.Args[1]
	port := os.Args[2]
	path := os.Args[3]

	conn, err := net.Dial("tcp", host+":"+port)
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	var req bytes.Buffer
	req.WriteString(fmt.Sprintf("GET %s HTTP/1.1\n", path))
	// req.WriteString("Connection: keep-alive\n")
	req.WriteString("Accept: text/html\n\r\n")
	_, err = conn.Write(req.Bytes())
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

func printUsaeg() {
	fmt.Printf("Usage:\nhttpfromtcpclient <host> <port> <path>\n")
}
