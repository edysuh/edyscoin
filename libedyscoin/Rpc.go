package libedyscoin

import (

)

type RpcService struct {
	node *Node
}

type HandshakeRequest struct {
	nodeId  Id
	address string
}

type HandshakeResponse struct {
	nodeId  Id
	address string
}

func (rs *RpcService) Handshake(req HandshakeRequest, res *HandshakeResponse) {
	// TODO net.Conn? or contact info?
	rs.node.Peers[req.nodeId] = req.address
	res = &HandshakeResponse{rs.node.Id, rs.node.Address}
}
