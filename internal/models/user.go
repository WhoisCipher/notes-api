package models

import (
	_ "gorm.io/gorm"
	"time"
)

type User struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Username  string    `json:"username" gorm:"unique;not null;size:50"`
	Email     string    `json:"email" gorm:"unique;not null;size:100"`
	Password  string    `json:"-" gorm:"not null"`
	CreatedAt time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`

	Notes []Note `json:"notes" gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
}
