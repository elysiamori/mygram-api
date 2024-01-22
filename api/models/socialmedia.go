package models

import "time"

type SocialMedia struct {
	ID             uint      `gorm:"primaryKey" json:"id"`
	Name           string    `gorm:"not null" json:"name" valid:"required"`
	SocialMediaURL string    `gorm:"not null" json:"social_media_url" valid:"required"`
	UserID         uint      `gorm:"not null" json:"user_id" valid:"required"`
	CreatedAt      time.Time `gorm:"autoCreateTime" json:"created_at" time_format:"2006-01-02 15:04:05"`
	UpdatedAt      time.Time `gorm:"autoUpdateTime" json:"updated_at" time_format:"2006-01-02 15:04:05"`
	User           User      `gorm:"foreignKey:UserID;Constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
