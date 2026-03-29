package models

import (
	"database/sql/driver"
	"fmt"
	"strings"
	"time"

	"gorm.io/gorm"
)

type Plan string

const (
	PlanBasic      Plan = "basic"
	PlanPro        Plan = "pro"
	PlanEnterprise Plan = "enterprise"
)

// StringSlice is a []string that serialises as a comma-separated TEXT column
// so it works on any Postgres/SQLite without requiring the pq driver.
type StringSlice []string

func (s StringSlice) Value() (driver.Value, error) {
	return strings.Join(s, ","), nil
}

func (s *StringSlice) Scan(value interface{}) error {
	if value == nil {
		*s = StringSlice{}
		return nil
	}
	var str string
	switch v := value.(type) {
	case string:
		str = v
	case []byte:
		str = string(v)
	default:
		return fmt.Errorf("StringSlice: unsupported type %T", value)
	}
	if str == "" {
		*s = StringSlice{}
		return nil
	}
	*s = strings.Split(str, ",")
	return nil
}

// DefaultModules are enabled for every new tenant.
var DefaultModules = StringSlice{"administracion", "ventas", "inventario"}

// StockDeductMode controls when inventory is decremented.
type StockDeductMode string

const (
	StockDeductOnSale    StockDeductMode = "al_vender"
	StockDeductOnDelivery StockDeductMode = "nota_de_entrega"
)

type Tenant struct {
	ID                  uint            `json:"id" gorm:"primarykey"`
	CreatedAt           time.Time       `json:"created_at"`
	UpdatedAt           time.Time       `json:"-"`
	DeletedAt           gorm.DeletedAt  `json:"-" gorm:"index"`
	Name                string          `json:"name" gorm:"not null"`
	Slug                string          `json:"slug" gorm:"uniqueIndex;not null"`
	Plan                Plan            `json:"plan" gorm:"type:varchar(20);default:'basic'"`
	LicenseKey          string          `json:"license_key" gorm:"uniqueIndex"`
	LicenseExpiresAt    time.Time       `json:"license_expires_at"`
	Active              bool            `json:"active" gorm:"default:true"`
	IsDemo              bool            `json:"is_demo" gorm:"default:true"`
	ShowBookmarkHint    bool            `json:"show_bookmark_hint" gorm:"default:true"`
	FiscalPrinterModel  string          `json:"fiscal_printer_model" gorm:"size:50"`
	LastLoginAt         *time.Time      `json:"last_login_at"`
	EnabledModules      StringSlice     `json:"enabled_modules" gorm:"type:text;default:'administracion,ventas,inventario'"`
	StockDeductMode     StockDeductMode `json:"stock_deduct_mode" gorm:"type:varchar(30);default:'al_vender'"`
}
