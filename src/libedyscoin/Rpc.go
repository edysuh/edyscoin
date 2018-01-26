package libedyscoin

import (
	"encoding/json"
	"fmt"
	"log"
)

type RpcService struct {
	node *Node
}

type Message struct {
	MsgId      Id          `json:"msgid"`
	SenderId   Id          `json:"senderid"`
	SenderAddr string      `json:"senderaddr"`
	Method     string      `json:"method,omitempty"`
	Params	   Params	   `json:"params,omitempty"`
}

type Params struct {
	Payload    []byte      `json:"payload,omitempty"`
	Seenlist   map[Id]bool `json:"seenlist,omitempty"`
	Success    []Id        `json:"success,omitempty"`
}

func NewMessage(node *Node, params ...Params) *Message {
	msg := &Message{
		MsgId:      NewId(node.Address),
		SenderId:   node.Id,
		SenderAddr: node.Address,
		Method:     "Broadcast",
	}
	if 0 < len(params) && len(params) < 2 {
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
	var txn *Transaction
	if err := json.Unmarshal(req.Params.Payload, txn); err != nil {
		log.Fatal(err)
	}
	rpcs.node.BlockChain.NewTransaction(*txn)

	_, err := rpcs.node.DoBroadcastNewTransaction(txn)
	if err != nil {
		fmt.Println(err)
	}

	*res = Message{
		MsgId:      req.MsgId,
		SenderId:   rpcs.node.Id,
		SenderAddr: rpcs.node.Address,
	}
	return nil
}

func (rpcs *RpcService) BroadcastNewBlockChain(req Message, res *Message) error {
	localbc := rpcs.node.BlockChain
	var remotebc *BlockChain
	if err := json.Unmarshal(req.Params.Payload, remotebc); err != nil {
		log.Fatal(err)
	}
	if err := localbc.Consensus(remotebc); err != nil {
		log.Fatal(err)
	}

	_, err := rpcs.node.DoBroadcastNewBlockChain(remotebc)
	if err != nil {
		fmt.Println(err)
	}

	*res = Message{
		MsgId:      req.MsgId,
		SenderId:   rpcs.node.Id,
		SenderAddr: rpcs.node.Address,
	}
	return nil
}
