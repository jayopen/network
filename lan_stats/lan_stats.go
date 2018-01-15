package main

import (
	"fmt"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

const (
	LISTEN_INTERFACE     = "wlan0"
	LISTEN_INTERFACE_MAC = "ab:cd:ef:gh:ij:kl"
	BROADCAST_MAC        = "ff:ff:ff:ff:ff:ff"
	UNKNOWN_MAC          = "UNKNOWN"
	SNAPLENGTH           = 1600
	PROMISC_MODE         = true
)

var BANNED_MAC_ADDRESSES = []string{BROADCAST_MAC, LISTEN_INTERFACE_MAC}
var lanMap = make(map[string]int)

func main() {
	// Create a PCAP handle to read packets off specified network interface.
	// OpenLive returns a pcap.*handle object. This object satisfies
	// the PacketDataSource interface.
	handle, err := pcap.OpenLive("wlan0", 14, true, pcap.BlockForever)
	if err != nil {
		panic(err)
	}
	// Using an object that satisfie PacketDataSource, we create a
	// gopacket.*PacketSource object.
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())

	// Packets() returns a channel of packets.
	for packet := range packetSource.Packets() {
		length := packet.Metadata().CaptureInfo.Length
		if linkLayer := packet.Layer(layers.LayerTypeEthernet); linkLayer != nil {
			// Type assert packet interface into Ethernet type.
			eth := linkLayer.(*layers.Ethernet)
			mac := GetNeutralMAC(eth.SrcMAC.String(), eth.DstMAC.String())
			lanMap[mac] = lanMap[mac] + int(length)
		} else {
			fmt.Println("Non ethernet packet")
		}
	}
}

// Return a single MAC address that is not broadcast MAC or the LISTEN_INTERFACE_MAC.
func GetNeutralMAC(a string, b string) string {
	for _, banned := range BANNED_MAC_ADDRESSES {
		if a == banned {
			return b
		}
		if b == banned {
			return a
		}
	}
	// Neither a or b is a banned address.
	return UNKNOWN_MAC
}
