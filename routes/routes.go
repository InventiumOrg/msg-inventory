package routes

import (
	"msg-inventory/handlers"
	"msg-inventory/middlewares"

	"github.com/gin-gonic/gin"
)

type Route struct {
	router *gin.Engine
}

func (r *Route) routes() {
	router := gin.Default()
	router.Use(middlewares.Authenticate())
	router.GET("/inventory", handlers.GetInventory)
	router.POST("/inventory", handlers.CreateInventory)
}
