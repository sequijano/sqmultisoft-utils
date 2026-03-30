package models

import "time"

// DispatchRecord tracks a unit of stock that has been sold but not yet
// physically dispatched from the warehouse.
// Only created when InventorySettings.HasDispatchWarehouse = true.
type DispatchRecord struct {
	ID           uint       `json:"id"            gorm:"primarykey"`
	CreatedAt    time.Time  `json:"created_at"`
	TenantID     uint       `json:"tenant_id"     gorm:"index;not null"`
	SaleID       *uint      `json:"sale_id,omitempty"`
	ProductID    uint       `json:"product_id"    gorm:"not null"`
	ProductName  string     `json:"product_name"  gorm:"size:200"`
	Quantity     float64    `json:"quantity"      gorm:"not null"`
	Status       string     `json:"status"        gorm:"type:varchar(15);default:'pending'"` // pending | dispatched
	DispatchedAt *time.Time `json:"dispatched_at,omitempty"`
	Notes        string     `json:"notes"         gorm:"size:300"`
}
