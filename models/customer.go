package models

import (
	"time"

	"gorm.io/gorm"
)

type DocumentType struct {
	ID           uint           `json:"id" gorm:"primarykey"`
	CreatedAt    time.Time      `json:"-"`
	UpdatedAt    time.Time      `json:"-"`
	DeletedAt    gorm.DeletedAt `json:"-" gorm:"index"`
	TenantID     uint           `json:"tenant_id" gorm:"uniqueIndex:idx_doctype_tenant_abbr;not null"`
	Abbreviation string         `json:"abbreviation" gorm:"uniqueIndex:idx_doctype_tenant_abbr;size:10;not null"`
	Description  string         `json:"description" gorm:"not null"`
	IsDefault    bool           `json:"is_default" gorm:"default:false"`
}

type Customer struct {
	ID             uint           `json:"id" gorm:"primarykey"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"-"`
	DeletedAt      gorm.DeletedAt `json:"-" gorm:"index"`
	TenantID       uint           `json:"tenant_id" gorm:"uniqueIndex:idx_customer_tenant_doc;not null"`
	DocumentTypeID uint           `json:"document_type_id" gorm:"uniqueIndex:idx_customer_tenant_doc;not null"`
	DocumentType   DocumentType   `json:"document_type" gorm:"foreignKey:DocumentTypeID"`
	DocumentNumber string         `json:"document_number" gorm:"uniqueIndex:idx_customer_tenant_doc;size:50;not null"`
	Name           string         `json:"name" gorm:"not null"`
}
