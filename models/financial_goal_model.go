package models

import (
	"time"
)

type FinancialGoal struct {
	GoalID        uint      `gorm:"primaryKey" json:"goal_id"`
	UserID        uint      `gorm:"not null" json:"user_id"`
	Name          string    `gorm:"not null" json:"name"`
	TargetAmount  float64   `gorm:"not null" json:"target_amount"`
	CurrentAmount float64   `gorm:"default:0" json:"current_amount"`
	TargetDate    time.Time `gorm:"not null" json:"target_date"`
	Notes         string    `json:"notes,omitempty"`
	CreatedAt     time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt     time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
