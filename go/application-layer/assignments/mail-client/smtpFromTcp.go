package mailclient

import (
	"fmt"
	"log"
	"net"
)

func smtpFromTcp() {
	conn, err := net.Dial("tcp", "127.0.0.1:1025")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	buffer := make([]byte, 1024)
	_, err = conn.Read(buffer)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(buffer))

	conn.Write([]byte("HELO test.com\r\n"))
	buffer = make([]byte, 1024)
	conn.Read(buffer)
	fmt.Println(string(buffer))

	conn.Write([]byte("MAIL FROM: <me@test.com>\r\n"))
	buffer = make([]byte, 1024)
	conn.Read(buffer)
	fmt.Println(string(buffer))

	conn.Write([]byte("RCPT TO: <you@test.com>\r\n"))
	buffer = make([]byte, 1024)
	conn.Read(buffer)
	fmt.Println(string(buffer))

	conn.Write([]byte("DATA\r\n"))
	buffer = make([]byte, 1024)
	conn.Read(buffer)
	fmt.Println(string(buffer))

	conn.Write([]byte("Received: from yout@test.com by me@test.com ; Thu, 21 May 1998\nFrom: ya boy\nTo: ma boy\nSubject: mr sir guy boy\r\n\r\nYo hi dude!\nme\r\n.\r\n"))
	buffer = make([]byte, 1024)
	conn.Read(buffer)
	fmt.Println(string(buffer))

	conn.Write([]byte("QUIT\r\n"))
	buffer = make([]byte, 1024)
	conn.Read(buffer)
	fmt.Println(string(buffer))

}
