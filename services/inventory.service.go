package services

import "msg-inventory/models"

type InventoryService interface {
	CreateInventory() (*models.Inventory, error)
	UpdateInventory() (*models.Inventory, error)
	DeleteInventory() (*models.Inventory, error)
}
