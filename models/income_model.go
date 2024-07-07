package models

import (
	"time"
)

type Income struct {
	IncomeID  uint      `gorm:"primaryKey" json:"income_id"`
	UserID    uint      `gorm:"not null" json:"user_id"`
	Amount    float64   `gorm:"not null" json:"amount"`
	Source    string    `json:"source"`
	Date      time.Time `gorm:"not null" json:"date"`
	Notes     string    `json:"notes,omitempty"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
