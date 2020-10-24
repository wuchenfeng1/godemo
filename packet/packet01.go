package main

import (
	"fmt"
	"github.com/google/gopacket/pcap"
	"log"
	"net"
	"strings"
)

type FaceInfo struct {
	NPFName     string
	Description string
	NickName    string
	IPv4        string
}

//获取window系统上设备地址  锁定抓取源
func getIfList() []FaceInfo {
	var ifaceInfoList []FaceInfo

	// 得到所有的(网络)设备
	devices, err := pcap.FindAllDevs()
	if err != nil {
		log.Fatal(err)
	}

	interfaceList, err := net.Interfaces()
	if err != nil {
		log.Fatal(err)
	}

	for _, i := range interfaceList {
		byName, err := net.InterfaceByName(i.Name)
		if err != nil {
			log.Fatal(err)
		}
		address, err := byName.Addrs()
		ifaceInfoList = append(ifaceInfoList, FaceInfo{NickName: byName.Name, IPv4: address[1].String()})
	}
	var validInterfaces []FaceInfo
	for _, device := range devices {
		for _, address := range device.Addresses {
			for _, inaction := range ifaceInfoList {
				if strings.Contains(inaction.IPv4, address.IP.String()) {
					validInterfaces = append(validInterfaces, FaceInfo{NPFName: device.Name, Description: device.Description, NickName: inaction.NickName, IPv4: inaction.IPv4})
					break
				}
			}
		}
	}

	return validInterfaces
}

func main() {

	list := getIfList()

	for _, item := range list {
		fmt.Println(item)
	}

}
