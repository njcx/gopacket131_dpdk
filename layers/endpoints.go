// Copyright 2012 Google, Inc. All rights reserved.
//
// Use of this source code is governed by a BSD-style license
// that can be found in the LICENSE file in the root of the source
// tree.

package layers

import (
	"encoding/binary"
	"net"
	"strconv"

	"github.com/njcx/gopacket131_dpdk"
)

var (
	// We use two different endpoint types for IPv4 vs IPv6 addresses, so that
	// ordering with endpointA.LessThan(endpointB) sanely groups all IPv4
	// addresses and all IPv6 addresses, such that IPv6 > IPv4 for all addresses.
	EndpointIPv4 = gopacket131_dpdk.RegisterEndpointType(1, gopacket131_dpdk.EndpointTypeMetadata{Name: "IPv4", Formatter: func(b []byte) string {
		return net.IP(b).String()
	}})
	EndpointIPv6 = gopacket131_dpdk.RegisterEndpointType(2, gopacket131_dpdk.EndpointTypeMetadata{Name: "IPv6", Formatter: func(b []byte) string {
		return net.IP(b).String()
	}})

	EndpointMAC = gopacket131_dpdk.RegisterEndpointType(3, gopacket131_dpdk.EndpointTypeMetadata{Name: "MAC", Formatter: func(b []byte) string {
		return net.HardwareAddr(b).String()
	}})
	EndpointTCPPort = gopacket131_dpdk.RegisterEndpointType(4, gopacket131_dpdk.EndpointTypeMetadata{Name: "TCP", Formatter: func(b []byte) string {
		return strconv.Itoa(int(binary.BigEndian.Uint16(b)))
	}})
	EndpointUDPPort = gopacket131_dpdk.RegisterEndpointType(5, gopacket131_dpdk.EndpointTypeMetadata{Name: "UDP", Formatter: func(b []byte) string {
		return strconv.Itoa(int(binary.BigEndian.Uint16(b)))
	}})
	EndpointSCTPPort = gopacket131_dpdk.RegisterEndpointType(6, gopacket131_dpdk.EndpointTypeMetadata{Name: "SCTP", Formatter: func(b []byte) string {
		return strconv.Itoa(int(binary.BigEndian.Uint16(b)))
	}})
	EndpointRUDPPort = gopacket131_dpdk.RegisterEndpointType(7, gopacket131_dpdk.EndpointTypeMetadata{Name: "RUDP", Formatter: func(b []byte) string {
		return strconv.Itoa(int(b[0]))
	}})
	EndpointUDPLitePort = gopacket131_dpdk.RegisterEndpointType(8, gopacket131_dpdk.EndpointTypeMetadata{Name: "UDPLite", Formatter: func(b []byte) string {
		return strconv.Itoa(int(binary.BigEndian.Uint16(b)))
	}})
	EndpointPPP = gopacket131_dpdk.RegisterEndpointType(9, gopacket131_dpdk.EndpointTypeMetadata{Name: "PPP", Formatter: func([]byte) string {
		return "point"
	}})
)

// NewIPEndpoint creates a new IP (v4 or v6) endpoint from a net.IP address.
// It returns gopacket131_dpdk.InvalidEndpoint if the IP address is invalid.
func NewIPEndpoint(a net.IP) gopacket131_dpdk.Endpoint {
	ipv4 := a.To4()
	if ipv4 != nil {
		return gopacket131_dpdk.NewEndpoint(EndpointIPv4, []byte(ipv4))
	}

	ipv6 := a.To16()
	if ipv6 != nil {
		return gopacket131_dpdk.NewEndpoint(EndpointIPv6, []byte(ipv6))
	}

	return gopacket131_dpdk.InvalidEndpoint
}

// NewMACEndpoint returns a new MAC address endpoint.
func NewMACEndpoint(a net.HardwareAddr) gopacket131_dpdk.Endpoint {
	return gopacket131_dpdk.NewEndpoint(EndpointMAC, []byte(a))
}
func newPortEndpoint(t gopacket131_dpdk.EndpointType, p uint16) gopacket131_dpdk.Endpoint {
	return gopacket131_dpdk.NewEndpoint(t, []byte{byte(p >> 8), byte(p)})
}

// NewTCPPortEndpoint returns an endpoint based on a TCP port.
func NewTCPPortEndpoint(p TCPPort) gopacket131_dpdk.Endpoint {
	return newPortEndpoint(EndpointTCPPort, uint16(p))
}

// NewUDPPortEndpoint returns an endpoint based on a UDP port.
func NewUDPPortEndpoint(p UDPPort) gopacket131_dpdk.Endpoint {
	return newPortEndpoint(EndpointUDPPort, uint16(p))
}

// NewSCTPPortEndpoint returns an endpoint based on a SCTP port.
func NewSCTPPortEndpoint(p SCTPPort) gopacket131_dpdk.Endpoint {
	return newPortEndpoint(EndpointSCTPPort, uint16(p))
}

// NewRUDPPortEndpoint returns an endpoint based on a RUDP port.
func NewRUDPPortEndpoint(p RUDPPort) gopacket131_dpdk.Endpoint {
	return gopacket131_dpdk.NewEndpoint(EndpointRUDPPort, []byte{byte(p)})
}

// NewUDPLitePortEndpoint returns an endpoint based on a UDPLite port.
func NewUDPLitePortEndpoint(p UDPLitePort) gopacket131_dpdk.Endpoint {
	return newPortEndpoint(EndpointUDPLitePort, uint16(p))
}
