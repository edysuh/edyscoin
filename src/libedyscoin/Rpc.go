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

func (rpcs *RpcService) Handshake(req Message, res *Message) error {
	if !rpcs.node.Id.Equals(req.SenderId) {
		rpcs.node.Peers[req.SenderId] = req.SenderAddr
	}
	*res = Message{SenderId: rpcs.node.Id, SenderAddr: rpcs.node.Address}
	return nil
}

func (rpcs *RpcService) Broadcast(req Message, res *Message) error {
	rpcs.node.DoBroadcast(&req)
	*res = Message{MsgId: req.MsgId, SenderId: rpcs.node.Id, SenderAddr: rpcs.node.Address}
	fmt.Printf("%+v\n\n", req)
	return nil
}
