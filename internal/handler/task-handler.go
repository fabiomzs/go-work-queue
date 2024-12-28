package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func TaskListHandler(ctx *gin.Context) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{
		"message": "operation from tasks successfull",
		"data":    "Contnet",
	})
}

func TaskByIdHandler(ctx *gin.Context) {
	ctx.Header("Content-type", "application/json")

	id := ctx.Param("id")

	ctx.JSON(http.StatusOK, gin.H{
		"message": "operation from tasks successfull",
		"data":    fmt.Sprintf("Contnet %s", id),
	})
}
