package main

import (
	"context"
	"go.uber.org/zap"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/unifi-blueberry/core/internal/addon"
	"github.com/unifi-blueberry/core/internal/core"
)

func init() {
	var logger *zap.Logger
	logger, _ = zap.NewDevelopment()
	zap.ReplaceGlobals(logger)
}

func main() {
	mux := http.NewServeMux()

	core.RegisterServer(mux)
	addon.RegisterServer(mux)

	server := &http.Server{
		Addr: ":10234",
		// Use h2c so we can serve HTTP/2 without TLS.
		Handler: h2c.NewHandler(mux, &http2.Server{}),
	}

	start(server)
}

// start starts an http server and manages a graceful shutdown
func start(server *http.Server) {

	zap.S().Infof("starting server, listening on %s", server.Addr)

	// publish OS signals to a channel
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	// run server in a goroutine so we can block this thread on the signal channel
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			zap.S().Panicln(err)
		}
	}()

	// wait for an os signal
	<-signals
	zap.S().Infoln("initiating server shutdown")

	// create a context with a hard timeout for shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// initiate graceful shutdown
	if err := server.Shutdown(ctx); err != nil {
		zap.S().Panicln(err)
	}

	zap.S().Infoln("server has shutdown")
}
