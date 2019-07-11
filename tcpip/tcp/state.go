package tcp

const (
	TCPEstablished = iota
	TCPSynSent
	TCPSynRcv
	TCPFinWait1
	TCPFinWait2
	TCPTimeWait
	TCPClose
	TCPCloseWait
	TCPLastAck
	TCPListen
	TCPClosing
	TCPNewRcv
)
