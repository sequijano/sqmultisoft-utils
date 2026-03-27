package models

import (
	"time"

	"gorm.io/gorm"
)

type Plan string

const (
	PlanBasic      Plan = "basic"
	PlanPro        Plan = "pro"
	PlanEnterprise Plan = "enterprise"
)

type Tenant struct {
	ID                  uint           `json:"id" gorm:"primarykey"`
	CreatedAt           time.Time      `json:"created_at"`
	UpdatedAt           time.Time      `json:"-"`
	DeletedAt           gorm.DeletedAt `json:"-" gorm:"index"`
	Name                string         `json:"name" gorm:"not null"`
	Slug                string         `json:"slug" gorm:"uniqueIndex;not null"`
	Plan                Plan           `json:"plan" gorm:"type:varchar(20);default:'basic'"`
	LicenseKey          string         `json:"license_key" gorm:"uniqueIndex"`
	LicenseExpiresAt    time.Time      `json:"license_expires_at"`
	Active              bool           `json:"active" gorm:"default:true"`
	IsDemo              bool           `json:"is_demo" gorm:"default:true"`
	ShowBookmarkHint    bool           `json:"show_bookmark_hint" gorm:"default:true"`
	FiscalPrinterModel  string         `json:"fiscal_printer_model" gorm:"size:50"`
	LastLoginAt         *time.Time     `json:"last_login_at"`
}
