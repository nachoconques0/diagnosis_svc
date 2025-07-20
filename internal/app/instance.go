package app

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/rs/zerolog/log"

	diagController "github.com/nachoconques0/diagnosis_svc/internal/controller/diagnosis"
	patientController "github.com/nachoconques0/diagnosis_svc/internal/controller/patient"
	userController "github.com/nachoconques0/diagnosis_svc/internal/controller/user"
	"github.com/nachoconques0/diagnosis_svc/internal/db"
	httpServer "github.com/nachoconques0/diagnosis_svc/internal/http"
	"github.com/nachoconques0/diagnosis_svc/internal/repo/diagnosis"
	"github.com/nachoconques0/diagnosis_svc/internal/repo/patient"
	"github.com/nachoconques0/diagnosis_svc/internal/repo/user"
	diagService "github.com/nachoconques0/diagnosis_svc/internal/service/diagnosis"
	patientService "github.com/nachoconques0/diagnosis_svc/internal/service/patient"
	userService "github.com/nachoconques0/diagnosis_svc/internal/service/user"
)

const (
	serviceNEW = iota
	serviceRUNNING
	serviceSTOPPING
	serviceSTOPPED
)

type Instance struct {
	servers []server
	timeout int
	state   int
	mu      sync.Mutex
}

type server interface {
	Run() error
	Stop(context.Context) error
}

func New(opts ...Option) error {
	options := Options{}
	for _, o := range opts {
		o(&options)
	}

	dbConn, err := db.New(options.dbOptions...)
	if err != nil {
		return err
	}

	// Init repo with DB
	patientRepo := patient.NewRepository(dbConn)
	diagRepo := diagnosis.NewRepository(dbConn)
	userRepo := user.NewRepository(dbConn)

	// Service
	patientSvc := patientService.New(patientRepo)
	diagSvc := diagService.New(diagRepo)
	userSvc := userService.New(userRepo)

	// Controller
	diagCtrl := diagController.New(diagSvc)
	patientCtrl := patientController.New(patientSvc)
	userCtrl := userController.New(userSvc, options.jwtSecret)

	// HTTP Server
	httpSrv, err := httpServer.New(httpServer.WithAddress(fmt.Sprintf(":%s", options.httpPort)))
	if err != nil {
		return err
	}
	httpRouter := httpServer.InitHTTPRouter(httpSrv)
	httpServer.InitRoutes(httpRouter, *userCtrl, *diagCtrl, *patientCtrl)

	i := Instance{
		timeout: 20,
		servers: []server{httpSrv},
	}

	quitCh := make(chan os.Signal, 1)
	signal.Notify(quitCh, syscall.SIGTERM, os.Interrupt)
	defer signal.Stop(quitCh)

	return i.Run(quitCh)
}

func (s *Instance) Run(quit chan os.Signal) error {
	if s.isStopped() {
		return errors.New("instance was already stopped. Can't start it again")
	}

	if s.isNew() {
		log.Info().Msg("Application: starting...")

		s.mu.Lock()
		s.state = serviceRUNNING
		errCh := make(chan error, 1)
		s.runServers(errCh)
		s.mu.Unlock()

		log.Info().Msg("Application: running...")

		select {
		case err := <-errCh:
			return err
		case <-quit:
			s.mu.Lock()
			s.state = serviceSTOPPING
			s.mu.Unlock()

			log.Info().Msgf("Application: stopping in %.0fs...", s.Timeout().Seconds())
			ctx, cancel := context.WithTimeout(context.Background(), s.Timeout())
			defer cancel()
			go s.stopServers(ctx)
			<-ctx.Done()

			s.mu.Lock()
			s.state = serviceSTOPPED
			s.mu.Unlock()
			log.Info().Msg("Application: stopped")
		}
	}
	return nil
}

func (s *Instance) runServers(errCh chan error) {
	for _, srv := range s.servers {
		go func(s server, ch chan error) {
			if err := s.Run(); err != nil {
				ch <- err
			}
		}(srv, errCh)
	}
}

func (s *Instance) stopServers(ctx context.Context) {
	for _, srv := range s.servers {
		_ = srv.Stop(ctx)
	}
}

func (s *Instance) Timeout() time.Duration {
	return time.Duration(s.timeout) * time.Second
}

func (s *Instance) isStopped() bool {
	return s.state == serviceSTOPPED
}

func (s *Instance) isNew() bool {
	return s.state == serviceNEW
}
