package libedyscoin

import (
	"fmt"
)

type RpcService struct {
	node *Node
}

type Message struct {
	MsgId      Id
	SenderId   Id
	SenderAddr string
	Method     string
	Args       string
	Seenlist   map[Id]bool
	Success    bool
}

type HandshakeRequest struct {
	NodeId  Id
	Address string
}

type HandshakeResponse struct {
	NodeId  Id
	Address string
}

func (rpcs *RpcService) Handshake(req Message, res *Message) error {
	rpcs.node.Peers[req.SenderId] = req.SenderAddr
	*res = Message{SenderId: rpcs.node.Id, SenderAddr: rpcs.node.Address}
	return nil
}

type bId Id		// BroadcastId

type BroadcastRequest struct {
	BId    bId
	Seen   map[Id]bool
	Method string
	Args   struct{}
}

type BroadcastResponse struct {
	BId     bId
	ResNode Id
}

type TransactionBroadcast struct {
	txn *Transaction
}

type BlockChainBroadcast struct {
	bc *BlockChain
}

// TODO NEED BID!!
func (rpcs *RpcService) Broadcast(req BroadcastRequest, res *BroadcastResponse) error {
	// args := req.Args
	fmt.Println(req.Method)
	rpcs.node.DoBroadcast(req.Method, req.Args, req.Seen)
	*res = BroadcastResponse{req.BId, rpcs.node.Id}
	return nil
}
