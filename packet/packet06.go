package main

/**
这个例子做了几件事情。首先将显示如何使用网络设备发送原始字节。
这样就可以像串行连接一样使用它来发送数据。这对于真正的低层数据传输非常有用，
但如果您想与应用程序进行交互，您应该构建可以识别该数据包的其他硬件和软件。
接下来，它将显示如何使用以太网，IP和TCP层创建一个数据包。一切都是默认空的。
要完成它，我们创建另一个数据包，但实际上填写了以太网层的一些MAC地址，
IPv4的一些IP地址和TCP层的端口号。你应该看到如何伪装数据包和仿冒网络设备。
TCP层结构体具有可读取和可设置的SYN，FIN，ACK标志。
这有助于操纵和模糊TCP三次握手，会话和端口扫描。
pcap库提供了一种发送字节的简单方法，
但gopacket中的图层可帮助我们为多层创建字节结构。
*/

import (
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"log"
	"net"
	"time"
)

func main() {
	var (
		buffer  gopacket.SerializeBuffer
		options gopacket.SerializeOptions
	)
	// Open device
	handle, err := pcap.OpenLive("\\Device\\NPF_{13478D8B-663A-47FC-9075-57D9DB23FEA0}", 1024, false, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()

	// Send raw bytes over wire
	rawBytes := []byte{10, 20, 30}
	err = handle.WritePacketData(rawBytes)
	if err != nil {
		log.Fatal(err)
	}

	// Create a properly formed packet, just with
	// empty details. Should fill out MAC addresses,
	// IP addresses, etc.
	buffer = gopacket.NewSerializeBuffer()
	gopacket.SerializeLayers(buffer, options,
		&layers.Ethernet{},
		&layers.IPv4{},
		&layers.TCP{},
		gopacket.Payload(rawBytes),
	)
	outgoingPacket := buffer.Bytes()
	// Send our packet
	err = handle.WritePacketData(outgoingPacket)
	if err != nil {
		log.Fatal(err)
	}

	// This time lets fill out some information
	ipLayer := &layers.IPv4{
		SrcIP: net.IP{127, 0, 0, 1},
		DstIP: net.IP{8, 8, 8, 8},
	}
	ethernetLayer := &layers.Ethernet{
		SrcMAC: net.HardwareAddr{0xFF, 0xAA, 0xFA, 0xAA, 0xFF, 0xAA},
		DstMAC: net.HardwareAddr{0xBD, 0xBD, 0xBD, 0xBD, 0xBD, 0xBD},
	}
	tcpLayer := &layers.TCP{
		SrcPort: layers.TCPPort(4321),
		DstPort: layers.TCPPort(80),
	}
	// And create the packet with the layers
	buffer = gopacket.NewSerializeBuffer()
	gopacket.SerializeLayers(buffer, options,
		ethernetLayer,
		ipLayer,
		tcpLayer,
		gopacket.Payload(rawBytes),
	)
	outgoingPacket = buffer.Bytes()
}
