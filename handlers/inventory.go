package handlers

import (
	"github.com/gin-gonic/gin"
)

type getInventoryRequest struct {
	Name string `json:"name"`
}

func GetInventory(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "Get Inventory"})
}

func ListInventory(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "List Inventory"})
}

func UpdateInventory(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Update Inventory"})
}

func CreateInventory(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Create Inventory"})
}
