package main

import (
	"context"

	_ "github.com/fabiomzs/go.work-queue/docs"

	"github.com/fabiomzs/go.work-queue/internal/logger"
	"github.com/fabiomzs/go.work-queue/internal/server"
)

// @title Tag Task API
// @version 1.0
// @description API in Go for Work Queue Pattern PoC
// @termsOfService http://swagger.io/terms/

// @contact.name   Fabio Muniz
// @contact.url    http://fabiomuniz.com.br

// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:3000
// @BasePath /api/v1
// @schemes http
func main() {
	ctx := context.Background()

	s := server.NewServer(server.ServerOptions{
		Context: ctx,
		Logger:  logger.NewSimpleLogger(),
	})

	s.Start()
}
