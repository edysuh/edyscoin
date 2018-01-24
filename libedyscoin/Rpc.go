package libedyscoin

import (

)

type RpcService struct {
	node *Node
}

type HandshakeRequest struct {
	NodeId  Id
	Address string
}

type HandshakeResponse struct {
	NodeId  Id
	Address string
}

func (rpcs *RpcService) Handshake(req HandshakeRequest, res *HandshakeResponse) error {
	rpcs.node.Peers[req.NodeId] = req.Address
	*res = HandshakeResponse{rpcs.node.Id, rpcs.node.Address}
	return nil
}

// TODO change this to a struct that includes a TTL
type bId Id		// BroadcastId

type BroadcastRequest struct {
	BId  bId
	Args interface{}
}

type BroadcastResponse struct {
	BId     bId
	ResNode Id
}

type TransactionBroadcast struct {
	Seen map[Id]bool
}

func (rpcs *RpcService) Broadcast(req BroadcastRequest, res *BroadcastResponse) error {
	breq := req.(BroadcastRequest)
	if id, ok := req.Args.Seen; ok {
		for _, node := range rpcs.node.Peers {

		}

	}
	return nil
}
