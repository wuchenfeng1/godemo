package main

import (
	"fmt"
	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
	"log"
	"os"
	"time"
)

//默认是80端口
func main() {
	handle, err := pcap.OpenLive("\\Device\\NPF_{13478D8B-663A-47FC-9075-57D9DB23FEA0}", 1024, false, 30*time.Second)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer handle.Close()
	err = handle.SetBPFFilter("tcp and port 8090")
	if err != nil {
		log.Fatal("监听8090端口错误：", err)
	}
	fmt.Println("仅捕获TCP端口8090包")
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		log.Println("程序metadata：", packet.Metadata())
		log.Println("程序Data：", packet.Data())
	}
}
