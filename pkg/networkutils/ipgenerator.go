package networkutils

import (
	"math/rand"
	"net"
)

type IPGenerator struct {
	netClient NetClient
	rand      *rand.Rand
}

func NewIPGenerator(netClient NetClient) IPGenerator {
	_ = "STUB: not implemented"
	return *new(IPGenerator)
}

func (ipgen IPGenerator) GenerateUniqueIP(cidrBlock string, usedIPs map[string]bool) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Count total usable IPs in the CIDR range (skip network address)

// Pick a random starting offset to reduce collisions between parallel instances

// Walk through all IPs starting from the random offset, wrapping around

// offset 0 = first host IP (network address + 1)

// cidrHostCount returns the number of usable host IPs in a CIDR range (excluding network address).
func cidrHostCount(cidr *net.IPNet) int { _ = "STUB: not implemented"; return 0 }

// Total addresses = 2^(bits-ones), subtract 1 for network address

// addToIP returns a new IP that is base + offset.
func addToIP(base net.IP, offset int) net.IP { _ = "STUB: not implemented"; return *new(net.IP) }

// copyIP creates a copy of the IP address.
func copyIP(ip net.IP) net.IP { _ = "STUB: not implemented"; return *new(net.IP) }
