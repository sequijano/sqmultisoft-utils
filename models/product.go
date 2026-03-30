package models

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID               uint           `json:"id" gorm:"primarykey"`
	CreatedAt        time.Time      `json:"-"`
	UpdatedAt        time.Time      `json:"-"`
	DeletedAt        gorm.DeletedAt `json:"-" gorm:"index"`
	TenantID         uint           `json:"tenant_id" gorm:"index;not null;default:1"`
	Name             string         `json:"name" gorm:"not null"`
	Barcode          string         `json:"barcode" gorm:"size:100;index"`
	CategoryID       uint           `json:"category_id"`
	Category         Category       `json:"category"`
	BasePrice        float64        `json:"base_price" gorm:"not null"`
	BaseCurrencyCode string         `json:"base_currency_code" gorm:"size:3;not null"`
	BasePriceUSD     float64        `json:"base_price_usd" gorm:"default:0"`
	Stock            float64        `json:"stock" gorm:"default:0"`          // existencia de ventas
	StockPhysical    float64        `json:"stock_physical" gorm:"default:0"` // existencia física
	Active           bool           `json:"active" gorm:"default:true"`
	UnitType         string         `json:"unit_type" gorm:"size:10;not null;default:'unit'"`
}
