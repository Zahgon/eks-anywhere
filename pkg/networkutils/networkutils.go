package networkutils

import (
	"net"
)

func IsPortValid(port string) bool { _ = "STUB: not implemented"; return false }

func ValidateIP(ip string) error { _ = "STUB: not implemented"; return nil }

// IsIPInUse performs a best effort check to see if an IP address is in use. It is not completely
// reliable as testing if an IP is in use is inherently difficult, particularly with non-trivial
// network topologies.
func IsIPInUse(client NetClient, ip string) bool {
	_ = "STUB: not implemented"
	// Dial and immediately close the connection if it was established as its superfluous for
	// our check. We use port 80 as its common and is more likely to get through firewalls
	// than other ports.
	return false
}

// If we establish a connection or we receive a response assume that address is in use.
// The latter case covers situations like an IP in use but the port requested is not open.

func IsPortInUse(client NetClient, host string, port string) bool {
	_ = "STUB: not implemented"
	return false
}

func GetLocalIP() (net.IP, error) { _ = "STUB: not implemented"; return *new(net.IP), nil }
