package tcpip

import "net"

type LinkAddr string

type NetworkProtocolNumber uint32

type Route struct {
	RemoteAddr     net.Addr
	RemoteLinkAddr LinkAddr

	LocalAddr     net.Addr
	LocalLinkAddr LinkAddr

	NextHop net.Addr

	NetProto NetworkProtocolNumber
}
