package main

/*下一个程序将显示如何创建自己的layer。构建gopacket
layer 包不包含的协议。如果您要创建自己的l33t协议，
甚至不使用TCP / IP或以太网，这是很有用的。
*/

import (
	"fmt"
	"github.com/google/gopacket"
)

//创建自定义层结构
type CustomLayer struct {
	//这个层前面只有两个字节
	SomeByte    byte
	AnotherByte byte
	restOfData  []byte
}

//注册图层类型以便我们可以使用它
//第一个参数是一个ID。使用负数
//或2000+用于自定义图层。它一定是独一无二的
var CustomLayerType = gopacket.RegisterLayerType(
	2001,
	gopacket.LayerTypeMetadata{
		Name:    "CustomLayerType",
		Decoder: gopacket.DecodeFunc(decodeCustomLayer),
	},
)

//注册图层类型以便我们可以使用它
//第一个参数是一个ID。使用负数
//或2000+用于自定义图层。它必须是唯一的//当我们查询类型时，应该是什么类型的层
//我们说是吗？我们希望它返回我们的自定义图层类型
func (l CustomLayer) LayerType() gopacket.LayerType {
	return CustomLayerType
}

//LayerContents返回我们的层
//提供。在本例中，它是一个头层，因此
//我们返回标题信息
func (l CustomLayer) LayerContents() []byte {
	return []byte{l.SomeByte, l.AnotherByte}
}

//LayerPayload返回生成的后续层
//在我们的层或原始有效载荷之上
func (l CustomLayer) LayerPayload() []byte {
	return l.restOfData
}

//自定义解码功能。我们想叫什么名字都行
//但是它应该有相同的参数和返回值
//当层被注册时，我们告诉它使用这个解码功能
func decodeCustomLayer(data []byte, p gopacket.PacketBuilder) error {
	//AddLayer附加到包所具有的层的列表中
	p.AddLayer(&CustomLayer{data[0], data[1], data[2:]})
	//返回值告诉数据包期望哪个层
	//以及其他数据。可能是另一个头层，
	//什么也没有，或者是有效载荷层。
	//nil表示这是最后一层。不再解码了
	//返回零
	//返回另一层类型告诉它解码
	//具有该层解码器功能的下一层
	//返回p.NextDecoder(layers.layertypethernet层)
	//返回有效负载类型意味着剩余的数据
	//是原始有效载荷。它将设置应用层
	//有效载荷的内容物
	return p.NextDecoder(gopacket.LayerTypePayload)
}

func main() {
	//如果你创建自己的编码和解码，你基本上可以
	//创建自己的协议或实现一个不是
	//已经在layers包中定义。在我们的例子中，我们只是
	//用我们自己的层包装一个普通的以太网包。
	//创建自己的协议是很好的，如果你想创建
	//一些模糊的二进制数据类型对其他人来说很难
	//解码
	//最后，解码您的数据包：
	rawBytes := []byte{0xF0, 0x0F, 65, 65, 66, 67, 68}
	packet := gopacket.NewPacket(
		rawBytes,
		CustomLayerType,
		gopacket.Default,
	)
	fmt.Println("Created packet out of raw bytes.")
	fmt.Println(packet)

	// Decode the packet as our custom layer
	customLayer := packet.Layer(CustomLayerType)
	if customLayer != nil {
		fmt.Println("Packet was successfully decoded with custom layer decoder.")
		customLayerContent, _ := customLayer.(*CustomLayer)
		// Now we can access the elements of the custom struct
		fmt.Println("Payload: ", customLayerContent.LayerPayload())
		fmt.Println("SomeByte element:", customLayerContent.SomeByte)
		fmt.Println("AnotherByte element:", customLayerContent.AnotherByte)
	}
}
