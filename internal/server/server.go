package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/fabiomzs/go.work-queue/internal/logger"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Server interface {
	Start()
}

type ServerOptions struct {
	Logger  logger.Logger
	Context context.Context
}

type server struct {
	router *gin.Engine
	ServerOptions
}

func NewServer(options ServerOptions) Server {
	return server{
		router:        gin.New(),
		ServerOptions: options,
	}
}

func (s server) Start() {
	s.setupServer()
	s.setupSwagger()
	s.setupMiddlewares()
	s.registerRoutes()
	s.startGracefulShutdown()
}

func (s server) setupServer() {
	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, _ string, _ int) {
		s.Logger.Info(fmt.Sprintf("mapped %v %v route", httpMethod, absolutePath))
	}

	s.router.SetTrustedProxies(nil)
}

func (s server) setupMiddlewares() {
	s.router.Use(gin.Recovery())
}

func (s server) setupSwagger() {
	url := ginSwagger.URL("http://localhost:3000/swagger/doc.json")
	s.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
}

func (s server) startGracefulShutdown() {
	port := ":3000"
	srv := &http.Server{
		Addr:    port,
		Handler: s.router,
	}

	go func() {
		s.Logger.Info(fmt.Sprintf("starting server on port %s...", port))
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			s.Logger.Error(fmt.Sprintf("start server error %s", err.Error()))
			os.Exit(1)
		}
	}()

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	q := <-quit

	s.Logger.Info(fmt.Sprintf("server shutdown... %s", q))

	ctx, cancel := context.WithTimeout(s.Context, 5*time.Second)

	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		s.Logger.Error("server shutdown ", err)
		os.Exit(1)
	}

	<-ctx.Done()

	s.Logger.Info("server shutdown done")
}
