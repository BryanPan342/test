// Copyright 2017 Google, Inc. All rights reserved.
//
// Use of this source code is governed by a BSD-style license
// that can be found in the LICENSE file in the root of the source
// tree.

package layers

import (
	"github.com/google/gopacket"
	"reflect"
	"testing"
)

// testPacketOSPF3Hello is the packet:
//   14:43:11.663317 IP6 fe80::1 > ff02::5: OSPFv3, Hello, length 36
//   	0x0000:  3333 0000 0005 c200 1ffa 0001 86dd 6e00  33............n.
//   	0x0010:  0000 0024 5901 fe80 0000 0000 0000 0000  ...$Y...........
//   	0x0020:  0000 0000 0001 ff02 0000 0000 0000 0000  ................
//   	0x0030:  0000 0000 0005 0301 0024 0101 0101 0000  .........$......
//   	0x0040:  0001 fb86 0000 0000 0005 0100 0013 000a  ................
//   	0x0050:  0028 0000 0000 0000 0000                 .(........
var testPacketOSPF3Hello = []byte{
	0x33, 0x33, 0x00, 0x00, 0x00, 0x05, 0xc2, 0x00, 0x1f, 0xfa, 0x00, 0x01, 0x86, 0xdd, 0x6e, 0x00,
	0x00, 0x00, 0x00, 0x24, 0x59, 0x01, 0xfe, 0x80, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0xff, 0x02, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x05, 0x03, 0x01, 0x00, 0x24, 0x01, 0x01, 0x01, 0x01, 0x00, 0x00,
	0x00, 0x01, 0xfb, 0x86, 0x00, 0x00, 0x00, 0x00, 0x00, 0x05, 0x01, 0x00, 0x00, 0x13, 0x00, 0x0a,
	0x00, 0x28, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
}

func TestPacketOSPF3Hello(t *testing.T) {
	p := gopacket.NewPacket(testPacketOSPF3Hello, LinkTypeEthernet, gopacket.Default)
	if p.ErrorLayer() != nil {
		t.Error("Failed to decode packet:", p.ErrorLayer().Error())
	}
	checkLayers(p, []gopacket.LayerType{LayerTypeEthernet, LayerTypeIPv6, LayerTypeOSPF}, t)

	ospf := p.Layer(LayerTypeOSPF).(*OSPFv3)
	if ospf.Version != 3 {
		t.Fatal("Invalid OSPF version")
	}
	if got, ok := p.Layer(LayerTypeOSPF).(*OSPFv3); ok {
		want := &OSPFv3{
			OSPF: OSPF{
				Version:      3,
				Type:         OSPFHello,
				PacketLength: 36,
				RouterID:     0x1010101,
				AreaID:       1,
				Checksum:     0xfb86,
				Content: HelloPkg{
					InterfaceID:              5,
					RtrPriority:              1,
					Options:                  0x000013,
					HelloInterval:            10,
					RouterDeadInterval:       40,
					DesignatedRouterID:       0,
					BackupDesignatedRouterID: 0,
				},
			},
			Instance: 0,
			Reserved: 0,
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("OSPF packet processing failed:\ngot  :\n%#v\n\nwant :\n%#v\n\n", got, want)
		}
	} else {
		t.Error("No OSPF layer type found in packet")
	}
}
func BenchmarkDecodePacketPacket0(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gopacket.NewPacket(testPacketOSPF3Hello, LinkTypeEthernet, gopacket.NoCopy)
	}
}

// testPacketOSPF3DBDesc is the packet:
//   14:43:51.657571 IP6 fe80::2 > fe80::1: OSPFv3, Database Description, length 28
//   	0x0000:  c200 1ffa 0001 c201 1ffa 0001 86dd 6e00  ..............n.
//   	0x0010:  0000 001c 5901 fe80 0000 0000 0000 0000  ....Y...........
//   	0x0020:  0000 0000 0002 fe80 0000 0000 0000 0000  ................
//   	0x0030:  0000 0000 0001 0302 001c 0202 0202 0000  ................
//   	0x0040:  0001 d826 0000 0000 0013 05dc 0007 0000  ...&............
//   	0x0050:  1d46                                     .F
var testPacketOSPF3DBDesc = []byte{
	0xc2, 0x00, 0x1f, 0xfa, 0x00, 0x01, 0xc2, 0x01, 0x1f, 0xfa, 0x00, 0x01, 0x86, 0xdd, 0x6e, 0x00,
	0x00, 0x00, 0x00, 0x1c, 0x59, 0x01, 0xfe, 0x80, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xfe, 0x80, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x03, 0x02, 0x00, 0x1c, 0x02, 0x02, 0x02, 0x02, 0x00, 0x00,
	0x00, 0x01, 0xd8, 0x26, 0x00, 0x00, 0x00, 0x00, 0x00, 0x13, 0x05, 0xdc, 0x00, 0x07, 0x00, 0x00,
	0x1d, 0x46,
}

func TestPacketOSPF3DBDesc(t *testing.T) {
	p := gopacket.NewPacket(testPacketOSPF3DBDesc, LinkTypeEthernet, gopacket.Default)
	if p.ErrorLayer() != nil {
		t.Error("Failed to decode packet:", p.ErrorLayer().Error())
	}
	checkLayers(p, []gopacket.LayerType{LayerTypeEthernet, LayerTypeIPv6, LayerTypeOSPF}, t)
	if got, ok := p.Layer(LayerTypeOSPF).(*OSPFv3); ok {
		want := &OSPFv3{
			OSPF: OSPF{
				Version:      3,
				Type:         OSPFDatabaseDescription,
				PacketLength: 28,
				RouterID:     0x2020202,
				AreaID:       1,
				Checksum:     0xd826,
				Content: DbDescPkg{
					Options:      0x000013,
					InterfaceMTU: 1500,
					Flags:        0x7,
					DDSeqNumber:  7494,
				},
			},
			Instance: 0,
			Reserved: 0,
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("OSPF packet processing failed:\ngot  :\n%#v\n\nwant :\n%#v\n\n", got, want)
		}
	} else {
		t.Error("No OSPF layer type found in packet")
	}
}
func BenchmarkDecodePacketPacket1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gopacket.NewPacket(testPacketOSPF3DBDesc, LinkTypeEthernet, gopacket.NoCopy)
	}
}

// testPacketOSPF3LSRequest is the packet:
//   14:43:51.673584 IP6 fe80::2 > fe80::1: OSPFv3, LS-Request, length 100
//   	0x0000:  c200 1ffa 0001 c201 1ffa 0001 86dd 6e00  ..............n.
//   	0x0010:  0000 0064 5901 fe80 0000 0000 0000 0000  ...dY...........
//   	0x0020:  0000 0000 0002 fe80 0000 0000 0000 0000  ................
//   	0x0030:  0000 0000 0001 0303 0064 0202 0202 0000  .........d......
//   	0x0040:  0001 2c9a 0000 0000 2001 0000 0000 0101  ..,.............
//   	0x0050:  0101 0000 2003 0000 0003 0101 0101 0000  ................
//   	0x0060:  2003 0000 0002 0101 0101 0000 2003 0000  ................
//   	0x0070:  0001 0101 0101 0000 2003 0000 0000 0101  ................
//   	0x0080:  0101 0000 0008 0000 0005 0101 0101 0000  ................
//   	0x0090:  2009 0000 0000 0101 0101                 ..........
var testPacketOSPF3LSRequest = []byte{
	0xc2, 0x00, 0x1f, 0xfa, 0x00, 0x01, 0xc2, 0x01, 0x1f, 0xfa, 0x00, 0x01, 0x86, 0xdd, 0x6e, 0x00,
	0x00, 0x00, 0x00, 0x64, 0x59, 0x01, 0xfe, 0x80, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xfe, 0x80, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x03, 0x03, 0x00, 0x64, 0x02, 0x02, 0x02, 0x02, 0x00, 0x00,
	0x00, 0x01, 0x2c, 0x9a, 0x00, 0x00, 0x00, 0x00, 0x20, 0x01, 0x00, 0x00, 0x00, 0x00, 0x01, 0x01,
	0x01, 0x01, 0x00, 0x00, 0x20, 0x03, 0x00, 0x00, 0x00, 0x03, 0x01, 0x01, 0x01, 0x01, 0x00, 0x00,
	0x20, 0x03, 0x00, 0x00, 0x00, 0x02, 0x01, 0x01, 0x01, 0x01, 0x00, 0x00, 0x20, 0x03, 0x00, 0x00,
	0x00, 0x01, 0x01, 0x01, 0x01, 0x01, 0x00, 0x00, 0x20, 0x03, 0x00, 0x00, 0x00, 0x00, 0x01, 0x01,
	0x01, 0x01, 0x00, 0x00, 0x00, 0x08, 0x00, 0x00, 0x00, 0x05, 0x01, 0x01, 0x01, 0x01, 0x00, 0x00,
	0x20, 0x09, 0x00, 0x00, 0x00, 0x00, 0x01, 0x01, 0x01, 0x01,
}

func TestPacketOSPF3LSRequest(t *testing.T) {
	p := gopacket.NewPacket(testPacketOSPF3LSRequest, LinkTypeEthernet, gopacket.Default)
	if p.ErrorLayer() != nil {
		t.Error("Failed to decode packet:", p.ErrorLayer().Error())
	}
	checkLayers(p, []gopacket.LayerType{LayerTypeEthernet, LayerTypeIPv6, LayerTypeOSPF}, t)
	if got, ok := p.Layer(LayerTypeOSPF).(*OSPFv3); ok {
		want := &OSPFv3{
			OSPF: OSPF{
				Version:      3,
				Type:         OSPFLinkStateRequest,
				PacketLength: 100,
				RouterID:     0x2020202,
				AreaID:       1,
				Checksum:     0x2c9a,
				Content: []LSReq{
					LSReq{
						LSType:    0x2001,
						LSID:      0x00000000,
						AdvRouter: 0x01010101,
					},
					LSReq{
						LSType:    0x2003,
						LSID:      0x00000003,
						AdvRouter: 0x01010101,
					},
					LSReq{
						LSType:    0x2003,
						LSID:      0x00000002,
						AdvRouter: 0x01010101,
					},
					LSReq{
						LSType:    0x2003,
						LSID:      0x00000001,
						AdvRouter: 0x01010101,
					},
					LSReq{
						LSType:    0x2003,
						LSID:      0x00000000,
						AdvRouter: 0x01010101,
					},
					LSReq{
						LSType:    0x0008,
						LSID:      0x00000005,
						AdvRouter: 0x01010101,
					},
					LSReq{
						LSType:    0x2009,
						LSID:      0x00000000,
						AdvRouter: 0x01010101,
					},
				},
			},
			Instance: 0,
			Reserved: 0,
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("OSPF packet processing failed:\ngot  :\n%#v\n\nwant :\n%#v\n\n", got, want)
		}
	} else {
		t.Error("No OSPF layer type found in packet")
	}
}
func BenchmarkDecodePacketPacket2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gopacket.NewPacket(testPacketOSPF3LSRequest, LinkTypeEthernet, gopacket.NoCopy)
	}
}

// testPacketOSPF3LSUpdate is the packet:
//   14:43:51.681554 IP6 fe80::1 > fe80::2: OSPFv3, LS-Update, length 288
//   	0x0000:  c201 1ffa 0001 c200 1ffa 0001 86dd 6e00  ..............n.
//   	0x0010:  0000 0120 5901 fe80 0000 0000 0000 0000  ....Y...........
//   	0x0020:  0000 0000 0001 fe80 0000 0000 0000 0000  ................
//   	0x0030:  0000 0000 0002 0304 0120 0101 0101 0000  ................
//   	0x0040:  0001 e556 0000 0000 0007 0028 2001 0000  ...V.......(....
//   	0x0050:  0000 0101 0101 8000 0002 d13a 0018 0100  ...........:....
//   	0x0060:  0033 0029 2003 0000 0003 0101 0101 8000  .3.)............
//   	0x0070:  0001 6259 0024 0000 004a 4000 0000 2001  ..bY.$...J@.....
//   	0x0080:  0db8 0000 0003 0029 2003 0000 0002 0101  .......)........
//   	0x0090:  0101 8000 0001 baf6 0024 0000 0054 4000  .........$...T@.
//   	0x00a0:  0000 2001 0db8 0000 0004 0029 2003 0000  ...........)....
//   	0x00b0:  0001 0101 0101 8000 0001 eba0 0024 0000  .............$..
//   	0x00c0:  004a 4000 0000 2001 0db8 0000 0034 0029  .J@..........4.)
//   	0x00d0:  2003 0000 0000 0101 0101 8000 0001 0ebd  ................
//   	0x00e0:  0024 0000 0040 4000 0000 2001 0db8 0000  .$...@@.........
//   	0x00f0:  0000 0023 0008 0000 0005 0101 0101 8000  ...#............
//   	0x0100:  0002 3d08 0038 0100 0033 fe80 0000 0000  ..=..8...3......
//   	0x0110:  0000 0000 0000 0000 0001 0000 0001 4000  ..............@.
//   	0x0120:  0000 2001 0db8 0000 0012 0023 2009 0000  ...........#....
//   	0x0130:  0000 0101 0101 8000 0001 e8d2 002c 0001  .............,..
//   	0x0140:  2001 0000 0000 0101 0101 4000 000a 2001  ..........@.....
//   	0x0150:  0db8 0000 0012                           ......
var testPacketOSPF3LSUpdate = []byte{
	0xc2, 0x01, 0x1f, 0xfa, 0x00, 0x01, 0xc2, 0x00, 0x1f, 0xfa, 0x00, 0x01, 0x86, 0xdd, 0x6e, 0x00,
	0x00, 0x00, 0x01, 0x20, 0x59, 0x01, 0xfe, 0x80, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0xfe, 0x80, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0x03, 0x04, 0x01, 0x20, 0x01, 0x01, 0x01, 0x01, 0x00, 0x00,
	0x00, 0x01, 0xe5, 0x56, 0x00, 0x00, 0x00, 0x00, 0x00, 0x07, 0x00, 0x28, 0x20, 0x01, 0x00, 0x00,
	0x00, 0x00, 0x01, 0x01, 0x01, 0x01, 0x80, 0x00, 0x00, 0x02, 0xd1, 0x3a, 0x00, 0x18, 0x01, 0x00,
	0x00, 0x33, 0x00, 0x29, 0x20, 0x03, 0x00, 0x00, 0x00, 0x03, 0x01, 0x01, 0x01, 0x01, 0x80, 0x00,
	0x00, 0x01, 0x62, 0x59, 0x00, 0x24, 0x00, 0x00, 0x00, 0x4a, 0x40, 0x00, 0x00, 0x00, 0x20, 0x01,
	0x0d, 0xb8, 0x00, 0x00, 0x00, 0x03, 0x00, 0x29, 0x20, 0x03, 0x00, 0x00, 0x00, 0x02, 0x01, 0x01,
	0x01, 0x01, 0x80, 0x00, 0x00, 0x01, 0xba, 0xf6, 0x00, 0x24, 0x00, 0x00, 0x00, 0x54, 0x40, 0x00,
	0x00, 0x00, 0x20, 0x01, 0x0d, 0xb8, 0x00, 0x00, 0x00, 0x04, 0x00, 0x29, 0x20, 0x03, 0x00, 0x00,
	0x00, 0x01, 0x01, 0x01, 0x01, 0x01, 0x80, 0x00, 0x00, 0x01, 0xeb, 0xa0, 0x00, 0x24, 0x00, 0x00,
	0x00, 0x4a, 0x40, 0x00, 0x00, 0x00, 0x20, 0x01, 0x0d, 0xb8, 0x00, 0x00, 0x00, 0x34, 0x00, 0x29,
	0x20, 0x03, 0x00, 0x00, 0x00, 0x00, 0x01, 0x01, 0x01, 0x01, 0x80, 0x00, 0x00, 0x01, 0x0e, 0xbd,
	0x00, 0x24, 0x00, 0x00, 0x00, 0x40, 0x40, 0x00, 0x00, 0x00, 0x20, 0x01, 0x0d, 0xb8, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x23, 0x00, 0x08, 0x00, 0x00, 0x00, 0x05, 0x01, 0x01, 0x01, 0x01, 0x80, 0x00,
	0x00, 0x02, 0x3d, 0x08, 0x00, 0x38, 0x01, 0x00, 0x00, 0x33, 0xfe, 0x80, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x01, 0x40, 0x00,
	0x00, 0x00, 0x20, 0x01, 0x0d, 0xb8, 0x00, 0x00, 0x00, 0x12, 0x00, 0x23, 0x20, 0x09, 0x00, 0x00,
	0x00, 0x00, 0x01, 0x01, 0x01, 0x01, 0x80, 0x00, 0x00, 0x01, 0xe8, 0xd2, 0x00, 0x2c, 0x00, 0x01,
	0x20, 0x01, 0x00, 0x00, 0x00, 0x00, 0x01, 0x01, 0x01, 0x01, 0x40, 0x00, 0x00, 0x0a, 0x20, 0x01,
	0x0d, 0xb8, 0x00, 0x00, 0x00, 0x12,
}

func TestPacketOSPF3LSUpdate(t *testing.T) {
	p := gopacket.NewPacket(testPacketOSPF3LSUpdate, LinkTypeEthernet, gopacket.Default)
	if p.ErrorLayer() != nil {
		t.Error("Failed to decode packet:", p.ErrorLayer().Error())
	}
	checkLayers(p, []gopacket.LayerType{LayerTypeEthernet, LayerTypeIPv6, LayerTypeOSPF}, t)
	if got, ok := p.Layer(LayerTypeOSPF).(*OSPFv3); ok {
		want := &OSPFv3{
			OSPF: OSPF{
				Version:      3,
				Type:         OSPFLinkStateUpdate,
				PacketLength: 288,
				RouterID:     0x1010101,
				AreaID:       1,
				Checksum:     0xe556,
				Content: LSUpdate{
					NumOfLSAs: 7,
					LSAs: []LSA{
						LSA{
							LSAheader: LSAheader{
								LSAge:       40,
								LSType:      0x2001,
								LinkStateID: 0x00000000,
								AdvRouter:   0x01010101,
								LSSeqNumber: 0x80000002,
								LSChecksum:  0xd13a,
								Length:      24,
							},
							Content: RouterLSA{
								Flags:   0x1,
								Options: 0x33,
							},
						},
						LSA{
							LSAheader: LSAheader{
								LSAge:       41,
								LSType:      0x2003,
								LinkStateID: 0x00000003,
								AdvRouter:   0x01010101,
								LSSeqNumber: 0x80000001,
								LSChecksum:  0x6259,
								Length:      36,
							},
							Content: InterAreaPrefixLSA{
								Metric:        74,
								PrefixLength:  64,
								PrefixOptions: 0,
								AddressPrefix: []byte{0x20, 0x01, 0x0d, 0xb8, 0x00, 0x00, 0x00, 0x03},
							},
						},
						LSA{
							LSAheader: LSAheader{
								LSAge:       41,
								LSType:      0x2003,
								LinkStateID: 0x00000002,
								AdvRouter:   0x01010101,
								LSSeqNumber: 0x80000001,
								LSChecksum:  0xbaf6,
								Length:      36,
							},
							Content: InterAreaPrefixLSA{
								Metric:        84,
								PrefixLength:  64,
								PrefixOptions: 0,
								AddressPrefix: []byte{0x20, 0x1, 0xd, 0xb8, 0x0, 0x0, 0x0, 0x4},
							},
						},
						LSA{
							LSAheader: LSAheader{
								LSAge:       41,
								LSType:      0x2003,
								LinkStateID: 0x00000001,
								AdvRouter:   0x01010101,
								LSSeqNumber: 0x80000001,
								LSChecksum:  0xeba0,
								Length:      36,
							},
							Content: InterAreaPrefixLSA{
								Metric:        74,
								PrefixLength:  64,
								PrefixOptions: 0,
								AddressPrefix: []byte{0x20, 0x1, 0xd, 0xb8, 0x0, 0x0, 0x0, 0x34},
							},
						},
						LSA{
							LSAheader: LSAheader{
								LSAge:       41,
								LSType:      0x2003,
								LinkStateID: 0x00000000,
								AdvRouter:   0x01010101,
								LSSeqNumber: 0x80000001,
								LSChecksum:  0xebd,
								Length:      36,
							},
							Content: InterAreaPrefixLSA{
								Metric:        64,
								PrefixLength:  64,
								PrefixOptions: 0,
								AddressPrefix: []byte{0x20, 0x1, 0xd, 0xb8, 0x0, 0x0, 0x0, 0x0},
							},
						},
						LSA{
							LSAheader: LSAheader{
								LSAge:       35,
								LSType:      0x8,
								LinkStateID: 0x00000005,
								AdvRouter:   0x01010101,
								LSSeqNumber: 0x80000002,
								LSChecksum:  0x3d08,
								Length:      56,
							},
							Content: LinkLSA{
								RtrPriority:      1,
								Options:          0x33,
								LinkLocalAddress: []byte{0xfe, 0x80, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01},
								NumOfPrefixes:    1,
								Prefixes: []Prefix{
									Prefix{
										PrefixLength:  64,
										PrefixOptions: 0,
										AddressPrefix: []byte{0x20, 0x01, 0x0d, 0xb8, 0x00, 0x00, 0x00, 0x12},
									},
								},
							},
						},
						LSA{
							LSAheader: LSAheader{
								LSAge:       35,
								LSType:      0x2009,
								LinkStateID: 0x00000000,
								AdvRouter:   0x01010101,
								LSSeqNumber: 0x80000001,
								LSChecksum:  0xe8d2,
								Length:      44,
							},
							Content: IntraAreaPrefixLSA{
								NumOfPrefixes: 1,
								RefLSType:     0x2001,
								RefAdvRouter:  0x01010101,
								Prefixes: []Prefix{
									Prefix{
										PrefixLength:  64,
										PrefixOptions: 0,
										Metric:        10,
										AddressPrefix: []byte{0x20, 0x1, 0xd, 0xb8, 0x0, 0x0, 0x0, 0x12},
									},
								},
							},
						},
					},
				},
			},
			Instance: 0,
			Reserved: 0,
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("OSPF packet processing failed:\ngot  :\n%#v\n\nwant :\n%#v\n\n", got, want)
		}
	} else {
		t.Error("No OSPF layer type found in packet")
	}
}
func BenchmarkDecodePacketPacket3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gopacket.NewPacket(testPacketOSPF3LSUpdate, LinkTypeEthernet, gopacket.NoCopy)
	}
}

// testPacketOSPF3LSAck is the packet:
//   14:43:54.185384 IP6 fe80::1 > ff02::5: OSPFv3, LS-Ack, length 136
//   	0x0000:  3333 0000 0005 c200 1ffa 0001 86dd 6e00  33............n.
//   	0x0010:  0000 0088 5901 fe80 0000 0000 0000 0000  ....Y...........
//   	0x0020:  0000 0000 0001 ff02 0000 0000 0000 0000  ................
//   	0x0030:  0000 0000 0005 0305 0088 0101 0101 0000  ................
//   	0x0040:  0001 9d2c 0000 0005 2001 0000 0000 0202  ...,............
//   	0x0050:  0202 8000 0002 b354 0018 0006 2003 0000  .......T........
//   	0x0060:  0003 0202 0202 8000 0001 4473 0024 0006  ..........Ds.$..
//   	0x0070:  2003 0000 0002 0202 0202 8000 0001 9c11  ................
//   	0x0080:  0024 0006 2003 0000 0001 0202 0202 8000  .$..............
//   	0x0090:  0001 cdba 0024 0006 2003 0000 0000 0202  .....$..........
//   	0x00a0:  0202 8000 0001 efd7 0024 0005 0008 0000  .........$......
//   	0x00b0:  0005 0202 0202 8000 0001 5433 002c       ..........T3.,
var testPacketOSPF3LSAck = []byte{
	0x33, 0x33, 0x00, 0x00, 0x00, 0x05, 0xc2, 0x00, 0x1f, 0xfa, 0x00, 0x01, 0x86, 0xdd, 0x6e, 0x00,
	0x00, 0x00, 0x00, 0x88, 0x59, 0x01, 0xfe, 0x80, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0xff, 0x02, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x05, 0x03, 0x05, 0x00, 0x88, 0x01, 0x01, 0x01, 0x01, 0x00, 0x00,
	0x00, 0x01, 0x9d, 0x2c, 0x00, 0x00, 0x00, 0x05, 0x20, 0x01, 0x00, 0x00, 0x00, 0x00, 0x02, 0x02,
	0x02, 0x02, 0x80, 0x00, 0x00, 0x02, 0xb3, 0x54, 0x00, 0x18, 0x00, 0x06, 0x20, 0x03, 0x00, 0x00,
	0x00, 0x03, 0x02, 0x02, 0x02, 0x02, 0x80, 0x00, 0x00, 0x01, 0x44, 0x73, 0x00, 0x24, 0x00, 0x06,
	0x20, 0x03, 0x00, 0x00, 0x00, 0x02, 0x02, 0x02, 0x02, 0x02, 0x80, 0x00, 0x00, 0x01, 0x9c, 0x11,
	0x00, 0x24, 0x00, 0x06, 0x20, 0x03, 0x00, 0x00, 0x00, 0x01, 0x02, 0x02, 0x02, 0x02, 0x80, 0x00,
	0x00, 0x01, 0xcd, 0xba, 0x00, 0x24, 0x00, 0x06, 0x20, 0x03, 0x00, 0x00, 0x00, 0x00, 0x02, 0x02,
	0x02, 0x02, 0x80, 0x00, 0x00, 0x01, 0xef, 0xd7, 0x00, 0x24, 0x00, 0x05, 0x00, 0x08, 0x00, 0x00,
	0x00, 0x05, 0x02, 0x02, 0x02, 0x02, 0x80, 0x00, 0x00, 0x01, 0x54, 0x33, 0x00, 0x2c,
}

func TestPacketOSPF3LSAck(t *testing.T) {
	p := gopacket.NewPacket(testPacketOSPF3LSAck, LinkTypeEthernet, gopacket.Default)
	if p.ErrorLayer() != nil {
		t.Error("Failed to decode packet:", p.ErrorLayer().Error())
	}
	checkLayers(p, []gopacket.LayerType{LayerTypeEthernet, LayerTypeIPv6, LayerTypeOSPF}, t)
	if got, ok := p.Layer(LayerTypeOSPF).(*OSPFv3); ok {
		want := &OSPFv3{
			OSPF: OSPF{
				Version:      3,
				Type:         OSPFLinkStateAcknowledgment,
				PacketLength: 136,
				RouterID:     0x1010101,
				AreaID:       1,
				Checksum:     0x9d2c,
				Content: []LSAheader{
					LSAheader{
						LSAge:       5,
						LSType:      0x2001,
						LinkStateID: 0x00000000,
						AdvRouter:   0x02020202,
						LSSeqNumber: 0x80000002,
						LSChecksum:  0xb354,
						Length:      24,
					},
					LSAheader{
						LSAge:       6,
						LSType:      0x2003,
						LinkStateID: 0x00000003,
						AdvRouter:   0x02020202,
						LSSeqNumber: 0x80000001,
						LSChecksum:  0x4473,
						Length:      36,
					},
					LSAheader{
						LSAge:       6,
						LSType:      0x2003,
						LinkStateID: 0x00000002,
						AdvRouter:   0x02020202,
						LSSeqNumber: 0x80000001,
						LSChecksum:  0x9c11,
						Length:      36,
					},
					LSAheader{
						LSAge:       6,
						LSType:      0x2003,
						LinkStateID: 0x00000001,
						AdvRouter:   0x02020202,
						LSSeqNumber: 0x80000001,
						LSChecksum:  0xcdba,
						Length:      36,
					},
					LSAheader{
						LSAge:       6,
						LSType:      0x2003,
						LinkStateID: 0x00000000,
						AdvRouter:   0x02020202,
						LSSeqNumber: 0x80000001,
						LSChecksum:  0xefd7,
						Length:      36,
					},
					LSAheader{
						LSAge:       5,
						LSType:      0x0008,
						LinkStateID: 0x00000005,
						AdvRouter:   0x02020202,
						LSSeqNumber: 0x80000001,
						LSChecksum:  0x5433,
						Length:      44,
					},
				},
			},
			Instance: 0,
			Reserved: 0,
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("OSPF packet processing failed:\ngot  :\n%#v\n\nwant :\n%#v\n\n", got, want)
		}
	} else {
		t.Error("No OSPF layer type found in packet")
	}
}
func BenchmarkDecodePacketPacket4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gopacket.NewPacket(testPacketOSPF3LSAck, LinkTypeEthernet, gopacket.NoCopy)
	}
}
