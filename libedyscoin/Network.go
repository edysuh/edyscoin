package libedyscoin

import (
	"fmt"
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

func (n *Node) ConnectToRemote(raddr string) (*rpc.Client, error) {
	client, err := rpc.DialHTTP("tcp", raddr)
	if err != nil {
		err = fmt.Errorf("remote node offline or incorrect address: %v\n", err)
	}
	return client, err
}

// func (n *Node) RequestHandler() {
// 	select {
// 	case req := <-n.HandshakeCh:
// 		fmt.Printf("%v\n", req)
// 	}
// }

func (n *Node) DoHandshake(raddr string) (*HandshakeResponse, error) {
	client, err := n.ConnectToRemote(raddr)
	if err != nil {
		return nil, err
	}
	req := HandshakeRequest{n.Id, n.Address}
	var res HandshakeResponse

	handshakeCall := client.Go("RpcService.Handshake", req, &res, nil)
	resCall := <-handshakeCall.Done
	if resCall.Error != nil {
		log.Fatal(resCall.Error)
	}

	return &res, nil
}

// TODO how do do lookups for nodes outside peer list?
func (n *Node) DoHandshakeWithId(Id Id) (*HandshakeResponse, error) {
	if raddr, ok := n.Peers[Id]; ok {
		return n.DoHandshake(raddr)
	}
	return nil, fmt.Errorf("node with this address is not a peer")
}

func (n *Node) DoBroadcast(method string, args interface{}) (*BroadcastResponse, error) {
	return nil, nil
}

func (n *Node) DoBroadcastNewTransaction(txn *Transaction) (*BroadcastResponse, error) {

	return nil, nil
}

func (n *Node) DoBroadcastNewBlockChain(bc *BlockChain) (*BroadcastResponse, error) {
	return nil, nil
}
