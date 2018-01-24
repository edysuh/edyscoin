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

type BroadcastRequest struct {

}

type BroadcastResponse struct {

}

func (rpcs *RpcService) Broadcast(req BroadcastRequest, res *BroadcastResponse) error {
	return nil
}
