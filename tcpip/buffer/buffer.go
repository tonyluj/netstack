package buffer

import (
	"github.com/tonyluj/netstack/tcpip"
)

type Buffer struct {
	Dev    tcpip.Device
}

func AllocBuffer() (b *Buffer) {
	b = new(Buffer)
	return
}
