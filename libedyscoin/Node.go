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
}

func NewNode(laddr string) *Node {
	n := new(Node)
	n.Id       = NewId(laddr)
	n.Address  = laddr
	n.Peers    = make(map[Id]string)

	rpc.Register(&RpcService{n})
	rpc.HandleHTTP()
	n.StartServer(laddr)

	return n
}

func (n *Node) StartServer(laddr string) {
	listener, err := net.Listen("tcp", laddr)
	if err != nil {
		log.Fatalf("could not set up listener at %s: %v", laddr, err)
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

func (n *Node) RpcCall(raddr string, method string, req *Message) (*Message, error) {
	client, err := n.ConnectToRemote(raddr)
	if err != nil {
		return nil, err
	}
	var res Message
	call := client.Go("RpcService." + method, req, &res, nil)
	callres := <-call.Done
	if callres.Error != nil {
		log.Fatal(callres.Error)
	}
	return &res, nil
}

func (n *Node) DoHandshake(raddr string) (*Message, error) {
	req := &Message{SenderId: n.Id, SenderAddr: n.Address}
	res, err := n.RpcCall(raddr, "Handshake", req)
	if err != nil {
		log.Fatal(err)
	}
	if !n.Id.Equals(res.SenderId) {
		n.Peers[res.SenderId] = res.SenderAddr
	}
	return res, err
}

// TODO should return some sort of success i.e. the array of nodes broadcasted to
func (n *Node) DoBroadcast(req *Message) {
	req.Seenlist[n.Id] = true
	for rid, raddr := range n.Peers {
		go func(rid Id, raddr string) {
			if !n.Id.Equals(rid) && !req.Seenlist[rid] {
				_, err := n.RpcCall(raddr, "Broadcast", req)
				if err != nil {
					fmt.Println(err)
				}
			}
		}(rid, raddr)
	}
}

func (n *Node) DoBroadcastNewTransaction(txn *Transaction) (*Message, error) {

	return nil, nil
}

func (n *Node) DoBroadcastNewBlockChain(bc *BlockChain) (*Message, error) {
	return nil, nil
}
