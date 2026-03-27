package models

import (
	"time"

	"gorm.io/gorm"
)

type Category struct {
	ID          uint           `json:"id" gorm:"primarykey"`
	CreatedAt   time.Time      `json:"-"`
	UpdatedAt   time.Time      `json:"-"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
	TenantID    uint           `json:"tenant_id" gorm:"uniqueIndex:idx_categories_tenant_name;not null;default:1"`
	Name        string         `json:"name" gorm:"uniqueIndex:idx_categories_tenant_name;not null"`
	Description string         `json:"description"`
	Visible     bool           `json:"visible" gorm:"default:true"`
	SortOrder   int            `json:"sort_order" gorm:"default:0"`
}
