package http

import (
	"context"
	"log"
	"net/http"
	"time"
)

type Server struct {
	server *http.Server
}

func NewServer(port string, handler http.Handler) *Server {
	return &Server{
		server: &http.Server{
			Addr:    ":" + port,
			Handler: handler,
			//ReadTimeout:  5 * time.Second,
			//WriteTimeout: 55 * time.Second,
		},
	}
}

func (e *Server) ListenAndServe() {
	go func() {
		if err := e.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Printf("error on ListenAndServe: %q", err)
		}
	}()
}

func (e *Server) Shutdown() {
	ctx, _ := context.WithTimeout(context.Background(), 60*time.Second)
	if err := e.server.Shutdown(ctx); err != nil && err != http.ErrServerClosed {
		log.Printf("could not shutdown in 60s: %q", err)
	}
}
