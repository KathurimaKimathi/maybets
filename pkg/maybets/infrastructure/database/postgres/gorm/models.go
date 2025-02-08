package gorm

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// AbstractBase model contains defines common fields across tables
type AbstractBase struct {
	ID        *string   `gorm:"primaryKey;unique;column:id"`
	CreatedAt time.Time `gorm:"column:created;not null"`
	UpdatedAt time.Time `gorm:"column:updated;not null"`
	CreatedBy *string   `gorm:"column:created_by"`
	UpdatedBy *string   `gorm:"column:updated_by"`
}

// BeforeCreate sets ID and CreatedAt before inserting a record
func (base *AbstractBase) BeforeCreate(_ *gorm.DB) error {
	if base.ID == nil {
		id := uuid.New().String()
		base.ID = &id
	}

	base.CreatedAt = time.Now()
	base.UpdatedAt = base.CreatedAt

	return nil
}

// BeforeUpdate updates UpdatedAt before modifying a record
func (base *AbstractBase) BeforeUpdate(_ *gorm.DB) error {
	base.UpdatedAt = time.Now()

	return nil
}

// Bet models the bet data class model
type Bet struct {
	AbstractBase
	BetID     string    `json:"bet_id" gorm:"column:bet_id;not null"`
	UserID    string    `json:"user_id" gorm:"column:user_id;not null"`
	Amount    float64   `json:"amount" gorm:"column:amount;not null"`
	Odds      float64   `json:"odds" gorm:"column:odds;not null"`
	Outcome   string    `json:"outcome" gorm:"column:outcome;not null"`
	Timestamp time.Time `json:"timestamp" gorm:"column:timestamp;not null"`
}

// TableName ....
func (Bet) TableName() string {
	return "bets"
}

type User struct {
	UserID    string `json:"user_id"`
	TotalBets int64  `json:"total_bets"`
}
