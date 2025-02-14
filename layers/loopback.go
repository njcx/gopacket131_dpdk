// Copyright 2012 Google, Inc. All rights reserved.
//
// Use of this source code is governed by a BSD-style license
// that can be found in the LICENSE file in the root of the source
// tree.

package layers

import (
	"encoding/binary"
	"errors"
	"fmt"

	"github.com/njcx/gopacket131_dpdk"
)

// Loopback contains the header for loopback encapsulation.  This header is
// used by both BSD and OpenBSD style loopback decoding (pcap's DLT_NULL
// and DLT_LOOP, respectively).
type Loopback struct {
	BaseLayer
	Family ProtocolFamily
}

// LayerType returns LayerTypeLoopback.
func (l *Loopback) LayerType() gopacket131_dpdk.LayerType { return LayerTypeLoopback }

// DecodeFromBytes decodes the given bytes into this layer.
func (l *Loopback) DecodeFromBytes(data []byte, df gopacket131_dpdk.DecodeFeedback) error {
	if len(data) < 4 {
		return errors.New("Loopback packet too small")
	}

	// The protocol could be either big-endian or little-endian, we're
	// not sure.  But we're PRETTY sure that the value is less than
	// 256, so we can check the first two bytes.
	var prot uint32
	if data[0] == 0 && data[1] == 0 {
		prot = binary.BigEndian.Uint32(data[:4])
	} else {
		prot = binary.LittleEndian.Uint32(data[:4])
	}
	if prot > 0xFF {
		return fmt.Errorf("Invalid loopback protocol %q", data[:4])
	}

	l.Family = ProtocolFamily(prot)
	l.BaseLayer = BaseLayer{data[:4], data[4:]}
	return nil
}

// CanDecode returns the set of layer types that this DecodingLayer can decode.
func (l *Loopback) CanDecode() gopacket131_dpdk.LayerClass {
	return LayerTypeLoopback
}

// NextLayerType returns the layer type contained by this DecodingLayer.
func (l *Loopback) NextLayerType() gopacket131_dpdk.LayerType {
	return l.Family.LayerType()
}

// SerializeTo writes the serialized form of this layer into the
// SerializationBuffer, implementing gopacket131_dpdk.SerializableLayer.
func (l *Loopback) SerializeTo(b gopacket131_dpdk.SerializeBuffer, opts gopacket131_dpdk.SerializeOptions) error {
	bytes, err := b.PrependBytes(4)
	if err != nil {
		return err
	}
	binary.LittleEndian.PutUint32(bytes, uint32(l.Family))
	return nil
}

func decodeLoopback(data []byte, p gopacket131_dpdk.PacketBuilder) error {
	l := Loopback{}
	if err := l.DecodeFromBytes(data, gopacket131_dpdk.NilDecodeFeedback); err != nil {
		return err
	}
	p.AddLayer(&l)
	return p.NextDecoder(l.Family)
}
