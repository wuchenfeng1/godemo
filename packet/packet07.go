package main

//.更多的解码/构造数据包的例子
import (
	"fmt"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

func main() {
	//如果我们没有设备或文件的句柄，但我们有一堆
	//对于原始字节，我们可以尝试将它们解码成包信息
	//NewPacket（）将组成数据包的原始字节作为第一个参数
	//第二个参数是要解码的最低层。会的
	//解码该层及其上的所有层。第三层
	//是解码的类型：default（一次完成）、lazy（按需）和NoCopy
	//不会创建缓冲区的副本
	//创建具有以太网、IP、TCP和负载层的数据包
	//我们正在创造一个我们知道会被正确解码，但是
	//你的字节源可以是任何东西。如果有任何包
	//返回为零，这意味着它无法将其解码为
	//正确层（格式错误或不正确的数据包类型）
	payload := []byte{2, 4, 6}
	options := gopacket.SerializeOptions{}
	buffer := gopacket.NewSerializeBuffer()
	gopacket.SerializeLayers(buffer, options,
		&layers.Ethernet{},
		&layers.IPv4{},
		&layers.TCP{},
		gopacket.Payload(payload),
	)
	rawBytes := buffer.Bytes()

	// Decode an ethernet packet
	ethPacket :=
		gopacket.NewPacket(
			rawBytes,
			layers.LayerTypeEthernet,
			gopacket.Default,
		)

	//通过懒惰解码，它只会在需要的时候解码它需要的东西
	//这不是并发安全的。如果使用并发，则使用默认值
	ipPacket :=
		gopacket.NewPacket(
			rawBytes,
			layers.LayerTypeIPv4,
			gopacket.Lazy,
		)

	//使用NoCopy选项，将引用底层切片
	//直接不复制。如果底层字节发生变化，那么
	//包裹
	tcpPacket :=
		gopacket.NewPacket(
			rawBytes,
			layers.LayerTypeTCP,
			gopacket.NoCopy,
		)

	fmt.Println(ethPacket)
	fmt.Println(ipPacket)
	fmt.Println(tcpPacket)
}
