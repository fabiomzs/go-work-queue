package main

import (
	"context"

	_ "github.com/fabiomzs/go.work-queue/docs"

	"github.com/fabiomzs/go.work-queue/internal/logger"
	"github.com/fabiomzs/go.work-queue/internal/server"
)

// @title Tag Task API
// @version 1.0
// @description API in Go for work queue pattern
// @host localhost:3000
// @BasePath /api/v1
func main() {
	ctx := context.Background()

	s := server.NewServer(server.ServerOptions{
		Context: ctx,
		Logger:  logger.NewSimpleLogger(),
	})

	s.Start()
}
