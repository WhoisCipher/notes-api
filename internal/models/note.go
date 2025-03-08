package models

import (
    _ "gorm.io/gorm"
    "time"
)

type Note struct {
    ID        uint      `json:"id" gorm:"primaryKey"`
    UserID    uint      `json:"user_id" gorm:"not null"`
    Title     string    `json:"title" gorm:"not null;size:255"`
    Content   string    `json:"content" gorm:"not null"`
    CreatedAt time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
    UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`

    User User `json:"-" gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
}


