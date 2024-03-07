package main

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"sync"
	"syscall"
	"time"
)

func main() {
	slog.SetDefault(logger)
	slog.Info(Name, "Version", Version)
	slog.Info(Name, "GitCommit", GitCommit)
	slog.Info(Name, "runtime.Version", runtime.Version())

	signals := make(chan os.Signal, 1) // Setup graceful shutdown
	// Docker and Kubernetes use the SIGTERM signal to gracefully shut down a container.
	signal.Notify(signals, os.Interrupt, syscall.SIGTERM)

	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	withServer := len(os.Args) == 1
	if withServer {
		wg.Add(1)
		go serve(&wg)
	}

	wg.Add(1)
	go retry(ctx, &wg)

	// Wait for the termination signal
	slog.Info("Starting graceful shutdown", "termination signal received", <-signals)
	if withServer {
		err := server.Shutdown(ctx)
		if err != nil {
			slog.Error("server", "error", err)
		}
	}

	cancel()

	wg.Wait()
	slog.Info("App: graceful shutdown completed")
}

const (
	firstInterval = 1 * time.Second
	retryInterval = 5 * time.Minute
)

func retry(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	t := time.NewTicker(firstInterval)
	defer t.Stop()
	for {
		select {
		case <-t.C:
			t.Reset(retryInterval)
			k, err := doYourJob(ctx)
			if err != nil {
				slog.Error("get", "error", err.Error())
				break
			}
			slog.Info("get", "key", k)

		case <-ctx.Done():
			slog.Info("getPeriodically: graceful shutdown completed")
			return
		}
	}
}

func doYourJob(ctx context.Context) (key string, err error) {
	get, err := http.NewRequestWithContext(ctx, http.MethodGet, address, nil)
	if err != nil {
		return
	}

	response, err := client.Do(get)
	if err != nil {
		return
	}
	defer response.Body.Close()

	var KeyId *struct{ Key string }
	err = json.NewDecoder(response.Body).Decode(&KeyId)
	if err != nil {
		return
	}

	key = KeyId.Key
	return
}

func serve(wg *sync.WaitGroup) {
	defer wg.Done()

	mux := http.NewServeMux()
	mux.HandleFunc("/logs/level", logsLevel)

	server = &http.Server{
		Addr:         serverAddr,
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	slog.Info("Server", "starting at", server.Addr)
	err := server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		slog.Error("Server", "error", err)
		os.Exit(1)
	}
	slog.Info("Server: graceful shutdown completed")
}

var (
	lvl    = new(slog.LevelVar)
	logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{AddSource: true, Level: lvl}))
)

func logsLevel(w http.ResponseWriter, r *http.Request) {
	slog.Info("home", "ClientIP", clientIP(r))

	type response struct {
		Status int
		Level  slog.Level
	}

	switch r.Method {
	case http.MethodGet:
		json.NewEncoder(w).Encode(response{Status: http.StatusOK, Level: lvl.Level()})

	case http.MethodPut:
		var v *response
		err := json.NewDecoder(r.Body).Decode(&v)
		if err != nil {
			httpError(w, http.StatusBadRequest)
			return
		}
		lvl.Set(v.Level)

	default:
		httpError(w, http.StatusMethodNotAllowed)
	}
}

func httpError(w http.ResponseWriter, code int) {
	http.Error(w, http.StatusText(code), code)
}

func clientIP(r *http.Request) string {
	ip := r.Header.Get("X-Real-IP")
	if ip == "" {
		ip = r.Header.Get("X-Forwarded-For")
		if ip == "" {
			ip = r.RemoteAddr
		}
	}
	return ip
}

var (
	address    = "http://127.0.0.1:8080/levels"
	serverAddr = ":8080"
	client     = http.Client{
		Timeout: 10 * time.Second,
	}
	server *http.Server

	Version   = "1.0.0"
	Name      = "MyApp"
	GitCommit = "git rev-parse HEAD"
	// CGO_ENABLED=0 go build -trimpath=true -ldflags="-s -X main.Version=$(Version) -X main.Name=$(Name) -X main.GitCommit=$(GIT_COMMIT)"
)
