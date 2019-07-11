package tcp

// TCP header
//     0                   1                   2                   3
//    0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1
//   +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
//   |          Source Port          |       Destination Port        |
//   +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
//   |                        Sequence Number                        |
//   +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
//   |                    Acknowledgment Number                      |
//   +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
//   |  Data |           |U|A|P|R|S|F|                               |
//   | Offset| Reserved  |R|C|S|S|Y|I|            Window             |
//   |       |           |G|K|H|T|N|N|                               |
//   +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
//   |           Checksum            |         Urgent Pointer        |
//   +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
//   |                    Options                    |    Padding    |
//   +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
//   |                             data                              |
//   +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+

type TCPHeader struct {
	// Source port
	SrcPort uint16

	// Destination port
	DstPort uint16

	// Sequence Number
	SeqNum uint32

	// Acknowledgment Number
	AckNum uint32

	// Data offset
	DataOffset uint8 // 4 bits

	// Reserved
	Reserved uint8 // 6 bits

	URG byte // Urgent Pointer field significant
	ACK byte // Acknowledgment field significant
	PSH byte // Push Function
	RST byte // Reset the connection
	SYN byte // Synchronize sequence numbers
	FIN byte // No more data from sender

	// Window
	Window uint16

	// Checksum
	Checksum uint16

	// Urgent Pointer
	UrgentPointer uint16

	// Options
	Options TCPOption // variable

	// Padding
	Padding []byte // variable
}

type TCPOption struct {
}
