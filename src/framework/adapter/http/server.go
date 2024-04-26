package http

import (
	"context"
	"net"
	"net/http"

	"go.uber.org/fx"
)

func NewServer(lifecycle fx.Lifecycle, router *http.ServeMux) *http.Server {
	server := &http.Server{Addr: ":8080", Handler: router}
	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			listener, err := net.Listen("tcp", server.Addr)
			if err != nil {
				return err
			}
			go server.Serve(listener)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return server.Shutdown(ctx)
		},
	})
	return server
}

func Invoke(*http.Server) {}
