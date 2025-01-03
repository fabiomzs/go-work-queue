package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary      List all tasks
// @Description  List all valid tasks
// @Tags         Tasks
// @Accept       json
// @Produce      json
// @Success      200 {object} TaskResponse
// @Failure      500 {object} ErrorResponse
// @Router       /tasks [get]
func TaskListHandler(ctx *gin.Context) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{
		"message": "operation from tasks successfull",
		"data":    "Contnet",
	})
}

// @Summary      Get task by id
// @Description  Get task by id
// @Tags         Tasks
// @Accept       json
// @Produce      json
// @Success      200 {object} TaskResponse
// @Failure      500 {object} ErrorResponse
// @Param 	     id path int true "Task ID"
// @Router       /tasks/{id} [get]
func TaskByIdHandler(ctx *gin.Context) {
	ctx.Header("Content-type", "application/json")

	id := ctx.Param("id")

	ctx.JSON(http.StatusOK, gin.H{
		"message": "operation from tasks successfull",
		"data":    fmt.Sprintf("Contnet %s", id),
	})
}
