package main

import (
	"bytes"
	"encoding/base64"
	"flag"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/evanphx/wildcat"
	"github.com/panjf2000/gnet/v2"
)

var (
	errMsg      = "Internal Server Error"
	errMsgBytes = []byte(errMsg)
)

type httpServer struct {
	gnet.BuiltinEventEngine

	addr      string
	multicore bool
	eng       gnet.Engine
}

type httpCodec struct {
	parser *wildcat.HTTPParser
	buf    []byte
}

func (hc *httpCodec) appendResponse(queryParam string) {
	hc.buf = append(hc.buf, "HTTP/1.1 200 OK\r\nServer: gnet\r\nContent-Type: text/plain\r\nDate: "...)
	hc.buf = time.Now().AppendFormat(hc.buf, "Mon, 02 Jan 2006 15:04:05 GMT")

	base64Param := base64.StdEncoding.EncodeToString([]byte(queryParam))
	hc.buf = append(hc.buf, fmt.Sprintf("\r\nContent-Length: %d\r\n\r\n%s", len(base64Param), base64Param)...)

}

func (hs *httpServer) OnBoot(eng gnet.Engine) gnet.Action {
	hs.eng = eng
	log.Printf("echo server with multi-core=%t is listening on %s\n", hs.multicore, hs.addr)
	return gnet.None
}

func (hs *httpServer) OnOpen(c gnet.Conn) ([]byte, gnet.Action) {
	c.SetContext(&httpCodec{parser: wildcat.NewHTTPParser()})
	return nil, gnet.None
}

func (hs *httpServer) OnTraffic(c gnet.Conn) gnet.Action {
	t0 := time.Now()
	defer func() {
		d := time.Since(t0)
		mu.Lock()
		requestCount++
		elapsedTime += d
		mu.Unlock()
	}()

	hc := c.Context().(*httpCodec)
	buf, _ := c.Next(-1)

pipeline:
	headerOffset, err := hc.parser.Parse(buf)
	if err != nil {
		c.Write(errMsgBytes)
		return gnet.Close
	}
	queryParam := []byte("Hi")
	q := []byte("/?q=")
	if bytes.HasPrefix(hc.parser.Path, q) {
		queryParam = hc.parser.Path[4:]
	}
	hc.appendResponse(base64.RawStdEncoding.EncodeToString(queryParam))

	bodyLen := int(hc.parser.ContentLength())
	if bodyLen == -1 {
		bodyLen = 0
	}

	buf = buf[headerOffset+bodyLen:]
	if len(buf) > 0 {
		goto pipeline
	}

	c.Write(hc.buf)
	hc.buf = hc.buf[:0]
	return gnet.None
}

func main() {
	go printStats()
	var port int
	var multicore bool

	flag.IntVar(&port, "port", 8080, "server port")
	flag.BoolVar(&multicore, "multicore", true, "multicore")
	flag.Parse()

	hs := &httpServer{addr: fmt.Sprintf("tcp://127.0.0.1:%d", port), multicore: multicore}

	log.Println("server exits:", gnet.Run(hs, hs.addr, gnet.WithMulticore(multicore)))
}

func printStats() {
	for {
		time.Sleep(time.Second)
		mu.Lock()
		n := requestCount
		d := elapsedTime
		mu.Unlock()
		fmt.Printf("Requests: %d, Total Elapsed Time: %s\n", n, d)
	}
}

var (
	requestCount int
	elapsedTime  time.Duration
	mu           sync.Mutex
)
