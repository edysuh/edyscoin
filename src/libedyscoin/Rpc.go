package libedyscoin

import (
	"fmt"
	// "log"
)

type RpcService struct {
	node *Node
}

type Message struct {
	MsgId      Id
	SenderId   Id    
	SenderAddr string
	Method     string
	Params	   Params
}

type Params struct {
	BlockChain  *BlockChain
	Transaction *Transaction
	Seenlist    map[Id]bool
	Success     []Id
}

func NewMessage(node *Node, method string, params ...Params) *Message {
	msg := &Message{
		MsgId:      NewId(node.Address),
		SenderId:   node.Id,
		SenderAddr: node.Address,
		Method:     method,
	}
	if len(params) == 1 {
		msg.Params = params[0]
	}
	return msg
}

func (rpcs *RpcService) Handshake(req Message, res *Message) error {
	if !rpcs.node.Id.Equals(req.SenderId) {
		rpcs.node.Peers[req.SenderId] = req.SenderAddr
	}

	*res = Message{
		MsgId:      req.MsgId,
		SenderId:   rpcs.node.Id,
		SenderAddr: rpcs.node.Address,
	}
	return nil
}

func (rpcs *RpcService) SyncBlockChain(req Message, res *Message) error {
	*res = Message{
		MsgId:      req.MsgId,
		SenderId:   rpcs.node.Id,
		SenderAddr: rpcs.node.Address,
		Params:		Params{BlockChain: rpcs.node.BlockChain},
	}
	return nil
}

func (rpcs *RpcService) Broadcast(req Message, res *Message) error {
	// TODO _?
	_, err := rpcs.node.DoBroadcast(&req)
	if err != nil {
		fmt.Println(err)
	}
	*res = Message{
		MsgId:      req.MsgId,
		SenderId:   rpcs.node.Id,
		SenderAddr: rpcs.node.Address,
	}
	fmt.Printf("%+v\n", req)
	return nil
}

func (rpcs *RpcService) BroadcastNewTransaction(req Message, res *Message) error {
	txn := req.Params.Transaction
	rpcs.node.BlockChain.NewTransaction(txn)
	// rpcs.node.BlockChain.ListTransactions()

	rpcs.Broadcast(req, res)
	return nil
}

func (rpcs *RpcService) BroadcastNewBlockChain(req Message, res *Message) error {
	localbc := rpcs.node.BlockChain
	remotebc := req.Params.BlockChain
	fmt.Printf("%#v\n", localbc)
	fmt.Printf("%#v\n", remotebc)

	if err := localbc.Consensus(remotebc); err != nil {
		panic(err)
	}
	*rpcs.node.BlockChain = *req.Params.BlockChain

	rpcs.Broadcast(req, res)
	return nil
}
