package models

import (
	"time"

	"gorm.io/gorm"
)

type PaymentMethod string

const (
	PaymentMethodCash     PaymentMethod = "cash"
	PaymentMethodCard     PaymentMethod = "card"
	PaymentMethodTransfer PaymentMethod = "transfer"
)

type SaleStatus string

const (
	SaleStatusCompleted   SaleStatus = "completed"
	SaleStatusVoided      SaleStatus = "voided"
	SaleStatusShiftClosed SaleStatus = "shift_closed"
)

type SalePayment struct {
	ID              uint          `json:"id" gorm:"primarykey"`
	CreatedAt       time.Time     `json:"-"`
	SaleID          uint          `json:"sale_id"`
	PaymentCurrency string        `json:"payment_currency" gorm:"size:3;not null"`
	PaymentMethod   PaymentMethod `json:"payment_method" gorm:"not null"`
	Amount          float64       `json:"amount" gorm:"not null"`
	AmountUSD       float64       `json:"amount_usd" gorm:"not null"`
	Reference       string        `json:"reference" gorm:"size:100"`
}

type Sale struct {
	ID                uint           `json:"id" gorm:"primarykey"`
	CreatedAt         time.Time      `json:"created_at"`
	UpdatedAt         time.Time      `json:"-"`
	DeletedAt         gorm.DeletedAt `json:"-" gorm:"index"`
	TenantID          uint           `json:"tenant_id" gorm:"index;not null;default:1"`
	CustomerID        *uint          `json:"customer_id" gorm:"index"`
	Customer          *Customer      `json:"customer,omitempty" gorm:"foreignKey:CustomerID"`
	CustomerName      string         `json:"customer_name" gorm:"size:255"`
	CustomerDocType   string         `json:"customer_doc_type" gorm:"size:20"`
	CustomerDocNumber string         `json:"customer_doc_number" gorm:"size:50"`
	CashierID         uint           `json:"cashier_id"`
	Cashier           User           `json:"cashier" gorm:"foreignKey:CashierID"`
	Items             []SaleItem     `json:"items"`
	Payments          []SalePayment  `json:"payments"`
	TotalUSD          float64        `json:"total_usd"`
	TotalVES          float64        `json:"total_ves"`
	TotalCOP          float64        `json:"total_cop"`
	AmountPaidUSD     float64        `json:"amount_paid_usd"`
	ChangeUSD         float64        `json:"change_usd"`
	ChangeCurrency    string         `json:"change_currency" gorm:"size:3;not null;default:'COP'"`
	ExRateUSDtoVES    float64        `json:"ex_rate_usd_to_ves"`
	ExRateUSDtoCOP    float64        `json:"ex_rate_usd_to_cop"`
	Status            SaleStatus     `json:"status" gorm:"default:'completed'"`
	DailyReportID     *uint          `json:"daily_report_id" gorm:"index"`
	ShiftNumber       int            `json:"shift_number"`
	SaleDate          time.Time      `json:"sale_date"`
	Notes             string         `json:"notes"`
}

type SaleItem struct {
	ID           uint           `json:"id" gorm:"primarykey"`
	CreatedAt    time.Time      `json:"-"`
	UpdatedAt    time.Time      `json:"-"`
	DeletedAt    gorm.DeletedAt `json:"-" gorm:"index"`
	SaleID       uint           `json:"sale_id"`
	ProductID    uint           `json:"product_id"`
	ProductName  string         `json:"product_name"`
	UnitType     string         `json:"unit_type" gorm:"size:10;not null;default:'unit'"`
	Quantity     float64        `json:"quantity" gorm:"type:decimal(10,3);not null"`
	UnitPriceUSD float64        `json:"unit_price_usd"`
	UnitPriceVES float64        `json:"unit_price_ves"`
	UnitPriceCOP float64        `json:"unit_price_cop"`
	SubtotalUSD  float64        `json:"subtotal_usd"`
	SubtotalVES  float64        `json:"subtotal_ves"`
	SubtotalCOP  float64        `json:"subtotal_cop"`
}
