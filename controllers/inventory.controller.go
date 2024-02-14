package controllers

import "msg-inventory/services"

type InventoryController struct {
	inventorySerivce services.InventoryService
}

func NewInventoryController(inventoryService services.InventoryService) InventoryController {
	return InventoryController{inventoryService}
}
