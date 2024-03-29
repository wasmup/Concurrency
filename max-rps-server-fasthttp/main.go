package main

import (
	"context"
	"encoding/base64"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"runtime/pprof"
	"sync"
	"syscall"
	"time"

	"github.com/valyala/fasthttp"
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

	var ap app

	wg.Add(1)
	go ap.serve(ctx, cancel, &wg)

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

func (p *app) serve(ctx context.Context, cancel context.CancelFunc, wg *sync.WaitGroup) {
	defer wg.Done()

	server := &fasthttp.Server{
		Handler: p.handleHome,
	}

	go func() {
		<-ctx.Done()
		slog.Info("Shutting down server...")
		err := server.ShutdownWithContext(ctx) // gracefully
		if err != nil {
			slog.Error("Server", "Shutdown", err)
		}
	}()

	err := server.ListenAndServe(":8080")
	if err != nil && err != http.ErrServerClosed {
		slog.Error("Server", "closed", err)
		cancel()
	}
}

func (p *app) handleHome(ctx *fasthttp.RequestCtx) {
	if info {
		defer p.update(time.Now())
	}

	q := string(ctx.QueryArgs().Peek("q"))
	if q == "" {
		q = "Hi"
	}
	ctx.Write([]byte(base64.StdEncoding.EncodeToString([]byte(q))))
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
