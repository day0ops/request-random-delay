package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"github.com/day0ops/request-random-delay/pkg/handlers"
	"github.com/day0ops/request-random-delay/pkg/logger"
	"go.uber.org/zap"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	port = flag.Int("port", 9001, "gRPC port")
)

func main() {
	os.Exit(start())
}

func start() int {
	l := logger.Get()

	flag.Parse()

	ctx, _ := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)

	h := handlers.NewHandler(handlers.LogWith(l))

	server := http.Server{
		Addr:    fmt.Sprintf(":%d", *port),
		Handler: h,
	}

	go func() {
		err := server.ListenAndServe()
		if err != nil && errors.Is(err, http.ErrServerClosed) {
			l.Info("server closed")
		} else {
			l.Fatal("failed to listen", zap.Error(err))
		}
	}()

	<-ctx.Done()
	l.Info("shutting down server")
	shoutDownCtx, shutdownRelease := context.WithTimeout(ctx, time.Second*10)
	defer shutdownRelease()
	if err := server.Shutdown(shoutDownCtx); err != nil {
		log.Fatal(err)
	}

	return 0
}
