package libedyscoin

import (
	// "fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

type Node struct {
	Id          Id
	Address		string
	Peers       map[Id]string
	// HandshakeCh chan HandshakeRequest
}

func NewNode(laddr string) *Node {
	n := new(Node)
	n.Id       = NewId(laddr)
	n.Address  = laddr
	n.Peers    = make(map[Id]string)
	// n.HandshakeCh = make(chan HandshakeRequest)

	rpc.Register(&RpcService{n})
	rpc.HandleHTTP()
	n.StartServer(laddr)

	return n
}

func (n *Node) StartServer(laddr string) {
	listener, err := net.Listen("tcp", laddr)
	if err != nil {
		log.Fatalf("could not set up listener at %s -> %v", laddr, err)
	}
	go http.Serve(listener, nil)
}

func (n *Node) ConnectToRemote(raddr string) *rpc.Client {
	client, err := rpc.DialHTTP("tcp", raddr)
	if err != nil {
		log.Fatalf("could not dial remote address %v -> %v", client, err)
	}
	return client
}

// func (n *Node) RequestHandler() {
// 	select {
// 	case req := <-n.HandshakeCh:
// 		fmt.Printf("%v\n", req)
// 	}
// }

func (n *Node) DoHandshake(raddr string) (*HandshakeResponse, error) {
	// fmt.Printf("from hs: %v, %v\n", n.Id, n.Address)

	client := n.ConnectToRemote(raddr)
	req := HandshakeRequest{n.Id, n.Address}
	var res HandshakeResponse

	err := client.Call("RpcService.Handshake", req, &res)
	if err != nil {
		log.Fatal("DoHandshake Error: ", err)
	}

	// fmt.Printf("response from DoHandshake: %v\n", res)
	return &res, nil
}

func (n *Node) DoHandshakeWithId(Id Id) {
	raddr := n.Peers[Id]
	n.DoHandshake(raddr)
}
