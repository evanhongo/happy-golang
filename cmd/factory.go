package main

import (
	"context"
	"net/http"
	"os/signal"
	"syscall"

	job_queue "github.com/evanhongo/happy-golang/pkg/job_queue"
	"github.com/evanhongo/happy-golang/pkg/logger"
	"golang.org/x/sync/errgroup"
)

type Cmd struct {
	httpServer *http.Server
	jobQueue   job_queue.IJobQueue
}

func (c *Cmd) Execute() error {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	eg, errCtx := errgroup.WithContext(ctx)

	eg.Go(func() error {
		defer logger.Info("Stop http server")
		if err := c.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			return err
		}
		return nil
	})

	eg.Go(func() error {
		defer logger.Info("Stop job queue")
		if err := c.jobQueue.Start(); err != nil {
			return err
		}
		return nil
	})

	eg.Go(func() error {
		<-errCtx.Done()
		return c.httpServer.Shutdown(ctx)
	})

	return eg.Wait()
}

func NewCmd(httpServer *http.Server, jobQueue job_queue.IJobQueue) *Cmd {
	return &Cmd{
		httpServer,
		jobQueue,
	}
}
