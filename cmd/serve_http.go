package cmd

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/cikupin/feature-flag-example/handler"
	"github.com/urfave/cli/v2"
)

// ServeHTTP will serve http rest API
var ServeHTTP = &cli.Command{
	Name:        "http",
	Usage:       "serve http rest API",
	Description: "serve http rest API",
	Action: func(c *cli.Context) error {
		serveHTTP(c)
		return nil
	},
}

func serveHTTP(ctx context.Context) {
	h := http.NewServeMux()
	h.HandleFunc("/toggle-feature", handler.ToggleFeature)
	h.HandleFunc("/toggle-provider", handler.ToggleProvider)

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	s := &http.Server{
		Addr:    ":8080",
		Handler: h,
	}

	go func() {
		s.ListenAndServe()
	}()
	fmt.Println("server is starting on port 8080 ...")

	<-done
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	if err := s.Shutdown(ctx); err != nil {
		log.Fatalf("Server Shutdown Failed: %s\n", err.Error())
	}
	fmt.Println("Server Exited Properly")

}
