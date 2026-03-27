package models

import "time"

type LicensePayment struct {
	ID                 uint       `json:"id" gorm:"primarykey"`
	CreatedAt          time.Time  `json:"created_at"`
	TenantID           uint       `json:"tenant_id" gorm:"index;not null"`
	Months             int        `json:"months" gorm:"not null"`
	AmountCOP          int64      `json:"amount_cop" gorm:"not null"`
	WompiTransactionID string     `json:"wompi_transaction_id" gorm:"uniqueIndex;size:100"`
	WompiReference     string     `json:"wompi_reference" gorm:"size:100"`
	Status             string     `json:"status" gorm:"size:20;default:'PENDING'"`
	PaidAt             *time.Time `json:"paid_at"`
}
