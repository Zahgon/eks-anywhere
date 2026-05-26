package networkutils

import (
	"net"
	"time"
)

type NetClient interface {
	DialTimeout(network, address string, timeout time.Duration) (net.Conn, error)
}

type DefaultNetClient struct{}

func (n *DefaultNetClient) DialTimeout(network, address string, timeout time.Duration) (net.Conn, error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), nil
}
