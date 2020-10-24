package main

//
//
////解码抓取的数据
///*我们可以使用原始数据包，并且可将其转换为已知格式。它与不同的层兼容，
//所以我们可以轻松访问以太网，IP和TCP层。layers包是Go库中新增的，
//在底层pcap库中不可用。这是一个令人难以置信的有用的包，它是gopacket库的一部分。
//它允许我们容易地识别包是否包含特定类型的层。该代码示例将显示如何使用layers包来查看数据包是以太网，IP和TCP，
//并轻松访问这些头文件中的元素。
//查找有效载荷取决于所涉及的所有层。每个协议是不同的，必须相应地计算。这就是layer包的魅力所在。 gopacket的作者花了时间为诸如以太网，IP，
//UDP和TCP等众多已知层创建了相应类型。有效载荷是应用层的一部分。
//*/
//
//import (
//	"fmt"
//	"github.com/google/gopacket"
//	"github.com/google/gopacket/layers"
//	"github.com/google/gopacket/pcap"
//	"log"
//	"strings"
//)
//
//////定义过滤器
////func getFilter(port uint16) string {
////	filter := fmt.Sprintf("tcp and ((src port %v) or (dst port %v))",  port, port)
////	return filter
////}
//
//func main() {
//	port := uint16(18090)
//	// Open device
//	handle, err := pcap.OpenLive("\\Device\\NPF_{13478D8B-663A-47FC-9075-57D9DB23FEA0}",  int32(65535) ,true, pcap.BlockForever)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	defer handle.Close()
//	filter := getFilter(port)
//	err = handle.SetBPFFilter(filter)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	fmt.Println(handle.Stats())
//	fmt.Println("仅捕获TCP端口8090包")
//
//	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
//	for packet := range packetSource.Packets() {
//		printPacketInfo(packet)
//	}
//}
//
//func printPacketInfo(packet gopacket.Packet) {
//	// Let’s see if the packet is an ethernet packet
//	ethernetLayer := packet.Layer(layers.LayerTypeEthernet)
//	if ethernetLayer != nil {
//		fmt.Println("Ethernet layer detected.")
//		ethernetPacket, _ := ethernetLayer.(*layers.Ethernet)
//		fmt.Println("Source MAC: ", ethernetPacket.SrcMAC)
//		fmt.Println("Destination MAC: ", ethernetPacket.DstMAC)
//		// Ethernet type is typically IPv4 but could be ARP or other
//		fmt.Println("Ethernet type: ", ethernetPacket.EthernetType)
//		fmt.Println()
//	}
//
//	// Let’s see if the packet is IP (even though the ether type told us)
//	ipLayer := packet.Layer(layers.LayerTypeIPv4)
//	if ipLayer != nil {
//		fmt.Println("IPv4 layer detected.")
//		ip, _ := ipLayer.(*layers.IPv4)
//		fmt.Printf("From #{ip.SrcIP} to #{ip.DstIP} \n" )
//		fmt.Println("Protocol: ", ip.Protocol)
//		fmt.Println()
//	}
//
//	// Let’s see if the packet is TCP
//	tcpLayer := packet.Layer(layers.LayerTypeTCP)
//	if tcpLayer != nil {
//		fmt.Println("TCP layer detected.")
//		tcp, _ := tcpLayer.(*layers.TCP)
//
//		// TCP layer variables:
//		// SrcPort, DstPort, Seq, Ack, DataOffset, Window, Checksum, Urgent
//		// Bool flags: FIN, SYN, RST, PSH, ACK, URG, ECE, CWR, NS
//		fmt.Printf("From port %d to %d\n", tcp.SrcPort, tcp.DstPort)
//		fmt.Println("Sequence number: ", tcp.Seq)
//		fmt.Println()
//	}
//
//	// Iterate over all layers, printing out each layer type
//	fmt.Println("All packet layers:")
//	for _, layer := range packet.Layers() {
//		fmt.Println("- ", layer.LayerType())
//	}
//
//	// When iterating through packet.Layers() above,
//	// if it lists Payload layer then that is the same as
//	// this applicationLayer. applicationLayer contains the payload
//	applicationLayer := packet.ApplicationLayer()
//	if applicationLayer != nil {
//		fmt.Println("Application layer/Payload found.")
//		fmt.Printf("%s\n", applicationLayer.Payload())
//
//		// Search for a string inside the payload
//		if strings.Contains(string(applicationLayer.Payload()), "HTTP") {
//		fmt.Println("HTTP found!")
//		}
//	}
//
//	// Check for errors
//	if err := packet.ErrorLayer(); err != nil {
//		fmt.Println("Error decoding some part of the packet:", err)
//	}
//}
