package main

import (
	"context"
	"flag"
	"log/slog"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

var debugOpt bool
var proxyAddress string

func main() {
	flag.BoolVar(&debugOpt, "debug", false, "enable debug output")
	flag.StringVar(&proxyAddress, "proxy", "", "proxy url")

	flag.Parse()

	// proxyUrl, _ := http.ProxyFromEnvironment() TODO: 从环境中获取代理配置
	if proxyAddress != "" {
		slog.Info("Using proxy URL:", "proxyAddress", proxyAddress)
	}

	if debugOpt {
		slog.Info("debug mode enabled.")
		slog.SetLogLoggerLevel(slog.LevelDebug)
	}

	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup

	slog.Info("Launching reservoir mirror server.")
	wg.Add(1)
	go ReservoirIndexServer(ctx, &wg)

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-signalChan
	slog.Info("Shutting down...")
	cancel()
	wg.Wait()
}
