package constant

import (
	"net/netip"

	"github.com/kitty314/1.17.0/transport/socks5"
)

const (
	BpfFSPath = "/sys/fs/bpf/clash"

	TcpAutoRedirPort  = 't'<<8 | 'r'<<0
	clashTrafficMark = 'c'<<24 | 'l'<<16 | 't'<<8 | 'm'<<0
)

type EBpf interface {
	Start() error
	Close()
	Lookup(srcAddrPort netip.AddrPort) (socks5.Addr, error)
}
