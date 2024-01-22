package models

import "time"

type Photo struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Title     string    `gorm:"not null" json:"title" valid:"required"`
	Caption   string    `gorm:"not null" json:"caption" valid:"required"`
	PhotoURL  string    `gorm:"not null" json:"photo_url" valid:"required"`
	UserID    uint      `gorm:"not null" json:"user_id" valid:"required"`
	User      User      `gorm:"foreignKey:UserID" json:"User"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at" time_format:"2006-01-02 15:04:05"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at" time_format:"2006-01-02 15:04:05"`
	Comment   []Comment `gorm:"foreignKey:PhotoID" json:"Comment"`
}
