package main

import (
	"bytes"
	"context"
	"encoding/base64"
	"flag"
	"fmt"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"runtime"
	"runtime/pprof"
	"sync"
	"syscall"
	"time"

	"github.com/evanphx/wildcat"
	"github.com/panjf2000/gnet/v2"
)

const info = true

func main() {
	{
		// Create the CPU profile file, overwrite if exists.
		// go tool pprof -http=":8787" cpu.out
		f, err := os.Create("cpu.out")
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		// Start CPU profiling.
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal(err)
		}
		defer pprof.StopCPUProfile()
	}

	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{AddSource: true})))
	slog.Info(runtime.Version())

	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var (
		port      int
		multicore bool
	)
	flag.IntVar(&port, "port", 8080, "server port")
	flag.BoolVar(&multicore, "multicore", true, "multicore")
	flag.Parse()

	var ap app

	wg.Add(1)
	go ap.serve(ctx, cancel, &wg, port, multicore)

	if info {
		wg.Add(1)
		go ap.printStats(ctx, &wg)
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	select {
	case sig := <-stop:
		slog.Info("terminated", "signal", sig)
		cancel()

	case <-ctx.Done():
		slog.Info("terminated by context")
	}

	wg.Wait()
	slog.Info("App closed")
}

func (p *app) serve(ctx context.Context, cancel context.CancelFunc, wg *sync.WaitGroup, port int, multicore bool) {
	defer wg.Done()

	server := &httpServer{
		addr:      fmt.Sprintf("tcp://127.0.0.1:%d", port),
		multicore: multicore,
		update:    p.update,
	}

	go func() {
		<-ctx.Done()
		slog.Info("Shutting down server...")
		err := server.eng.Stop(ctx) // gracefully
		if err != nil {
			slog.Error("Server", "Shutdown", err)
		}
	}()

	err := gnet.Run(server, server.addr, gnet.WithMulticore(multicore))
	if err != nil {
		slog.Error("Server", "closed", err)
		cancel()
	}
}

type httpCodec struct {
	parser *wildcat.HTTPParser
	buf    []byte
}

func (p *httpCodec) appendResponse(queryParam string) {
	p.buf = append(p.buf, "HTTP/1.1 200 OK\r\nServer: gnet\r\nContent-Type: text/plain\r\nDate: "...)
	p.buf = time.Now().AppendFormat(p.buf, "Mon, 02 Jan 2006 15:04:05 GMT")

	base64Param := base64.StdEncoding.EncodeToString([]byte(queryParam))
	p.buf = append(p.buf, fmt.Sprintf("\r\nContent-Length: %d\r\n\r\n%s", len(base64Param), base64Param)...)
}

type httpServer struct {
	gnet.BuiltinEventEngine

	addr      string
	multicore bool
	eng       gnet.Engine
	update    func(time.Time)
}

func (p *httpServer) OnBoot(eng gnet.Engine) gnet.Action {
	p.eng = eng
	log.Printf("echo server with multi-core=%t is listening on %s\n", p.multicore, p.addr)
	return gnet.None
}

func (p *httpServer) OnOpen(c gnet.Conn) ([]byte, gnet.Action) {
	c.SetContext(&httpCodec{parser: wildcat.NewHTTPParser()})
	return nil, gnet.None
}

func (p *httpServer) OnTraffic(c gnet.Conn) gnet.Action {
	if info {
		defer p.update(time.Now())
	}

	hc := c.Context().(*httpCodec)
	buf, _ := c.Next(-1)

	for {
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
		if len(buf) == 0 {
			break
		}
	}

	c.Write(hc.buf)
	hc.buf = hc.buf[:0]
	return gnet.None
}

type app struct {
	sync.RWMutex
	count   int
	elapsed time.Duration
}

func (p *app) update(t0 time.Time) {
	d := time.Since(t0)
	p.Lock()
	p.count++
	p.elapsed += d
	p.Unlock()
}

func (p *app) printStats(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	tick := time.NewTicker(10 * time.Second)
	defer tick.Stop()

	for {
		select {
		case <-tick.C:
			p.RLock()
			n := p.count
			d := p.elapsed
			p.RUnlock()
			slog.Info("Stats", "Requests", n, "Elapsed", d)

		case <-ctx.Done():
			return
		}
	}
}

var (
	errMsg      = "Internal Server Error"
	errMsgBytes = []byte(errMsg)
)
