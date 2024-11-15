package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/spf13/cobra"
)

func NewServerCommand() *cobra.Command {
	var (
		network string
		address string
	)
	cmd := &cobra.Command{
		Use: "server",
		RunE: func(cmd *cobra.Command, args []string) error {
			config := &net.ListenConfig{}
			ln, err := config.Listen(cmd.Context(), network, address)
			if err != nil {
				return err
			}
			defer ln.Close()

			mux := http.NewServeMux()
			mux.HandleFunc("/", greet)

			server := &http.Server{
				Handler: mux,
			}

			// Start HTTP Server
			go func() {
				log.Printf("start http server on %s", ln.Addr())
				if err := server.Serve(ln); err != nil && !errors.Is(err, http.ErrServerClosed) {
					log.Printf("start http server fail: %v", err)
				}
			}()

			stop := make(chan os.Signal, 1)
			signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

			// wait for signal os.Interrupt
			<-stop
			log.Printf("shutdown http server")

			ctx, cancel := context.WithTimeout(cmd.Context(), time.Second*5)
			defer cancel()

			if err := server.Shutdown(ctx); err != nil {
				log.Printf("shutdown http server fail: %v", err)
			}

			return nil
		},
	}
	cmd.Flags().StringVar(&network, "network", "", "listen on the network")
	cmd.Flags().StringVar(&address, "address", "", "listen to the address")

	return cmd
}

func greet(w http.ResponseWriter, r *http.Request) {
	log.Printf("receive http request: method=%s, url=%s, remote=%s", r.Method, r.URL, r.RemoteAddr)

	fmt.Fprintf(w, "Hello World! %s\n", time.Now().Format(time.RFC3339))
}
