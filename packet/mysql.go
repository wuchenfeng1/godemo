package main

import (
	"fmt"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

func main() {
	// 得到所有的(网络)设备
	main01("\\Device\\NPF_{13478D8B-663A-47FC-9075-57D9DB23FEA0}")

}
func main01(item string) {
	// 得到所有的(网络)设备
	fmt.Println("数据包开始抓取...")
	snapLen := int32(65535)
	port := uint16(18090)
	filter := getFilter1(port)
	//打开网络接口，抓取在线数据
	handle, err := pcap.OpenLive(item, snapLen, true, pcap.BlockForever)
	if err != nil {
		fmt.Printf("pcap open live failed: %v", err)
		return
	}
	defer handle.Close()
	// 设置过滤器
	if err := handle.SetBPFFilter(filter); err != nil {
		fmt.Printf("set bpf filter failed: %v", err)
		return
	}
	// 抓包
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	packetSource.NoCopy = true
	for packet := range packetSource.Packets() {
		if packet.NetworkLayer() == nil || packet.TransportLayer() == nil || packet.TransportLayer().LayerType() != layers.LayerTypeTCP {
			fmt.Println("意外数据包")
			continue
		}
		fmt.Printf("数据包:#{packet}\n")

		// tcp 层
		tcp := packet.TransportLayer().(*layers.TCP)
		fmt.Printf("tcp:%v\n", tcp)
		// tcp payload，也即是tcp传输的数据
		fmt.Printf("tcp payload:%v\n", tcp.Payload)
	}
}

//定义过滤器
func getFilter1(port uint16) string {
	filter := fmt.Sprintf("tcp and ((src port %v) or (dst port %v))", port, port)
	return filter
}
