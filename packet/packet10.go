package main

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"github.com/google/gopacket/tcpassembly"
	"github.com/google/gopacket/tcpassembly/tcpreader"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

type httpStreamFactory struct{}

type httpStream struct {
	net, transport gopacket.Flow
	r              tcpreader.ReaderStream
}

func (h *httpStream) runResponse() {

	buf := bufio.NewReader(&h.r)
	defer tcpreader.DiscardBytesToEOF(buf)
	for {
		resp, err := http.ReadResponse(buf, nil)
		if err == io.EOF || err == io.ErrUnexpectedEOF {
			return
		} else if err != nil {
			log.Println("读取流时出错：", h.net, h.transport, ":", err)
			return
		} else {
			printResponse(resp, h)
		}
	}
}
func (h *httpStream) runRequest() {

	buf := bufio.NewReader(&h.r)
	defer tcpreader.DiscardBytesToEOF(buf)
	for {
		req, err := http.ReadRequest(buf)
		if err == io.EOF || err == io.ErrUnexpectedEOF {
			// We must read until we see an EOF... very important!
			return
		} else if err != nil {
			log.Println("读取流时出错：", h.net, h.transport, ":", err)
		} else {
			printRequest(req, h)
		}
	}
}

//打印请求数据
func printRequest(req *http.Request, req2 *httpStream) {
	fmt.Println("请求:")
	fmt.Println(req2.net, req2.transport)
	fmt.Println("\n\r")
	fmt.Println("请求方法:", req.Method, "请求uri:", req.RequestURI, "请求协议:", req.Proto)
	printHeader(req.Header)
	defer req.Body.Close()
	fmt.Println("请求body：", req.Body)
}

// 打印响应
func printResponse(resp *http.Response, h *httpStream) {

	fmt.Println("响应协议：", resp.Proto, "响应状态：", resp.Status)
	printHeader(resp.Header)
	defer resp.Body.Close()
	body := resp.Body
	if strings.Contains(resp.Header.Get("Content-Encoding"), "gzip") {
		unbody, err := gzip.NewReader(resp.Body)
		if err != nil {
			fmt.Println("http resp unzip is failed,err: ", err)
		} else {
			body = unbody
		}
	}
	data, _ := ioutil.ReadAll(body)
	fmt.Println("响应的数据是：", string(data))
}

// 打印头部
func printHeader(h http.Header) {
	for k, v := range h {
		fmt.Println("打印头部的键值对：", k, v)
	}
}
func (h httpStreamFactory) New(netFlow, tcpFlow gopacket.Flow) tcpassembly.Stream {
	hstream := &httpStream{
		net:       netFlow,
		transport: tcpFlow,
		r:         tcpreader.NewReaderStream(),
	}

	src, _ := tcpFlow.Endpoints()
	if fmt.Sprintf("%v", src) == "18090" {
		go hstream.runResponse()
	} else {
		go hstream.runRequest()
	}

	// ReaderStream implements tcpassembly.Stream, so we can return a pointer to it.
	return &hstream.r
}

//定义过滤器
func getFilter(port uint16) string {
	filter := fmt.Sprintf("tcp and  port %v", port)
	return filter
}

func main() {
	handle, err := pcap.OpenLive("\\Device\\NPF_{13478D8B-663A-47FC-9075-57D9DB23FEA0}", 1600, false, pcap.BlockForever)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer handle.Close()
	filter := getFilter(18090)
	err = handle.SetBPFFilter(filter)
	if err != nil {
		log.Fatal("监听端口错误：", err)
	}
	fmt.Println("仅捕获TCP端口18090包")
	// Set up assembly
	streamFactory := &httpStreamFactory{}
	streamPool := tcpassembly.NewStreamPool(streamFactory)
	assembler := tcpassembly.NewAssembler(streamPool)

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	packets := packetSource.Packets()
	ticker := time.Tick(time.Minute)
	for {
		select {
		case packet := <-packets:
			// 零数据包表示pcap文件的结尾.
			if packet == nil {
				return
			}
			if packet.NetworkLayer() == nil || packet.TransportLayer() == nil || packet.TransportLayer().LayerType() != layers.LayerTypeTCP {
				log.Println("Unusable packet")
				continue
			}
			tcp := packet.TransportLayer().(*layers.TCP)
			assembler.AssembleWithTimestamp(packet.NetworkLayer().NetworkFlow(), tcp, packet.Metadata().Timestamp)
		case <-ticker:
			// 每分钟刷新过去2分钟内没有活动的连接。
			assembler.FlushOlderThan(time.Now().Add(time.Minute * -2))
		}
	}
}
