package server

import (
	docs "github.com/fabiomzs/go.work-queue/docs"
	"github.com/fabiomzs/go.work-queue/internal/handler"
	//swaggerFiles "github.com/swaggo/files"
	//ginSwagger "github.com/swaggo/gin-swagger"
)

func (s server) registerRoutes() {
	basePath := "/api/v1"
	docs.SwaggerInfo.BasePath = basePath

	//url := ginSwagger.URL("http://localhost:3000/swagger/doc.json")
	//s.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	v1 := s.router.Group(basePath)

	v1.GET("/tasks", handler.TaskListHandler)
	v1.GET("/tasks/:id", handler.TaskByIdHandler)
}
