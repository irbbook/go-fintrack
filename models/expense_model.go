package models

import (
	"time"
)

type Expense struct {
	ExpenseID uint      `gorm:"primaryKey" json:"expense_id"`
	UserID    uint      `gorm:"not null" json:"user_id"`
	Amount    float64   `gorm:"not null" json:"amount"`
	Category  string    `json:"category"`
	Date      time.Time `gorm:"not null" json:"date"`
	Notes     string    `json:"notes,omitempty"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
