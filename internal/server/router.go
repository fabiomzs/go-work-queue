package server

import (
	docs "github.com/fabiomzs/go.work-queue/docs"
	"github.com/fabiomzs/go.work-queue/internal/handler"
)

func (s server) registerRoutes() {
	basePath := "/api/v1"
	docs.SwaggerInfo.BasePath = basePath

	v1 := s.router.Group(basePath)

	// Tasks         godoc
	// @Summary      List all tasks
	// @Description  List all valid tasks
	// @Tags         Tasks
	// @Accept       json
	// @Produce      json
	// @Success      200 {object} handler.TaskResponse
	// @Failure      500 {object} handler.ErrorResponse
	// @Router       /tasks [get]
	v1.GET("/tasks", handler.TaskListHandler)
	v1.GET("/tasks/:id", handler.TaskByIdHandler)
}
