package ip

import "net"

const (
	IPv4Version uint8 = 4
)

// Flags in IPv4 header
const (
	IPv4FlagReserved uint8 = 0
	IPv4FlagDontFragment uint8 = 1 << 1
	IPv4FlagMoreFragment uint8 = 1 << 2
)

// IPv4 Header
//    0                   1                   2                   3
//    0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1
//   +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
//   |Version|  IHL  |Type of Service|          Total Length         |
//   +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
//   |         Identification        |Flags|      Fragment Offset    |
//   +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
//   |  Time to Live |    Protocol   |         Header Checksum       |
//   +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
//   |                       Source Address                          |
//   +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
//   |                    Destination Address                        |
//   +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
//   |                    Options                    |    Padding    |
//   +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
type IPv4Header struct {
	Version        uint8      // 4 bits
	IHL            uint8      // 4 bits
	TypeOfService  uint8      // 8 bits
	TotalLength    uint16     // 16 bits
	ID             uint16     // 16 bits
	Flags          uint8      // 3 bits
	FragmentOffset uint16     // 13 bits
	TimeToLive     uint8      // 8 bits
	Protocol       uint8      // 8 bits
	HeaderChecksum uint16     // 16 bits
	SrcAddr        net.IPAddr   // 32 bits
	DstAddr        net.IPAddr   // 32 bits
	Option         IPv4Option // variable
	Padding        []byte     // variable
}

type IPv4Option []byte
