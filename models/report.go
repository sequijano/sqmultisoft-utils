package models

import (
	"time"

	"gorm.io/gorm"
)

type DailyIncomeSummary struct {
	Date        string  `json:"date"`
	ShiftsCount int     `json:"shifts_count"`
	CashUSD     float64 `json:"cash_usd"`
	TransferUSD float64 `json:"transfer_usd"`
	CashCOP     float64 `json:"cash_cop"`
	TransferCOP float64 `json:"transfer_cop"`
	CashVES     float64 `json:"cash_ves"`
	CardVES     float64 `json:"card_ves"`
	TransferVES float64 `json:"transfer_ves"`
	ChangeUSD   float64 `json:"change_usd"`
	ChangeCOP   float64 `json:"change_cop"`
	ChangeVES   float64 `json:"change_ves"`
}

type DailyReport struct {
	ID                    uint           `json:"id" gorm:"primarykey"`
	CreatedAt             time.Time      `json:"created_at"`
	UpdatedAt             time.Time      `json:"-"`
	DeletedAt             gorm.DeletedAt `json:"-" gorm:"index"`
	TenantID              uint           `json:"tenant_id" gorm:"index;not null;default:1"`
	Status                string         `json:"status" gorm:"size:20;not null;default:'completed'"`
	VoidedByID            *uint          `json:"voided_by_id"`
	VoidedBy              *User          `json:"voided_by,omitempty" gorm:"foreignKey:VoidedByID"`
	ReportDate            time.Time      `json:"report_date" gorm:"index"`
	ShiftNumber           int            `json:"shift_number"`
	ShiftStartTime        time.Time      `json:"shift_start_time"`
	ShiftEndTime          time.Time      `json:"shift_end_time"`
	SalesCount            int            `json:"sales_count"`
	TotalSalesUSD         float64        `json:"total_sales_usd"`
	TotalSalesVES         float64        `json:"total_sales_ves"`
	TotalSalesCOP         float64        `json:"total_sales_cop"`
	CashUSD               float64        `json:"cash_usd"`
	CashCOP               float64        `json:"cash_cop"`
	CashVES               float64        `json:"cash_ves"`
	CardVES               float64        `json:"card_ves"`
	TransferUSD           float64        `json:"transfer_usd"`
	TransferCOP           float64        `json:"transfer_cop"`
	TransferVES           float64        `json:"transfer_ves"`
	ChangeUSD             float64        `json:"change_usd"`
	ChangeCOP             float64        `json:"change_cop"`
	ChangeVES             float64        `json:"change_ves"`
	ClosingExRateUSDtoVES float64        `json:"closing_ex_rate_usd_to_ves"`
	ClosingExRateUSDtoCOP float64        `json:"closing_ex_rate_usd_to_cop"`
	ClosedByID            uint           `json:"closed_by_id"`
	ClosedBy              User           `json:"closed_by" gorm:"foreignKey:ClosedByID"`
	Notes                 string         `json:"notes"`
}
