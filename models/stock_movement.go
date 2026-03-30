package models

import "time"

// StockMovementReason identifies the source of a stock change.
type StockMovementReason string

const (
	StockReasonSale       StockMovementReason = "sale"
	StockReasonManual     StockMovementReason = "manual"
	StockReasonCategory   StockMovementReason = "category"
	StockReasonAdjustment StockMovementReason = "adjustment"
	StockReasonDispatch   StockMovementReason = "dispatch"
)

// StockMovement records every stock change (deduction or load) for a product.
type StockMovement struct {
	ID          uint                `json:"id" gorm:"primarykey"`
	CreatedAt   time.Time           `json:"created_at"`
	TenantID    uint                `json:"tenant_id" gorm:"index;not null"`
	ProductID   uint                `json:"product_id" gorm:"index;not null"`
	ProductName string              `json:"product_name" gorm:"size:200"`
	Delta       float64             `json:"delta"`                             // positive = load, negative = deduction
	Reason      StockMovementReason `json:"reason" gorm:"type:varchar(30);not null"`
	StockType   string              `json:"stock_type" gorm:"type:varchar(10);default:'both'"` // sales | physical | both
	ReferenceID *uint               `json:"reference_id,omitempty"`            // sale_id when reason=sale
	Notes       string              `json:"notes" gorm:"size:300"`
}
