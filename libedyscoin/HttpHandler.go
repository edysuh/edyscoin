package libedyscoin

import (
	// "io"
	"log"
	"net"
)

type HttpHandler struct {
	peers []net.Conn
}

func (hh *HttpHandler) NewHandler(port string) {
	listener, err := net.Listen("tcp", ":" + port)
	if err != nil {
		log.Fatalf("could not open listener on localhost:%s -> %s", port, err)
	}

	log.Printf("listening on localhost port %s\n", port)

	for {
		conn, err := listener.Accept()
		log.Printf("\nconnection accepted: %v", conn.RemoteAddr())
		if err != nil {
			log.Fatal("could not accept connection ->", err)
		}
		hh.peers = append(hh.peers, conn)
	}
}

func (hh *HttpHandler) Broadcast(msg string) {
	for _, peer := range hh.peers {
		go peer.Write(([]byte)(msg))
	}
}

func Connect(addrport string) {
	_, err := net.Dial("tcp", addrport)
	if err != nil {
		log.Fatal("could not connect to remote node -> ", err)
	}
	log.Printf("connected to remote node %s", addrport)
}
