// Copyright 2012 Google, Inc. All rights reserved.
//
// Use of this source code is governed by a BSD-style license
// that can be found in the LICENSE file in the root of the source
// tree.

// This binary provides sample code for using the gopacket TCP assembler and TCP
// stream reader.  It reads packets off the wire and reconstructs HTTP requests
// it sees, logging them.
package main

import (
	"bufio"
	"compress/gzip"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/examples/util"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"github.com/google/gopacket/tcpassembly"
	"github.com/google/gopacket/tcpassembly/tcpreader"
)

//\\Device\\NPF_{93699EAA-B3C9-4D00-9445-73C7C361F433}
var iface = flag.String("i", "\\Device\\NPF_{13478D8B-663A-47FC-9075-57D9DB23FEA0}", "Interface to get packets from")
var fname = flag.String("r", "", "Filename to read from, overrides -i")
var snaplen = flag.Int("s", 1600, "SnapLen for pcap packet capture")
var filter = flag.String("f", "tcp and port 18090", "BPF filter for pcap")
var logAllPackets = flag.Bool("v", false, "Logs every packet in great detail")

// Build a simple HTTP request parser using tcpassembly.StreamFactory and tcpassembly.Stream interfaces

// httpStreamFactory implements tcpassembly.StreamFactory
type httpStreamFactory struct{}

// httpStream will handle the actual decoding of http requests.
type httpStream struct {
	net, transport gopacket.Flow
	r              tcpreader.ReaderStream
}

func (h *httpStreamFactory) New(net, transport gopacket.Flow) tcpassembly.Stream {
	hstream := &httpStream{
		net:       net,
		transport: transport,
		r:         tcpreader.NewReaderStream(),
	}
	src, _ := transport.Endpoints()
	log.Println(fmt.Sprintf("%v", src))
	if fmt.Sprintf("%v", src) == "18090" {
		go hstream.runResponse() // Important... we must guarantee that data from the reader stream is read.
	} else {
		go hstream.runRequest() // Important... we must guarantee that data from the reader stream is read.
	}
	// ReaderStream implements tcpassembly.Stream, so we can return a pointer to it.
	return &hstream.r
}

func (h *httpStream) runResponse() {
	buf := bufio.NewReader(&h.r)
	defer tcpreader.DiscardBytesToEOF(buf)
	for {
		resp, err := http.ReadResponse(buf, nil)
		if err == io.EOF || err == io.ErrUnexpectedEOF {
			log.Println("我们必须阅读直到我们看到一个EOF。。。非常重要！")
			return
		} else if err != nil {
			log.Println("Error reading stream", h.net, h.transport, ":", err)
			return
		} else {
			printResponse(resp, h)
			// log.Println("Received response from stream", h.net, h.transport, ":", resp, "with", bodyBytes, "bytes in response body")
		}
	}
}
func (h *httpStream) runRequest() {

	buf := bufio.NewReader(&h.r)
	defer tcpreader.DiscardBytesToEOF(buf)
	for {
		req, err := http.ReadRequest(buf)
		if err == io.EOF || err == io.ErrUnexpectedEOF {
			log.Println("我们必须阅读直到我们看到一个EOF。。。非常重要！", err)
			return
		} else if err != nil {
			log.Println("Error reading stream", h.net, h.transport, ":", err)
		} else {
			printRequest(req, h)
			// log.Println("Received request from stream", h.net, h.transport, ":", req, "with", bodyBytes, "bytes in request body")
		}
	}
}

// 打印头部
func printHeader(h http.Header) {
	for k, v := range h {
		fmt.Println(k, v)
	}
}

// 打印请求
func printRequest(req *http.Request, h *httpStream) {

	fmt.Println("\n\rRequest -->")
	fmt.Println("\n\r")
	fmt.Println(h.net, h.transport)
	fmt.Println("\n\r")
	fmt.Println(req.Method, req.RequestURI, req.Proto)
	printHeader(req.Header)
	defer req.Body.Close()
	//body, _ := ioutil.ReadAll(req.Body)
	//fmt.Println(string(body))
	fmt.Println(req.Body)
}

// 打印响应
func printResponse(resp *http.Response, h *httpStream) {

	fmt.Println("\n\rResponse -->")
	fmt.Println("\n\r")
	fmt.Println(resp.Proto, resp.Status)
	//printHeader(resp.Header)
	defer resp.Body.Close()
	//body, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println(string(body))
	//fmt.Println(resp.Body)
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
	fmt.Println(string(data))
}

func main() {
	defer util.Run()()
	var handle *pcap.Handle
	var err error

	// Set up pcap packet capture
	if *fname != "" {
		log.Printf("Reading from pcap dump %q", *fname)
		handle, err = pcap.OpenOffline(*fname)
	} else {
		log.Printf("Starting capture on interface %q", *iface)
		handle, err = pcap.OpenLive(*iface, int32(*snaplen), true, pcap.BlockForever)
	}
	if err != nil {
		log.Fatal(err)
	}

	if err := handle.SetBPFFilter(*filter); err != nil {
		log.Fatal(err)
	}

	// Set up assembly
	streamFactory := &httpStreamFactory{}
	streamPool := tcpassembly.NewStreamPool(streamFactory)
	assembler := tcpassembly.NewAssembler(streamPool)

	log.Println("reading in packets")
	// Read in packets, pass to assembler.
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	packets := packetSource.Packets()
	ticker := time.Tick(time.Minute)
	for {
		select {
		case packet := <-packets:
			// A nil packet indicates the end of a pcap file.
			if packet == nil {
				return
			}
			if *logAllPackets {
				log.Println(packet)
			}
			if packet.NetworkLayer() == nil || packet.TransportLayer() == nil || packet.TransportLayer().LayerType() != layers.LayerTypeTCP {
				log.Println("Unusable packet")
				continue
			}
			tcp := packet.TransportLayer().(*layers.TCP)
			assembler.AssembleWithTimestamp(packet.NetworkLayer().NetworkFlow(), tcp, packet.Metadata().Timestamp)
		case <-ticker:
			// Every minute, flush connections that haven't seen activity in the past 2 minutes.
			assembler.FlushOlderThan(time.Now().Add(time.Minute * -2))
		}
	}
}
