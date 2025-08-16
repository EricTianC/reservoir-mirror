package main

import (
	"flag"
	"log/slog"
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

	slog.Info("Launching reservoir mirror server.")
	syncReservoirIndex()
}
