package edyscoin

import (
	"io"
	"log"
	"net"
)

type HttpHandler struct {
	listener net.Listener
	conn     net.Conn
}

func (hh *HttpHandler) NewHandler(addr, port string) HttpHandler {
	// TODO should prolly only need one of these
	listener, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatal("could not open listener on localhost:8080 ->", err)
	}

	conn, err := net.Dial("tcp", addr + ":" + port)
	if err != nil {
		log.Fatal("could not connect to remote node ->", err)
	}

	return HttpHandler{listener, conn}
}

func (hh *HttpHandler) Listener() {
	for {
		conn, err := hh.listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go func (c net.Conn) {
			io.Copy(c, c)
			c.Close()
		}(conn)
	}
}
