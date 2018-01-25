package libedyscoin

import (
	"encoding/json"
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
	BlockChain  *BlockChain
}

func NewNode(laddr string) *Node {
	n := new(Node)
	n.Id       = NewId(laddr)
	n.Address  = laddr
	n.Peers    = make(map[Id]string)
	n.BlockChain    = NewBlockChain()

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

func (n *Node) DoBroadcast(req *Message) ([]Id, error) {
	req.Params.Seenlist[n.Id] = true
	for rid, raddr := range n.Peers {
		// go TODO goroutine: should return the array of nodes broadcasted to
		func(rid Id, raddr string) {
			if !n.Id.Equals(rid) && !req.Params.Seenlist[rid] {
				res, err := n.RpcCall(raddr, "Broadcast", req)
				if err != nil {
					fmt.Println(err)
				}
				req.Params.Success = append(req.Params.Success, res.SenderId)
			}
		}(rid, raddr)
	}
	return req.Params.Success, nil
}

func (n *Node) DoBroadcastNewTransaction(txn *Transaction) (*Message, error) {
	marsh, err := json.Marshal(txn)
	if err != nil {
		log.Fatal("error in marshalling into json ->", err)
	}
	n.DoBroadcast(NewMessage(n, Params{Payload: marsh}))
	return nil, nil
}

func (n *Node) DoBroadcastNewBlockChain(bc *BlockChain) (*Message, error) {
	marsh, err := json.Marshal(bc)
	if err != nil {
		log.Fatal("error in marshalling into json ->", err)
	}
	n.DoBroadcast(NewMessage(n, Params{Payload: marsh}))
	return nil, nil
}
