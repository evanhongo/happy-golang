package main

import (
	"context"
	"net/http"
	"os/signal"
	"syscall"

	job_queue "github.com/evanhongo/happy-golang/internal/job_queue"
	"github.com/evanhongo/happy-golang/pkg/logger"
	"golang.org/x/sync/errgroup"
)

type Server struct {
	httpServer *http.Server
	jobQueue   job_queue.IJobQueue
}

func (s *Server) Start() error {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	errWg, errCtx := errgroup.WithContext(ctx)

	errWg.Go(func() error {
		defer logger.Info("Stop http server")
		if err := s.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			return err
		}
		return nil
	})

	errWg.Go(func() error {
		defer logger.Info("Stop job queue")
		if err := s.jobQueue.Start(); err != nil {
			return err
		}
		return nil
	})

	errWg.Go(func() error {
		<-errCtx.Done()
		return s.httpServer.Shutdown(ctx)
	})

	return errWg.Wait()
}

func NewServer(httpServer *http.Server, jobQueue job_queue.IJobQueue) *Server {
	return &Server{
		httpServer,
		jobQueue,
	}
}
