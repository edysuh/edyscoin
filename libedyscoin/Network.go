package libedyscoin

import (
	// "io"
	"log"
	"net"
	"strconv"
)

var id = 0

type Node struct {
	ID    int
	peers []*net.Conn
}

func NewNode(laddr string) Node {
	n := Node{id, make([]*net.Conn, 0)}
	id++
	n.StartServer(laddr)
	return n
}

func (n *Node) StartServer(laddr string) {
	server, err := net.Listen("tcp", laddr)
	if err != nil {
		log.Fatalf("could not set up server at %s -> %v", laddr, err)
	}
	for {
		conn, err := server.Accept()
		if err != nil {
			log.Fatalf("could not accept connection from %v -> %v", conn, err)
		}
		n.peers = append(n.peers, &conn)
	}
}

func (n *Node) ConnectToRemote(raddr string) *net.Conn {
	conn, err := net.Dial("tcp", raddr)
	if err != nil {
		log.Fatalf("could not connect to remote address %v -> %v", conn, err)
	}
	return &conn
}

func (n *Node) Handshake(raddr string) {
	conn := n.ConnectToRemote(raddr)
	(*conn).Write([]byte(strconv.Itoa(n.ID)))
}
