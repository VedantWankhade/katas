package webserver

import (
	"crypto/tls"
	"fmt"
	"log"
	"net"
	"strings"
)

func httpFromTcpServerTLS() {

	cert, err := tls.LoadX509KeyPair("cert.pem", "key.pem")
	if err != nil {
		log.Fatal(err)
	}

	cfg := &tls.Config{Certificates: []tls.Certificate{cert}}

	lis, err := tls.Listen("tcp", "127.0.0.1:8090", cfg)
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := lis.Accept()
		if err != nil {
			log.Println("error accepting the connection:", err)
			continue
		}
		go func(c net.Conn) {
			msg := make([]byte, 1024)
			_, err := conn.Read(msg)
			if err != nil {
				log.Println("error reading message:", err)
				return
			}
			fmt.Println(string(msg))

			lines := strings.Split(string(msg), "\n")
			if len(lines) < 1 {
				log.Println("something wrong with the http message:", err)
				return
			}

			httpLine := lines[0]
			path := strings.Split(httpLine, " ")[1]

			res, err := getResForPath(path)
			if err != nil {
				log.Println("error getting response for", path, ":", err)
				return
			}

			conn.Write(res.Bytes())
			conn.Close()
		}(conn)
	}
}

// func getResForPath(path string) (*bytes.Buffer, error) {
// 	var res *bytes.Buffer
// 	var err error
//
// 	switch path {
// 	case "/", "/hello":
// 		res, err = getHello()
// 	case "/favicon.ico":
// 		res, err = getFavicon()
// 	default:
// 		res, err = nil, fmt.Errorf("invalid path: %s: %d", path, http.StatusNotFound)
// 	}
// 	return res, err
// }
//
// func getFavicon() (*bytes.Buffer, error) {
// 	var res bytes.Buffer
// 	var resBody bytes.Buffer
//
// 	ico, err := os.Open("favicon.png")
// 	if err != nil {
// 		return nil, err
// 	}
// 	_, err = io.Copy(&resBody, ico)
// 	if err != nil {
// 		return nil, err
// 	}
// 	res.WriteString("HTTP/1.1 200 OK\n")
// 	res.WriteString("Content-Type: image/png\n")
// 	res.WriteString(fmt.Sprintf("Content-Length: %d\n", resBody.Len()))
// 	res.WriteString("\r\n")
// 	res.WriteString(resBody.String())
//
// 	return &res, nil
// }
//
// func getHello() (*bytes.Buffer, error) {
// 	var res bytes.Buffer
// 	var resBody bytes.Buffer
//
// 	file, err := os.Open("hello.html")
// 	if err != nil {
// 		return nil, err
// 	}
// 	_, err = io.Copy(&resBody, file)
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	res.WriteString("HTTP/1.1 200 OK\n")
// 	res.WriteString("Content-Type: text/html\n")
// 	res.WriteString("Connection: close\n")
// 	res.WriteString(fmt.Sprintf("Content-Length: %d\n", resBody.Len()))
// 	res.WriteString("\r\n")
// 	res.WriteString(resBody.String())
//
// 	return &res, nil
// }
