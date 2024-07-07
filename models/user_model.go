package models

import (
	"time"
)

type User struct {
	UserID         uint            `gorm:"primaryKey" json:"user_id"`
	Username       string          `gorm:"unique;not null" json:"username"`
	Email          string          `gorm:"unique;not null" json:"email"`
	PasswordHash   string          `gorm:"not null" json:"-"`
	CreatedAt      time.Time       `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt      time.Time       `gorm:"autoUpdateTime" json:"updated_at"`
	Incomes        []Income        `gorm:"foreignKey:UserID" json:"incomes,omitempty"`
	Expenses       []Expense       `gorm:"foreignKey:UserID" json:"expenses,omitempty"`
	Budgets        []Budget        `gorm:"foreignKey:UserID" json:"budgets,omitempty"`
	FinancialGoals []FinancialGoal `gorm:"foreignKey:UserID" json:"financial_goals,omitempty"`
}
