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

func (rs *RpcService) Handshake(req HandshakeRequest, res *HandshakeResponse) error {
	rs.node.Peers[req.NodeId] = req.Address
	*res = HandshakeResponse{rs.node.Id, rs.node.Address}
	return nil
}
