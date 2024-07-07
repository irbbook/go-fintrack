package models

import (
	"time"
)

type Budget struct {
	BudgetID  uint      `gorm:"primaryKey" json:"budget_id"`
	UserID    uint      `gorm:"not null" json:"user_id"`
	Category  string    `gorm:"not null" json:"category"`
	Amount    float64   `gorm:"not null" json:"amount"`
	StartDate time.Time `gorm:"not null" json:"start_date"`
	EndDate   time.Time `gorm:"not null" json:"end_date"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
