package models

import (
	"time"

	"gorm.io/gorm"
)

type Currency struct {
	ID        uint           `json:"id" gorm:"primarykey"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
	Code      string         `json:"code" gorm:"uniqueIndex;size:3;not null"`
	Name      string         `json:"name"`
	Symbol    string         `json:"symbol"`
}

type ExchangeRate struct {
	ID          uint           `json:"id" gorm:"primarykey"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"-"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
	TenantID    uint           `json:"tenant_id" gorm:"index;not null;default:1"`
	USDtoVES    float64        `json:"usd_to_ves" gorm:"column:usd_to_ves;not null"`
	USDtoCOP    float64        `json:"usd_to_cop" gorm:"column:usd_to_cop;not null"`
	UpdatedByID uint           `json:"updated_by_id"`
	UpdatedBy   User           `json:"updated_by" gorm:"foreignKey:UpdatedByID"`
}
