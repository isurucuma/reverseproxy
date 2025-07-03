package main

import (
	"flag"
	"fmt"
	"github.com/isurucuma/reverseproxy/internal/proxy"
	"log/slog"
)

func main() {
	listenAddr := flag.String("listen", ":8080", "listen address")
	backendAddr := flag.String("backend", ":8081", "forward address")
	flag.Parse()

	slog.Info(fmt.Sprintf("starting proxy server on %s, forwarding to %s", *listenAddr, *backendAddr))

	if err := proxy.Run(*listenAddr, *backendAddr); err != nil {
		slog.Error("failed to start proxy server", "error", err)
		return
	}
}
