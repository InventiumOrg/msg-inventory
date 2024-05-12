package handlers

import (
	"github.com/gin-gonic/gin"
)

type getInventoryRequest struct {
	Name string `json:"name"`
}

func GetInventory(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Get Inventory"})
}

func listInventory(c *gin.Context) {
	c.JSON(200, gin.H{"message": "List Inventory"})
}

func updateInventory(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Update Inventory"})
}

func CreateInventory(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Create Inventory"})
}
