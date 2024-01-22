package models

import "time"

type Comment struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	UserID    uint      `gorm:"not null" json:"user_id" valid:"required"`
	PhotoID   uint      `gorm:"not null" json:"photo_id" valid:"required"`
	Message   string    `gorm:"not null" json:"message" valid:"required"`
	User      User      `gorm:"foreignKey:UserID" json:"User"`
	Photo     Photo     `gorm:"foreignKey:PhotoID" json:"Photo"`
	CreatedAt time.Time `gorm:"not null" json:"created_at" time_format:"2006-01-02 15:04:05"`
	UpdateAt  time.Time `gorm:"not null" json:"updated_at" time_format:"2006-01-02 15:04:05"`
}
