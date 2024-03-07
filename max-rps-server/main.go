package main

import (
	"context"
	"io"
	"log"
	"log/slog"
	"net"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"runtime"
	"runtime/pprof"
	"strings"
	"sync"
	"syscall"
	"time"
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
	localIP()

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

	mux := http.NewServeMux()
	mux.HandleFunc("/", p.handleHome)

	server := http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	go func() {
		<-ctx.Done()
		slog.Info("Shutting down server...")
		err := server.Shutdown(ctx) // gracefully
		if err != nil {
			slog.Error("Server", "Shutdown", err)
		}
	}()

	err := server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		slog.Error("Server", "closed", err)
		cancel()
	}
}

type app struct {
	sync.RWMutex
	count   int
	elapsed time.Duration
}

func (p *app) handleHome(w http.ResponseWriter, r *http.Request) {
	if info {
		defer p.update(time.Now())
	}

	echo := r.URL.Query().Get("q")
	if echo != "" {
		io.WriteString(w, echo)
	}
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

func localIP() {
	hostname()
	hostname("-i")

	a, err := net.InterfaceAddrs()
	if err != nil {
		slog.Error("localIP", "net.InterfaceAddrs", err)
		return
	}

	for _, face := range a {
		if net, ok := face.(*net.IPNet); ok && !net.IP.IsLoopback() {
			slog.Info("localIP", "IP", net.IP.String())
		}
	}
}

func hostname(arg ...string) {
	cmd := exec.Command("hostname", arg...)
	output, err := cmd.Output()
	if err != nil {
		slog.Error("hostname", "Command", err)
		return
	}
	name := strings.TrimSpace(string(output))
	slog.Info("hostname", "output", name)
}
