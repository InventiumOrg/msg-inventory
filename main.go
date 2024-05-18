package main

import (
	"msg-inventory/config"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Run(":13740")
	config.LoadConfig("app.env")
}
