package models

// InventorySettings holds per-tenant inventory configuration.
type InventorySettings struct {
	TenantID             uint `json:"tenant_id"              gorm:"primarykey"`
	HasDispatchWarehouse bool `json:"has_dispatch_warehouse" gorm:"default:false"`
}
