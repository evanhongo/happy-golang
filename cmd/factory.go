package main

import (
	"context"
	"net/http"
	"os/signal"
	"syscall"

	"github.com/evanhongo/happy-golang/api"
	auth_route "github.com/evanhongo/happy-golang/api/route/auth"
	docs_route "github.com/evanhongo/happy-golang/api/route/docs"
	health_route "github.com/evanhongo/happy-golang/api/route/health"
	job_route "github.com/evanhongo/happy-golang/api/route/job"
	"github.com/evanhongo/happy-golang/internal/env"
	job_queue "github.com/evanhongo/happy-golang/pkg/job_queue"
	"github.com/evanhongo/happy-golang/pkg/logger"
	"golang.org/x/sync/errgroup"
)

type Cmd struct {
	server   *api.Server
	jobQueue job_queue.IJobQueue
}

func (c *Cmd) Execute() error {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	eg, errCtx := errgroup.WithContext(ctx)

	eg.Go(func() error {
		defer logger.Info("Stop http server")
		c.server.Init()

		docsRouter, _ := docs_route.CreateRouter()
		healthRouter, _ := health_route.CreateRouter()
		authRouter, _ := auth_route.CreateRouter()
		jobRouter, _ := job_route.CreateRouter()

		env := env.GetEnv()
		if env.ENVIRONMENT != "production" {
			c.server.RegisterRouter(docsRouter)
		}

		c.server.RegisterRouter(healthRouter)
		c.server.RegisterRouter(authRouter)
		c.server.RegisterRouter(jobRouter)

		if err := c.server.Start(); err != nil && err != http.ErrServerClosed {
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
		return c.server.Shutdown(errCtx)
	})

	return eg.Wait()
}

func NewCmd(server *api.Server, jobQueue job_queue.IJobQueue) *Cmd {
	return &Cmd{
		server,
		jobQueue,
	}
}
