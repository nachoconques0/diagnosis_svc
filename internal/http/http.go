package server

import (
	"context"
	"fmt"
	"net"
	netHTTP "net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type Server interface {
	Address() string
	Run() error
	Stop(context.Context) error
}

// Server HTTP server definition
type server struct {
	opts   Options
	router *gin.Engine
	server *netHTTP.Server
}

// New Initialize a new HTTP server with gin framework
func New(opts ...Option) (*server, error) { //nolint:revive // no need to return Server because we need is the struct not the interface
	// Init and apply options
	options := Options{
		Context: context.Background(),
	}
	for _, o := range opts {
		o(&options)
	}

	// Validate the address
	_, _, err := net.SplitHostPort(options.Address)
	if err != nil {
		return nil, err
	}

	// Initialize the server
	router := gin.New()

	s := server{
		opts:   options,
		router: router,
		server: &netHTTP.Server{
			Addr:              options.Address,
			Handler:           router,
			ReadHeaderTimeout: time.Second * 60,
		},
	}

	return &s, nil
}

func InitHTTPRouter(srv *server) *gin.Engine {
	r := srv.router
	r.Use(gin.Logger())
	// Health endpoint
	r.GET("/health", func(ctx *gin.Context) {
		ctx.Status(netHTTP.StatusOK)
	})
	return r
}

// Address Return address where the server is running
func (s server) Address() string {
	return s.opts.Address
}

// Run runs the server and starts listening to HTTP requests
func (s server) Run() error {
	log.Info().Msg(fmt.Sprintf("HTTP server listening on :%s", s.Address()))
	return s.server.ListenAndServe()
}

// Stop Stops the server without interrupting any active connection
func (s server) Stop(ctx context.Context) error {
	log.Info().Msg("HTTP server: graceful stop...")
	err := s.server.Shutdown(ctx)
	log.Info().Msg("HTTP server: gracefully stopped!")
	return err
}

// Router return gub router instance
func (s server) Router() *gin.Engine {
	return s.router
}
