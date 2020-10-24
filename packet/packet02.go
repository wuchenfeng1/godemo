package main

import (
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"github.com/google/gopacket/pcapgo"
	"log"
	"os"
	"time"
)

//锁定源抓取数据包保存到test.pcap文件中
func main() {
	var (
		snapshotLen int32 = 1024
		packetCount int   = 0
	)
	handle, err := pcap.OpenLive("\\Device\\NPF_{13478D8B-663A-47FC-9075-57D9DB23FEA0}", 1024, false, 30*time.Second)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	f, _ := os.Create("test.pcap") // 如果文件已存在，会将文件清空。
	w := pcapgo.NewWriter(f)
	w.WriteFileHeader(uint32(snapshotLen), layers.LinkTypeEthernet)
	defer f.Close()
	defer handle.Close()

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		log.Println("程序metadata：", packet.Metadata())
		log.Println("程序Data：", packet.Data())
		w.WritePacket(packet.Metadata().CaptureInfo, packet.Data())
		packetCount++
		// Only capture 100 and then stop
		if packetCount > 100 {
			break
		}
	}
}
