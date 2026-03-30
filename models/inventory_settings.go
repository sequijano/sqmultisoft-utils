package models

// DispatchPrintMode controls how dispatch notes are printed.
type DispatchPrintMode string

const (
	DispatchPrintNone   DispatchPrintMode = "none"
	DispatchPrintTicket DispatchPrintMode = "ticket"
	DispatchPrintLetter DispatchPrintMode = "letter"
)

// InventorySettings holds per-tenant inventory configuration.
type InventorySettings struct {
	TenantID             uint              `json:"tenant_id"              gorm:"primarykey"`
	HasDispatchWarehouse bool              `json:"has_dispatch_warehouse" gorm:"default:false"`
	AllowNegativeStock   bool              `json:"allow_negative_stock"   gorm:"default:false"`
	DispatchPrintMode    DispatchPrintMode `json:"dispatch_print_mode"    gorm:"type:varchar(10);default:'none'"`
}
