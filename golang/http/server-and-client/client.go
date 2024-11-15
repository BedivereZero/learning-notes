package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"net/http"

	"github.com/spf13/cobra"
)

func NewClientCommand() *cobra.Command {
	var (
		network string
		address string

		method string
		path   string
	)
	cmd := &cobra.Command{
		Use: "client",
		RunE: func(cmd *cobra.Command, args []string) error {
			client := &http.Client{
				Transport: &UnixSocketTransport{
					network: network,
					address: address,
				},
			}

			resp, err := client.Get(path)
			if err != nil {
				return fmt.Errorf("send http request fail: %v", err)
			}
			defer resp.Body.Close()

			if resp.StatusCode != http.StatusOK {
				body, err := io.ReadAll(resp.Body)
				if err != nil {
					body = fmt.Appendf(body, "read response body fail: %v", err)
				}
				return fmt.Errorf("get %s fail: code=%d, body=%s", path, resp.StatusCode, body)
			}

			_, err = io.Copy(cmd.OutOrStdout(), resp.Body)
			return err
		},
	}
	cmd.Flags().StringVar(&network, "network", "", "connect to the network")
	cmd.Flags().StringVar(&address, "address", "", "connect to the address")
	cmd.Flags().StringVar(&method, "method", "", "http request method")
	cmd.Flags().StringVar(&path, "path", "", "http request path")

	return cmd
}

type UnixSocketTransport struct {
	network string
	address string
}

// RoundTrip implements http.RoundTripper.
func (t *UnixSocketTransport) RoundTrip(req *http.Request) (resp *http.Response, err error) {
	dialer := &net.Dialer{}

	conn, err := dialer.DialContext(req.Context(), t.network, t.address)
	if err != nil {
		err = fmt.Errorf("dial %s:%s fail: %w", t.network, t.address, err)
		return
	}
	defer conn.Close()

	if err = req.Write(conn); err != nil {
		err = fmt.Errorf("write http request to connection fail: %w", err)
		return
	}

	return http.ReadResponse(bufio.NewReader(conn), req)
}

var _ http.RoundTripper = &UnixSocketTransport{}
