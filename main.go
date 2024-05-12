package main

import "github.com/gin-gonic/gin"

type Server struct {
	router *gin.Engine
}

func main() {
	router := gin.Default()
	router.Run(":8080")
}
