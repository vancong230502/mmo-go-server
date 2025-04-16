package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/panjf2000/gnet/v2"
)

type tcpServer struct {
	gnet.BuiltinEventEngine
	eng       gnet.Engine
	addr      string
	multicore bool
}

func (ts *tcpServer) OnBoot(eng gnet.Engine) gnet.Action {
	ts.eng = eng
	log.Printf("TCP server is listening on %s (multi-cores: %t)\n", ts.addr, ts.multicore)
	return gnet.None
}

func (ts *tcpServer) OnTraffic(c gnet.Conn) gnet.Action {
	// Echo the received data
	buf, _ := c.Next(-1)
	c.Write(buf)
	return gnet.None
}

func main() {
	var port int
	var multicore bool

	// Parse command line flags
	flag.IntVar(&port, "port", 9000, "server port")
	flag.BoolVar(&multicore, "multicore", true, "multicore")
	flag.Parse()

	addr := fmt.Sprintf("tcp://:%d", port)
	ts := &tcpServer{addr: addr, multicore: multicore}

	// Start server
	log.Printf("Starting server on port %d...\n", port)
	err := gnet.Run(ts, addr,
		gnet.WithMulticore(multicore),
		gnet.WithReusePort(true),
		gnet.WithTCPKeepAlive(time.Minute*5),
	)
	log.Printf("Tesing log on port %d...\n", port)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
} 