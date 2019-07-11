package netstack

import "net"

type Domain uint8

type SocketType uint8

const (
	SockStream uint8 = 0
	SockDgram  uint8 = 1
)

type Protocol uint8

type SocketAddr struct {
	net.IPAddr
	Port int
}

type Socket struct {
	Domain     Domain
	SocketType SocketType
	Protocol   Protocol
	SrcAddr    SocketAddr
	DstAddr    SocketAddr
	Backlog    int
}

func NewSocket(domain Domain, socketType SocketType, protocol Protocol) (socket *Socket) {
	socket = new(Socket)
	socket.Domain = domain
	socket.SocketType = socketType
	socket.Protocol = protocol

	return
}

func (s *Socket) Bind(addr SocketAddr) (err error) {
	s.SrcAddr = addr
	return
}

func (s *Socket) Listen(backlog int) (err error) {
	s.Backlog = backlog

	// get packet from ring buffer
	return
}

func (s *Socket) Accept() (ns *Socket, err error) {
	// create connection
	return
}

func (s *Socket) Connect(addr SocketAddr) {
	s.DstAddr = addr
}
