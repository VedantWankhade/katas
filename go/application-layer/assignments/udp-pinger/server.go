package udppinger

import (
	"log"
	"math/rand/v2"
	"net"
)

func server() {
	lis, err := net.ListenPacket("udp", "127.0.0.1:12000")
	if err != nil {
		log.Fatal(err)
	}
	defer lis.Close()
	for {
		req := make([]byte, 1024)
		_, addr, err := lis.ReadFrom(req)
		r := rand.IntN(10)
		// reject 30% requests
		if err != nil || r < 4 {
			continue
		}
		go func(retAddr net.Addr, req []byte) {
			conn, err := net.Dial(addr.Network(), addr.String())
			if err != nil {
				return
			}
			defer conn.Close()
			_, err = conn.Write(req)
		}(addr, req)
	}
}
