package main

/*11.更快地解码数据包
如果我们知道我们要预期的得到的层，我们可以使用现有的结构来存储分组信息，
而不是为每个需要时间和内存的分组创建新的结构。使用DecodingLayerParser更快。
就像编组和解组数据一样。
*/

import (
	"fmt"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"log"
	"time"
)

func main() {
	var (
		// Will reuse these for each packet
		ethLayer layers.Ethernet
		ipLayer  layers.IPv4
		tcpLayer layers.TCP
	)
	// Open device
	handle, err := pcap.OpenLive("\\Device\\NPF_{13478D8B-663A-47FC-9075-57D9DB23FEA0}", 1024, false, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	//var filter = flag.String("f", "tcp and port 8090", "BPF filter for pcap")
	if err := handle.SetBPFFilter("tcp and port 18090"); err != nil {
		log.Println("错误。。。")
	}

	defer handle.Close()

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		parser := gopacket.NewDecodingLayerParser(
			layers.LayerTypeEthernet,
			&ethLayer,
			&ipLayer,
			&tcpLayer,
		)
		var foundLayerTypes []gopacket.LayerType

		err := parser.DecodeLayers(packet.Data(), &foundLayerTypes)
		if err != nil {
			fmt.Println("Trouble decoding layers: ", err)
		}

		for _, layerType := range foundLayerTypes {
			if layerType == layers.LayerTypeIPv4 {
				fmt.Println("IPv4: ", ipLayer.SrcIP, "->", ipLayer.DstIP)
			}
			if layerType == layers.LayerTypeTCP {
				fmt.Println("TCP Port: ", tcpLayer.SrcPort, "->", tcpLayer.DstPort)
				fmt.Println("TCP SYN:", tcpLayer.SYN, " | ACK:", tcpLayer.ACK)
			}
		}
	}
}
