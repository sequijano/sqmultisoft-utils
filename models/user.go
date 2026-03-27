package models

import (
	"time"

	"gorm.io/gorm"
)

type Role string

const (
	RoleAdmin   Role = "admin"
	RoleManager Role = "manager"
	RoleCashier Role = "cashier"
)

type User struct {
	ID                 uint           `json:"id" gorm:"primarykey"`
	CreatedAt          time.Time      `json:"created_at"`
	UpdatedAt          time.Time      `json:"-"`
	DeletedAt          gorm.DeletedAt `json:"-" gorm:"index"`
	TenantID           uint           `json:"tenant_id" gorm:"uniqueIndex:idx_users_username_tenant;index;not null;default:1"`
	Name               string         `json:"name"`
	Username           string         `json:"username" gorm:"uniqueIndex:idx_users_username_tenant;not null"`
	Password           string         `json:"-"`
	Role               Role           `json:"role" gorm:"type:varchar(20);not null"`
	Email              string         `json:"email" gorm:"size:255"`
	Active             bool           `json:"active" gorm:"default:true"`
	AuthToken          string         `json:"auth_token" gorm:"size:6"`
	MustChangePassword bool           `json:"must_change_password" gorm:"default:false"`
	SessionVersion     int            `json:"-" gorm:"default:0"`
}
