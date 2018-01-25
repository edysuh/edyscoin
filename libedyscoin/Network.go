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
	// Marks		map[bId]bool
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

func (n *Node) RpcCall(raddr string, method string) (*Message, error) {
	client, err := n.ConnectToRemote(raddr)
	if err != nil {
		return nil, err
	}
	req := Message{SenderId: n.Id, SenderAddr: n.Address}
	var res Message

	call := client.Go("RpcService." + method, req, &res, nil)
	callres := <-call.Done
	if callres.Error != nil {
		log.Fatal(callres.Error)
	}

	return &res, nil
}

func (n *Node) DoHandshake(raddr string) (*Message, error) {
	res, err := n.RpcCall(raddr, "Handshake")
	n.Peers[res.SenderId] = res.SenderAddr
	return res, err
}

// TODO how do do lookups for nodes outside peer list?
func (n *Node) DoHandshakeWithId(Id Id) (*Message, error) {
	if raddr, ok := n.Peers[Id]; ok {
		return n.DoHandshake(raddr)
	}
	return nil, fmt.Errorf("node with this address is not a peer")
}

// TODO should return some sort of success i.e. the array of nodes broadcasted to
func (n *Node) DoBroadcast(method string, args struct{}, seen map[Id]bool) {
	seen[n.Id] = true
	breq := BroadcastRequest{(bId)(NewId(n.Address)), seen, method, args}

	var bres BroadcastResponse
	fmt.Println(n.Peers)

	for rId, raddr := range n.Peers {
		fmt.Printf("rid and raddr: %v %s\n", rId, raddr)
		go func(rId Id, raddr string) {
			if !seen[rId] {
				fmt.Println("here")
				client, err := n.ConnectToRemote(raddr)
				if err != nil {
					fmt.Println("couldnt connect to ", raddr, err)
				}

				broadcastCall := client.Go("RpcService.Broadcast", breq, &bres, nil)
				resCall := <- broadcastCall.Done
				if resCall.Error != nil {
					log.Fatal(resCall)
				}
				fmt.Println(bres)
			}
		}(rId, raddr)
	}
}

func (n *Node) DoBroadcastNewTransaction(txn *Transaction) (*BroadcastResponse, error) {

	return nil, nil
}

func (n *Node) DoBroadcastNewBlockChain(bc *BlockChain) (*BroadcastResponse, error) {
	return nil, nil
}
