// Copyright 2012 Google, Inc. All rights reserved.
//
// Use of this source code is governed by a BSD-style license
// that can be found in the LICENSE file in the root of the source
// tree.

package layers

import (
	"github.com/njcx/gopacket131_dpdk"
)

var (
	LayerTypeARP                          = gopacket131_dpdk.RegisterLayerType(10, gopacket131_dpdk.LayerTypeMetadata{Name: "ARP", Decoder: gopacket131_dpdk.DecodeFunc(decodeARP)})
	LayerTypeCiscoDiscovery               = gopacket131_dpdk.RegisterLayerType(11, gopacket131_dpdk.LayerTypeMetadata{Name: "CiscoDiscovery", Decoder: gopacket131_dpdk.DecodeFunc(decodeCiscoDiscovery)})
	LayerTypeEthernetCTP                  = gopacket131_dpdk.RegisterLayerType(12, gopacket131_dpdk.LayerTypeMetadata{Name: "EthernetCTP", Decoder: gopacket131_dpdk.DecodeFunc(decodeEthernetCTP)})
	LayerTypeEthernetCTPForwardData       = gopacket131_dpdk.RegisterLayerType(13, gopacket131_dpdk.LayerTypeMetadata{Name: "EthernetCTPForwardData", Decoder: nil})
	LayerTypeEthernetCTPReply             = gopacket131_dpdk.RegisterLayerType(14, gopacket131_dpdk.LayerTypeMetadata{Name: "EthernetCTPReply", Decoder: nil})
	LayerTypeDot1Q                        = gopacket131_dpdk.RegisterLayerType(15, gopacket131_dpdk.LayerTypeMetadata{Name: "Dot1Q", Decoder: gopacket131_dpdk.DecodeFunc(decodeDot1Q)})
	LayerTypeEtherIP                      = gopacket131_dpdk.RegisterLayerType(16, gopacket131_dpdk.LayerTypeMetadata{Name: "EtherIP", Decoder: gopacket131_dpdk.DecodeFunc(decodeEtherIP)})
	LayerTypeEthernet                     = gopacket131_dpdk.RegisterLayerType(17, gopacket131_dpdk.LayerTypeMetadata{Name: "Ethernet", Decoder: gopacket131_dpdk.DecodeFunc(decodeEthernet)})
	LayerTypeGRE                          = gopacket131_dpdk.RegisterLayerType(18, gopacket131_dpdk.LayerTypeMetadata{Name: "GRE", Decoder: gopacket131_dpdk.DecodeFunc(decodeGRE)})
	LayerTypeICMPv4                       = gopacket131_dpdk.RegisterLayerType(19, gopacket131_dpdk.LayerTypeMetadata{Name: "ICMPv4", Decoder: gopacket131_dpdk.DecodeFunc(decodeICMPv4)})
	LayerTypeIPv4                         = gopacket131_dpdk.RegisterLayerType(20, gopacket131_dpdk.LayerTypeMetadata{Name: "IPv4", Decoder: gopacket131_dpdk.DecodeFunc(decodeIPv4)})
	LayerTypeIPv6                         = gopacket131_dpdk.RegisterLayerType(21, gopacket131_dpdk.LayerTypeMetadata{Name: "IPv6", Decoder: gopacket131_dpdk.DecodeFunc(decodeIPv6)})
	LayerTypeLLC                          = gopacket131_dpdk.RegisterLayerType(22, gopacket131_dpdk.LayerTypeMetadata{Name: "LLC", Decoder: gopacket131_dpdk.DecodeFunc(decodeLLC)})
	LayerTypeSNAP                         = gopacket131_dpdk.RegisterLayerType(23, gopacket131_dpdk.LayerTypeMetadata{Name: "SNAP", Decoder: gopacket131_dpdk.DecodeFunc(decodeSNAP)})
	LayerTypeMPLS                         = gopacket131_dpdk.RegisterLayerType(24, gopacket131_dpdk.LayerTypeMetadata{Name: "MPLS", Decoder: gopacket131_dpdk.DecodeFunc(decodeMPLS)})
	LayerTypePPP                          = gopacket131_dpdk.RegisterLayerType(25, gopacket131_dpdk.LayerTypeMetadata{Name: "PPP", Decoder: gopacket131_dpdk.DecodeFunc(decodePPP)})
	LayerTypePPPoE                        = gopacket131_dpdk.RegisterLayerType(26, gopacket131_dpdk.LayerTypeMetadata{Name: "PPPoE", Decoder: gopacket131_dpdk.DecodeFunc(decodePPPoE)})
	LayerTypeRUDP                         = gopacket131_dpdk.RegisterLayerType(27, gopacket131_dpdk.LayerTypeMetadata{Name: "RUDP", Decoder: gopacket131_dpdk.DecodeFunc(decodeRUDP)})
	LayerTypeSCTP                         = gopacket131_dpdk.RegisterLayerType(28, gopacket131_dpdk.LayerTypeMetadata{Name: "SCTP", Decoder: gopacket131_dpdk.DecodeFunc(decodeSCTP)})
	LayerTypeSCTPUnknownChunkType         = gopacket131_dpdk.RegisterLayerType(29, gopacket131_dpdk.LayerTypeMetadata{Name: "SCTPUnknownChunkType", Decoder: nil})
	LayerTypeSCTPData                     = gopacket131_dpdk.RegisterLayerType(30, gopacket131_dpdk.LayerTypeMetadata{Name: "SCTPData", Decoder: nil})
	LayerTypeSCTPInit                     = gopacket131_dpdk.RegisterLayerType(31, gopacket131_dpdk.LayerTypeMetadata{Name: "SCTPInit", Decoder: nil})
	LayerTypeSCTPSack                     = gopacket131_dpdk.RegisterLayerType(32, gopacket131_dpdk.LayerTypeMetadata{Name: "SCTPSack", Decoder: nil})
	LayerTypeSCTPHeartbeat                = gopacket131_dpdk.RegisterLayerType(33, gopacket131_dpdk.LayerTypeMetadata{Name: "SCTPHeartbeat", Decoder: nil})
	LayerTypeSCTPError                    = gopacket131_dpdk.RegisterLayerType(34, gopacket131_dpdk.LayerTypeMetadata{Name: "SCTPError", Decoder: nil})
	LayerTypeSCTPShutdown                 = gopacket131_dpdk.RegisterLayerType(35, gopacket131_dpdk.LayerTypeMetadata{Name: "SCTPShutdown", Decoder: nil})
	LayerTypeSCTPShutdownAck              = gopacket131_dpdk.RegisterLayerType(36, gopacket131_dpdk.LayerTypeMetadata{Name: "SCTPShutdownAck", Decoder: nil})
	LayerTypeSCTPCookieEcho               = gopacket131_dpdk.RegisterLayerType(37, gopacket131_dpdk.LayerTypeMetadata{Name: "SCTPCookieEcho", Decoder: nil})
	LayerTypeSCTPEmptyLayer               = gopacket131_dpdk.RegisterLayerType(38, gopacket131_dpdk.LayerTypeMetadata{Name: "SCTPEmptyLayer", Decoder: nil})
	LayerTypeSCTPInitAck                  = gopacket131_dpdk.RegisterLayerType(39, gopacket131_dpdk.LayerTypeMetadata{Name: "SCTPInitAck", Decoder: nil})
	LayerTypeSCTPHeartbeatAck             = gopacket131_dpdk.RegisterLayerType(40, gopacket131_dpdk.LayerTypeMetadata{Name: "SCTPHeartbeatAck", Decoder: nil})
	LayerTypeSCTPAbort                    = gopacket131_dpdk.RegisterLayerType(41, gopacket131_dpdk.LayerTypeMetadata{Name: "SCTPAbort", Decoder: nil})
	LayerTypeSCTPShutdownComplete         = gopacket131_dpdk.RegisterLayerType(42, gopacket131_dpdk.LayerTypeMetadata{Name: "SCTPShutdownComplete", Decoder: nil})
	LayerTypeSCTPCookieAck                = gopacket131_dpdk.RegisterLayerType(43, gopacket131_dpdk.LayerTypeMetadata{Name: "SCTPCookieAck", Decoder: nil})
	LayerTypeTCP                          = gopacket131_dpdk.RegisterLayerType(44, gopacket131_dpdk.LayerTypeMetadata{Name: "TCP", Decoder: gopacket131_dpdk.DecodeFunc(decodeTCP)})
	LayerTypeUDP                          = gopacket131_dpdk.RegisterLayerType(45, gopacket131_dpdk.LayerTypeMetadata{Name: "UDP", Decoder: gopacket131_dpdk.DecodeFunc(decodeUDP)})
	LayerTypeIPv6HopByHop                 = gopacket131_dpdk.RegisterLayerType(46, gopacket131_dpdk.LayerTypeMetadata{Name: "IPv6HopByHop", Decoder: gopacket131_dpdk.DecodeFunc(decodeIPv6HopByHop)})
	LayerTypeIPv6Routing                  = gopacket131_dpdk.RegisterLayerType(47, gopacket131_dpdk.LayerTypeMetadata{Name: "IPv6Routing", Decoder: gopacket131_dpdk.DecodeFunc(decodeIPv6Routing)})
	LayerTypeIPv6Fragment                 = gopacket131_dpdk.RegisterLayerType(48, gopacket131_dpdk.LayerTypeMetadata{Name: "IPv6Fragment", Decoder: gopacket131_dpdk.DecodeFunc(decodeIPv6Fragment)})
	LayerTypeIPv6Destination              = gopacket131_dpdk.RegisterLayerType(49, gopacket131_dpdk.LayerTypeMetadata{Name: "IPv6Destination", Decoder: gopacket131_dpdk.DecodeFunc(decodeIPv6Destination)})
	LayerTypeIPSecAH                      = gopacket131_dpdk.RegisterLayerType(50, gopacket131_dpdk.LayerTypeMetadata{Name: "IPSecAH", Decoder: gopacket131_dpdk.DecodeFunc(decodeIPSecAH)})
	LayerTypeIPSecESP                     = gopacket131_dpdk.RegisterLayerType(51, gopacket131_dpdk.LayerTypeMetadata{Name: "IPSecESP", Decoder: gopacket131_dpdk.DecodeFunc(decodeIPSecESP)})
	LayerTypeUDPLite                      = gopacket131_dpdk.RegisterLayerType(52, gopacket131_dpdk.LayerTypeMetadata{Name: "UDPLite", Decoder: gopacket131_dpdk.DecodeFunc(decodeUDPLite)})
	LayerTypeFDDI                         = gopacket131_dpdk.RegisterLayerType(53, gopacket131_dpdk.LayerTypeMetadata{Name: "FDDI", Decoder: gopacket131_dpdk.DecodeFunc(decodeFDDI)})
	LayerTypeLoopback                     = gopacket131_dpdk.RegisterLayerType(54, gopacket131_dpdk.LayerTypeMetadata{Name: "Loopback", Decoder: gopacket131_dpdk.DecodeFunc(decodeLoopback)})
	LayerTypeEAP                          = gopacket131_dpdk.RegisterLayerType(55, gopacket131_dpdk.LayerTypeMetadata{Name: "EAP", Decoder: gopacket131_dpdk.DecodeFunc(decodeEAP)})
	LayerTypeEAPOL                        = gopacket131_dpdk.RegisterLayerType(56, gopacket131_dpdk.LayerTypeMetadata{Name: "EAPOL", Decoder: gopacket131_dpdk.DecodeFunc(decodeEAPOL)})
	LayerTypeICMPv6                       = gopacket131_dpdk.RegisterLayerType(57, gopacket131_dpdk.LayerTypeMetadata{Name: "ICMPv6", Decoder: gopacket131_dpdk.DecodeFunc(decodeICMPv6)})
	LayerTypeLinkLayerDiscovery           = gopacket131_dpdk.RegisterLayerType(58, gopacket131_dpdk.LayerTypeMetadata{Name: "LinkLayerDiscovery", Decoder: gopacket131_dpdk.DecodeFunc(decodeLinkLayerDiscovery)})
	LayerTypeCiscoDiscoveryInfo           = gopacket131_dpdk.RegisterLayerType(59, gopacket131_dpdk.LayerTypeMetadata{Name: "CiscoDiscoveryInfo", Decoder: gopacket131_dpdk.DecodeFunc(decodeCiscoDiscoveryInfo)})
	LayerTypeLinkLayerDiscoveryInfo       = gopacket131_dpdk.RegisterLayerType(60, gopacket131_dpdk.LayerTypeMetadata{Name: "LinkLayerDiscoveryInfo", Decoder: nil})
	LayerTypeNortelDiscovery              = gopacket131_dpdk.RegisterLayerType(61, gopacket131_dpdk.LayerTypeMetadata{Name: "NortelDiscovery", Decoder: gopacket131_dpdk.DecodeFunc(decodeNortelDiscovery)})
	LayerTypeIGMP                         = gopacket131_dpdk.RegisterLayerType(62, gopacket131_dpdk.LayerTypeMetadata{Name: "IGMP", Decoder: gopacket131_dpdk.DecodeFunc(decodeIGMP)})
	LayerTypePFLog                        = gopacket131_dpdk.RegisterLayerType(63, gopacket131_dpdk.LayerTypeMetadata{Name: "PFLog", Decoder: gopacket131_dpdk.DecodeFunc(decodePFLog)})
	LayerTypeRadioTap                     = gopacket131_dpdk.RegisterLayerType(64, gopacket131_dpdk.LayerTypeMetadata{Name: "RadioTap", Decoder: gopacket131_dpdk.DecodeFunc(decodeRadioTap)})
	LayerTypeDot11                        = gopacket131_dpdk.RegisterLayerType(65, gopacket131_dpdk.LayerTypeMetadata{Name: "Dot11", Decoder: gopacket131_dpdk.DecodeFunc(decodeDot11)})
	LayerTypeDot11Ctrl                    = gopacket131_dpdk.RegisterLayerType(66, gopacket131_dpdk.LayerTypeMetadata{Name: "Dot11Ctrl", Decoder: gopacket131_dpdk.DecodeFunc(decodeDot11Ctrl)})
	LayerTypeDot11Data                    = gopacket131_dpdk.RegisterLayerType(67, gopacket131_dpdk.LayerTypeMetadata{Name: "Dot11Data", Decoder: gopacket131_dpdk.DecodeFunc(decodeDot11Data)})
	LayerTypeDot11DataCFAck               = gopacket131_dpdk.RegisterLayerType(68, gopacket131_dpdk.LayerTypeMetadata{Name: "Dot11DataCFAck", Decoder: gopacket131_dpdk.DecodeFunc(decodeDot11DataCFAck)})
	LayerTypeDot11DataCFPoll              = gopacket131_dpdk.RegisterLayerType(69, gopacket131_dpdk.LayerTypeMetadata{Name: "Dot11DataCFPoll", Decoder: gopacket131_dpdk.DecodeFunc(decodeDot11DataCFPoll)})
	LayerTypeDot11DataCFAckPoll           = gopacket131_dpdk.RegisterLayerType(70, gopacket131_dpdk.LayerTypeMetadata{Name: "Dot11DataCFAckPoll", Decoder: gopacket131_dpdk.DecodeFunc(decodeDot11DataCFAckPoll)})
	LayerTypeDot11DataNull                = gopacket131_dpdk.RegisterLayerType(71, gopacket131_dpdk.LayerTypeMetadata{Name: "Dot11DataNull", Decoder: gopacket131_dpdk.DecodeFunc(decodeDot11DataNull)})
	LayerTypeDot11DataCFAckNoData         = gopacket131_dpdk.RegisterLayerType(72, gopacket131_dpdk.LayerTypeMetadata{Name: "Dot11DataCFAck", Decoder: gopacket131_dpdk.DecodeFunc(decodeDot11DataCFAck)})
	LayerTypeDot11DataCFPollNoData        = gopacket131_dpdk.RegisterLayerType(73, gopacket131_dpdk.LayerTypeMetadata{Name: "Dot11DataCFPoll", Decoder: gopacket131_dpdk.DecodeFunc(decodeDot11DataCFPoll)})
	LayerTypeDot11DataCFAckPollNoData     = gopacket131_dpdk.RegisterLayerType(74, gopacket131_dpdk.LayerTypeMetadata{Name: "Dot11DataCFAckPoll", Decoder: gopacket131_dpdk.DecodeFunc(decodeDot11DataCFAckPoll)})
	LayerTypeDot11DataQOSData             = gopacket131_dpdk.RegisterLayerType(75, gopacket131_dpdk.LayerTypeMetadata{Name: "Dot11DataQOSData", Decoder: gopacket131_dpdk.DecodeFunc(decodeDot11DataQOSData)})
	LayerTypeDot11DataQOSDataCFAck        = gopacket131_dpdk.RegisterLayerType(76, gopacket131_dpdk.LayerTypeMetadata{Name: "Dot11DataQOSDataCFAck", Decoder: gopacket131_dpdk.DecodeFunc(decodeDot11DataQOSDataCFAck)})
	LayerTypeDot11DataQOSDataCFPoll       = gopacket131_dpdk.RegisterLayerType(77, gopacket131_dpdk.LayerTypeMetadata{Name: "Dot11DataQOSDataCFPoll", Decoder: gopacket131_dpdk.DecodeFunc(decodeDot11DataQOSDataCFPoll)})
	LayerTypeDot11DataQOSDataCFAckPoll    = gopacket131_dpdk.RegisterLayerType(78, gopacket131_dpdk.LayerTypeMetadata{Name: "Dot11DataQOSDataCFAckPoll", Decoder: gopacket131_dpdk.DecodeFunc(decodeDot11DataQOSDataCFAckPoll)})
	LayerTypeDot11DataQOSNull             = gopacket131_dpdk.RegisterLayerType(79, gopacket131_dpdk.LayerTypeMetadata{Name: "Dot11DataQOSNull", Decoder: gopacket131_dpdk.DecodeFunc(decodeDot11DataQOSNull)})
	LayerTypeDot11DataQOSCFPollNoData     = gopacket131_dpdk.RegisterLayerType(80, gopacket131_dpdk.LayerTypeMetadata{Name: "Dot11DataQOSCFPoll", Decoder: gopacket131_dpdk.DecodeFunc(decodeDot11DataQOSCFPollNoData)})
	LayerTypeDot11DataQOSCFAckPollNoData  = gopacket131_dpdk.RegisterLayerType(81, gopacket131_dpdk.LayerTypeMetadata{Name: "Dot11DataQOSCFAckPoll", Decoder: gopacket131_dpdk.DecodeFunc(decodeDot11DataQOSCFAckPollNoData)})
	LayerTypeDot11InformationElement      = gopacket131_dpdk.RegisterLayerType(82, gopacket131_dpdk.LayerTypeMetadata{Name: "Dot11InformationElement", Decoder: gopacket131_dpdk.DecodeFunc(decodeDot11InformationElement)})
	LayerTypeDot11CtrlCTS                 = gopacket131_dpdk.RegisterLayerType(83, gopacket131_dpdk.LayerTypeMetadata{Name: "Dot11CtrlCTS", Decoder: gopacket131_dpdk.DecodeFunc(decodeDot11CtrlCTS)})
	LayerTypeDot11CtrlRTS                 = gopacket131_dpdk.RegisterLayerType(84, gopacket131_dpdk.LayerTypeMetadata{Name: "Dot11CtrlRTS", Decoder: gopacket131_dpdk.DecodeFunc(decodeDot11CtrlRTS)})
	LayerTypeDot11CtrlBlockAckReq         = gopacket131_dpdk.RegisterLayerType(85, gopacket131_dpdk.LayerTypeMetadata{Name: "Dot11CtrlBlockAckReq", Decoder: gopacket131_dpdk.DecodeFunc(decodeDot11CtrlBlockAckReq)})
	LayerTypeDot11CtrlBlockAck            = gopacket131_dpdk.RegisterLayerType(86, gopacket131_dpdk.LayerTypeMetadata{Name: "Dot11CtrlBlockAck", Decoder: gopacket131_dpdk.DecodeFunc(decodeDot11CtrlBlockAck)})
	LayerTypeDot11CtrlPowersavePoll       = gopacket131_dpdk.RegisterLayerType(87, gopacket131_dpdk.LayerTypeMetadata{Name: "Dot11CtrlPowersavePoll", Decoder: gopacket131_dpdk.DecodeFunc(decodeDot11CtrlPowersavePoll)})
	LayerTypeDot11CtrlAck                 = gopacket131_dpdk.RegisterLayerType(88, gopacket131_dpdk.LayerTypeMetadata{Name: "Dot11CtrlAck", Decoder: gopacket131_dpdk.DecodeFunc(decodeDot11CtrlAck)})
	LayerTypeDot11CtrlCFEnd               = gopacket131_dpdk.RegisterLayerType(89, gopacket131_dpdk.LayerTypeMetadata{Name: "Dot11CtrlCFEnd", Decoder: gopacket131_dpdk.DecodeFunc(decodeDot11CtrlCFEnd)})
	LayerTypeDot11CtrlCFEndAck            = gopacket131_dpdk.RegisterLayerType(90, gopacket131_dpdk.LayerTypeMetadata{Name: "Dot11CtrlCFEndAck", Decoder: gopacket131_dpdk.DecodeFunc(decodeDot11CtrlCFEndAck)})
	LayerTypeDot11MgmtAssociationReq      = gopacket131_dpdk.RegisterLayerType(91, gopacket131_dpdk.LayerTypeMetadata{Name: "Dot11MgmtAssociationReq", Decoder: gopacket131_dpdk.DecodeFunc(decodeDot11MgmtAssociationReq)})
	LayerTypeDot11MgmtAssociationResp     = gopacket131_dpdk.RegisterLayerType(92, gopacket131_dpdk.LayerTypeMetadata{Name: "Dot11MgmtAssociationResp", Decoder: gopacket131_dpdk.DecodeFunc(decodeDot11MgmtAssociationResp)})
	LayerTypeDot11MgmtReassociationReq    = gopacket131_dpdk.RegisterLayerType(93, gopacket131_dpdk.LayerTypeMetadata{Name: "Dot11MgmtReassociationReq", Decoder: gopacket131_dpdk.DecodeFunc(decodeDot11MgmtReassociationReq)})
	LayerTypeDot11MgmtReassociationResp   = gopacket131_dpdk.RegisterLayerType(94, gopacket131_dpdk.LayerTypeMetadata{Name: "Dot11MgmtReassociationResp", Decoder: gopacket131_dpdk.DecodeFunc(decodeDot11MgmtReassociationResp)})
	LayerTypeDot11MgmtProbeReq            = gopacket131_dpdk.RegisterLayerType(95, gopacket131_dpdk.LayerTypeMetadata{Name: "Dot11MgmtProbeReq", Decoder: gopacket131_dpdk.DecodeFunc(decodeDot11MgmtProbeReq)})
	LayerTypeDot11MgmtProbeResp           = gopacket131_dpdk.RegisterLayerType(96, gopacket131_dpdk.LayerTypeMetadata{Name: "Dot11MgmtProbeResp", Decoder: gopacket131_dpdk.DecodeFunc(decodeDot11MgmtProbeResp)})
	LayerTypeDot11MgmtMeasurementPilot    = gopacket131_dpdk.RegisterLayerType(97, gopacket131_dpdk.LayerTypeMetadata{Name: "Dot11MgmtMeasurementPilot", Decoder: gopacket131_dpdk.DecodeFunc(decodeDot11MgmtMeasurementPilot)})
	LayerTypeDot11MgmtBeacon              = gopacket131_dpdk.RegisterLayerType(98, gopacket131_dpdk.LayerTypeMetadata{Name: "Dot11MgmtBeacon", Decoder: gopacket131_dpdk.DecodeFunc(decodeDot11MgmtBeacon)})
	LayerTypeDot11MgmtATIM                = gopacket131_dpdk.RegisterLayerType(99, gopacket131_dpdk.LayerTypeMetadata{Name: "Dot11MgmtATIM", Decoder: gopacket131_dpdk.DecodeFunc(decodeDot11MgmtATIM)})
	LayerTypeDot11MgmtDisassociation      = gopacket131_dpdk.RegisterLayerType(100, gopacket131_dpdk.LayerTypeMetadata{Name: "Dot11MgmtDisassociation", Decoder: gopacket131_dpdk.DecodeFunc(decodeDot11MgmtDisassociation)})
	LayerTypeDot11MgmtAuthentication      = gopacket131_dpdk.RegisterLayerType(101, gopacket131_dpdk.LayerTypeMetadata{Name: "Dot11MgmtAuthentication", Decoder: gopacket131_dpdk.DecodeFunc(decodeDot11MgmtAuthentication)})
	LayerTypeDot11MgmtDeauthentication    = gopacket131_dpdk.RegisterLayerType(102, gopacket131_dpdk.LayerTypeMetadata{Name: "Dot11MgmtDeauthentication", Decoder: gopacket131_dpdk.DecodeFunc(decodeDot11MgmtDeauthentication)})
	LayerTypeDot11MgmtAction              = gopacket131_dpdk.RegisterLayerType(103, gopacket131_dpdk.LayerTypeMetadata{Name: "Dot11MgmtAction", Decoder: gopacket131_dpdk.DecodeFunc(decodeDot11MgmtAction)})
	LayerTypeDot11MgmtActionNoAck         = gopacket131_dpdk.RegisterLayerType(104, gopacket131_dpdk.LayerTypeMetadata{Name: "Dot11MgmtActionNoAck", Decoder: gopacket131_dpdk.DecodeFunc(decodeDot11MgmtActionNoAck)})
	LayerTypeDot11MgmtArubaWLAN           = gopacket131_dpdk.RegisterLayerType(105, gopacket131_dpdk.LayerTypeMetadata{Name: "Dot11MgmtArubaWLAN", Decoder: gopacket131_dpdk.DecodeFunc(decodeDot11MgmtArubaWLAN)})
	LayerTypeDot11WEP                     = gopacket131_dpdk.RegisterLayerType(106, gopacket131_dpdk.LayerTypeMetadata{Name: "Dot11WEP", Decoder: gopacket131_dpdk.DecodeFunc(decodeDot11WEP)})
	LayerTypeDNS                          = gopacket131_dpdk.RegisterLayerType(107, gopacket131_dpdk.LayerTypeMetadata{Name: "DNS", Decoder: gopacket131_dpdk.DecodeFunc(decodeDNS)})
	LayerTypeUSB                          = gopacket131_dpdk.RegisterLayerType(108, gopacket131_dpdk.LayerTypeMetadata{Name: "USB", Decoder: gopacket131_dpdk.DecodeFunc(decodeUSB)})
	LayerTypeUSBRequestBlockSetup         = gopacket131_dpdk.RegisterLayerType(109, gopacket131_dpdk.LayerTypeMetadata{Name: "USBRequestBlockSetup", Decoder: gopacket131_dpdk.DecodeFunc(decodeUSBRequestBlockSetup)})
	LayerTypeUSBControl                   = gopacket131_dpdk.RegisterLayerType(110, gopacket131_dpdk.LayerTypeMetadata{Name: "USBControl", Decoder: gopacket131_dpdk.DecodeFunc(decodeUSBControl)})
	LayerTypeUSBInterrupt                 = gopacket131_dpdk.RegisterLayerType(111, gopacket131_dpdk.LayerTypeMetadata{Name: "USBInterrupt", Decoder: gopacket131_dpdk.DecodeFunc(decodeUSBInterrupt)})
	LayerTypeUSBBulk                      = gopacket131_dpdk.RegisterLayerType(112, gopacket131_dpdk.LayerTypeMetadata{Name: "USBBulk", Decoder: gopacket131_dpdk.DecodeFunc(decodeUSBBulk)})
	LayerTypeLinuxSLL                     = gopacket131_dpdk.RegisterLayerType(113, gopacket131_dpdk.LayerTypeMetadata{Name: "Linux SLL", Decoder: gopacket131_dpdk.DecodeFunc(decodeLinuxSLL)})
	LayerTypeSFlow                        = gopacket131_dpdk.RegisterLayerType(114, gopacket131_dpdk.LayerTypeMetadata{Name: "SFlow", Decoder: gopacket131_dpdk.DecodeFunc(decodeSFlow)})
	LayerTypePrismHeader                  = gopacket131_dpdk.RegisterLayerType(115, gopacket131_dpdk.LayerTypeMetadata{Name: "Prism monitor mode header", Decoder: gopacket131_dpdk.DecodeFunc(decodePrismHeader)})
	LayerTypeVXLAN                        = gopacket131_dpdk.RegisterLayerType(116, gopacket131_dpdk.LayerTypeMetadata{Name: "VXLAN", Decoder: gopacket131_dpdk.DecodeFunc(decodeVXLAN)})
	LayerTypeNTP                          = gopacket131_dpdk.RegisterLayerType(117, gopacket131_dpdk.LayerTypeMetadata{Name: "NTP", Decoder: gopacket131_dpdk.DecodeFunc(decodeNTP)})
	LayerTypeDHCPv4                       = gopacket131_dpdk.RegisterLayerType(118, gopacket131_dpdk.LayerTypeMetadata{Name: "DHCPv4", Decoder: gopacket131_dpdk.DecodeFunc(decodeDHCPv4)})
	LayerTypeVRRP                         = gopacket131_dpdk.RegisterLayerType(119, gopacket131_dpdk.LayerTypeMetadata{Name: "VRRP", Decoder: gopacket131_dpdk.DecodeFunc(decodeVRRP)})
	LayerTypeGeneve                       = gopacket131_dpdk.RegisterLayerType(120, gopacket131_dpdk.LayerTypeMetadata{Name: "Geneve", Decoder: gopacket131_dpdk.DecodeFunc(decodeGeneve)})
	LayerTypeSTP                          = gopacket131_dpdk.RegisterLayerType(121, gopacket131_dpdk.LayerTypeMetadata{Name: "STP", Decoder: gopacket131_dpdk.DecodeFunc(decodeSTP)})
	LayerTypeBFD                          = gopacket131_dpdk.RegisterLayerType(122, gopacket131_dpdk.LayerTypeMetadata{Name: "BFD", Decoder: gopacket131_dpdk.DecodeFunc(decodeBFD)})
	LayerTypeOSPF                         = gopacket131_dpdk.RegisterLayerType(123, gopacket131_dpdk.LayerTypeMetadata{Name: "OSPF", Decoder: gopacket131_dpdk.DecodeFunc(decodeOSPF)})
	LayerTypeICMPv6RouterSolicitation     = gopacket131_dpdk.RegisterLayerType(124, gopacket131_dpdk.LayerTypeMetadata{Name: "ICMPv6RouterSolicitation", Decoder: gopacket131_dpdk.DecodeFunc(decodeICMPv6RouterSolicitation)})
	LayerTypeICMPv6RouterAdvertisement    = gopacket131_dpdk.RegisterLayerType(125, gopacket131_dpdk.LayerTypeMetadata{Name: "ICMPv6RouterAdvertisement", Decoder: gopacket131_dpdk.DecodeFunc(decodeICMPv6RouterAdvertisement)})
	LayerTypeICMPv6NeighborSolicitation   = gopacket131_dpdk.RegisterLayerType(126, gopacket131_dpdk.LayerTypeMetadata{Name: "ICMPv6NeighborSolicitation", Decoder: gopacket131_dpdk.DecodeFunc(decodeICMPv6NeighborSolicitation)})
	LayerTypeICMPv6NeighborAdvertisement  = gopacket131_dpdk.RegisterLayerType(127, gopacket131_dpdk.LayerTypeMetadata{Name: "ICMPv6NeighborAdvertisement", Decoder: gopacket131_dpdk.DecodeFunc(decodeICMPv6NeighborAdvertisement)})
	LayerTypeICMPv6Redirect               = gopacket131_dpdk.RegisterLayerType(128, gopacket131_dpdk.LayerTypeMetadata{Name: "ICMPv6Redirect", Decoder: gopacket131_dpdk.DecodeFunc(decodeICMPv6Redirect)})
	LayerTypeGTPv1U                       = gopacket131_dpdk.RegisterLayerType(129, gopacket131_dpdk.LayerTypeMetadata{Name: "GTPv1U", Decoder: gopacket131_dpdk.DecodeFunc(decodeGTPv1u)})
	LayerTypeEAPOLKey                     = gopacket131_dpdk.RegisterLayerType(130, gopacket131_dpdk.LayerTypeMetadata{Name: "EAPOLKey", Decoder: gopacket131_dpdk.DecodeFunc(decodeEAPOLKey)})
	LayerTypeLCM                          = gopacket131_dpdk.RegisterLayerType(131, gopacket131_dpdk.LayerTypeMetadata{Name: "LCM", Decoder: gopacket131_dpdk.DecodeFunc(decodeLCM)})
	LayerTypeICMPv6Echo                   = gopacket131_dpdk.RegisterLayerType(132, gopacket131_dpdk.LayerTypeMetadata{Name: "ICMPv6Echo", Decoder: gopacket131_dpdk.DecodeFunc(decodeICMPv6Echo)})
	LayerTypeSIP                          = gopacket131_dpdk.RegisterLayerType(133, gopacket131_dpdk.LayerTypeMetadata{Name: "SIP", Decoder: gopacket131_dpdk.DecodeFunc(decodeSIP)})
	LayerTypeDHCPv6                       = gopacket131_dpdk.RegisterLayerType(134, gopacket131_dpdk.LayerTypeMetadata{Name: "DHCPv6", Decoder: gopacket131_dpdk.DecodeFunc(decodeDHCPv6)})
	LayerTypeMLDv1MulticastListenerReport = gopacket131_dpdk.RegisterLayerType(135, gopacket131_dpdk.LayerTypeMetadata{Name: "MLDv1MulticastListenerReport", Decoder: gopacket131_dpdk.DecodeFunc(decodeMLDv1MulticastListenerReport)})
	LayerTypeMLDv1MulticastListenerDone   = gopacket131_dpdk.RegisterLayerType(136, gopacket131_dpdk.LayerTypeMetadata{Name: "MLDv1MulticastListenerDone", Decoder: gopacket131_dpdk.DecodeFunc(decodeMLDv1MulticastListenerDone)})
	LayerTypeMLDv1MulticastListenerQuery  = gopacket131_dpdk.RegisterLayerType(137, gopacket131_dpdk.LayerTypeMetadata{Name: "MLDv1MulticastListenerQuery", Decoder: gopacket131_dpdk.DecodeFunc(decodeMLDv1MulticastListenerQuery)})
	LayerTypeMLDv2MulticastListenerReport = gopacket131_dpdk.RegisterLayerType(138, gopacket131_dpdk.LayerTypeMetadata{Name: "MLDv2MulticastListenerReport", Decoder: gopacket131_dpdk.DecodeFunc(decodeMLDv2MulticastListenerReport)})
	LayerTypeMLDv2MulticastListenerQuery  = gopacket131_dpdk.RegisterLayerType(139, gopacket131_dpdk.LayerTypeMetadata{Name: "MLDv2MulticastListenerQuery", Decoder: gopacket131_dpdk.DecodeFunc(decodeMLDv2MulticastListenerQuery)})
	LayerTypeTLS                          = gopacket131_dpdk.RegisterLayerType(140, gopacket131_dpdk.LayerTypeMetadata{Name: "TLS", Decoder: gopacket131_dpdk.DecodeFunc(decodeTLS)})
	LayerTypeModbusTCP                    = gopacket131_dpdk.RegisterLayerType(141, gopacket131_dpdk.LayerTypeMetadata{Name: "ModbusTCP", Decoder: gopacket131_dpdk.DecodeFunc(decodeModbusTCP)})
	LayerTypeRMCP                         = gopacket131_dpdk.RegisterLayerType(142, gopacket131_dpdk.LayerTypeMetadata{Name: "RMCP", Decoder: gopacket131_dpdk.DecodeFunc(decodeRMCP)})
	LayerTypeASF                          = gopacket131_dpdk.RegisterLayerType(143, gopacket131_dpdk.LayerTypeMetadata{Name: "ASF", Decoder: gopacket131_dpdk.DecodeFunc(decodeASF)})
	LayerTypeASFPresencePong              = gopacket131_dpdk.RegisterLayerType(144, gopacket131_dpdk.LayerTypeMetadata{Name: "ASFPresencePong", Decoder: gopacket131_dpdk.DecodeFunc(decodeASFPresencePong)})
	LayerTypeERSPANII                     = gopacket131_dpdk.RegisterLayerType(145, gopacket131_dpdk.LayerTypeMetadata{Name: "ERSPAN Type II", Decoder: gopacket131_dpdk.DecodeFunc(decodeERSPANII)})
	LayerTypeRADIUS                       = gopacket131_dpdk.RegisterLayerType(146, gopacket131_dpdk.LayerTypeMetadata{Name: "RADIUS", Decoder: gopacket131_dpdk.DecodeFunc(decodeRADIUS)})
	LayerTypeLinuxSLL2                    = gopacket131_dpdk.RegisterLayerType(276, gopacket131_dpdk.LayerTypeMetadata{Name: "Linux SLL2", Decoder: gopacket131_dpdk.DecodeFunc(decodeLinuxSLL2)})
	LayerTypeMDP                          = gopacket131_dpdk.RegisterLayerType(147, gopacket131_dpdk.LayerTypeMetadata{Name: "MDP", Decoder: gopacket131_dpdk.DecodeFunc(decodeMDP)})
)

var (
	// LayerClassIPNetwork contains TCP/IP network layer types.
	LayerClassIPNetwork = gopacket131_dpdk.NewLayerClass([]gopacket131_dpdk.LayerType{
		LayerTypeIPv4,
		LayerTypeIPv6,
	})
	// LayerClassIPTransport contains TCP/IP transport layer types.
	LayerClassIPTransport = gopacket131_dpdk.NewLayerClass([]gopacket131_dpdk.LayerType{
		LayerTypeTCP,
		LayerTypeUDP,
		LayerTypeSCTP,
	})
	// LayerClassIPControl contains TCP/IP control protocols.
	LayerClassIPControl = gopacket131_dpdk.NewLayerClass([]gopacket131_dpdk.LayerType{
		LayerTypeICMPv4,
		LayerTypeICMPv6,
	})
	// LayerClassSCTPChunk contains SCTP chunk types (not the top-level SCTP
	// layer).
	LayerClassSCTPChunk = gopacket131_dpdk.NewLayerClass([]gopacket131_dpdk.LayerType{
		LayerTypeSCTPUnknownChunkType,
		LayerTypeSCTPData,
		LayerTypeSCTPInit,
		LayerTypeSCTPSack,
		LayerTypeSCTPHeartbeat,
		LayerTypeSCTPError,
		LayerTypeSCTPShutdown,
		LayerTypeSCTPShutdownAck,
		LayerTypeSCTPCookieEcho,
		LayerTypeSCTPEmptyLayer,
		LayerTypeSCTPInitAck,
		LayerTypeSCTPHeartbeatAck,
		LayerTypeSCTPAbort,
		LayerTypeSCTPShutdownComplete,
		LayerTypeSCTPCookieAck,
	})
	// LayerClassIPv6Extension contains IPv6 extension headers.
	LayerClassIPv6Extension = gopacket131_dpdk.NewLayerClass([]gopacket131_dpdk.LayerType{
		LayerTypeIPv6HopByHop,
		LayerTypeIPv6Routing,
		LayerTypeIPv6Fragment,
		LayerTypeIPv6Destination,
	})
	LayerClassIPSec = gopacket131_dpdk.NewLayerClass([]gopacket131_dpdk.LayerType{
		LayerTypeIPSecAH,
		LayerTypeIPSecESP,
	})
	// LayerClassICMPv6NDP contains ICMPv6 neighbor discovery protocol
	// messages.
	LayerClassICMPv6NDP = gopacket131_dpdk.NewLayerClass([]gopacket131_dpdk.LayerType{
		LayerTypeICMPv6RouterSolicitation,
		LayerTypeICMPv6RouterAdvertisement,
		LayerTypeICMPv6NeighborSolicitation,
		LayerTypeICMPv6NeighborAdvertisement,
		LayerTypeICMPv6Redirect,
	})
	// LayerClassMLDv1 contains multicast listener discovery protocol
	LayerClassMLDv1 = gopacket131_dpdk.NewLayerClass([]gopacket131_dpdk.LayerType{
		LayerTypeMLDv1MulticastListenerQuery,
		LayerTypeMLDv1MulticastListenerReport,
		LayerTypeMLDv1MulticastListenerDone,
	})
	// LayerClassMLDv2 contains multicast listener discovery protocol v2
	LayerClassMLDv2 = gopacket131_dpdk.NewLayerClass([]gopacket131_dpdk.LayerType{
		LayerTypeMLDv1MulticastListenerReport,
		LayerTypeMLDv1MulticastListenerDone,
		LayerTypeMLDv2MulticastListenerReport,
		LayerTypeMLDv1MulticastListenerQuery,
		LayerTypeMLDv2MulticastListenerQuery,
	})
)
