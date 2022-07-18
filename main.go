package main

import (
	"context"
	"errors"
	"flag"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"urlservice/api"
	"urlservice/middleware"

	"github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
)

var (
	dev = flag.Bool("dev", false, "development mode")
)

func main() {
	flag.Parse()

	addr := osLookupEnv("LISTEN_ADDR", ":8081")
	tlsCertFile := osLookupEnv("TLS_CERT_FILE", "")
	tlsKeyFile := osLookupEnv("TLS_KEY_FILE", "")

	log := logrus.WithField("app", "url_service").WithField("dev", *dev)

	handler := middleware.WithMiddleware(api.New(log), log, *dev)

	server := &http.Server{Addr: addr, Handler: handler}

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	group, ctx := errgroup.WithContext(ctx)

	group.Go(func() error {
		var err error
		ll := log.WithField("addr", addr)
		if tlsCertFile != "" && tlsKeyFile != "" {
			ll.Info("starting server with TLS")
			err = server.ListenAndServeTLS(tlsCertFile, tlsKeyFile)
		} else {
			ll.Info("starting server without TLS")
			err = server.ListenAndServe()
		}
		if errors.Is(err, http.ErrServerClosed) {
			ll.Info("server stopped")
			return nil
		}
		return err
	})
	group.Go(func() error {
		signalChan := make(chan os.Signal, 1)
		signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
		select {
		case <-signalChan:
			log.Info("shutdown received")

			server.Shutdown(context.Background())
		case <-ctx.Done():
			return ctx.Err()
		}

		return nil
	})
	if err := group.Wait(); err != nil {
		log.WithError(err).Fatalln("application failed")
	}
	log.Info("application stopped")
}

func osLookupEnv(key, defaultValue string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return defaultValue
}
