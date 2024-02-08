package main

import (
	"log/slog"
	"os"
	"runtime"
)

var Version, Name, GitCommit string

func main() {
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{AddSource: true})))
	slog.Info(Name, "runtime.Version", runtime.Version())
	slog.Info(Name, "Version", Version)
	slog.Info(Name, "GitCommit", GitCommit)
}
