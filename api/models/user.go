package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Username  string    `gorm:"uniqueIndex" json:"username" valid:"required"`
	Email     string    `gorm:"uniqueIndex" json:"email" valid:"required,email"`
	Password  string    `gorm:"not null" json:"password" valid:"required"`
	Age       uint      `gorm:"not null" json:"age" valid:"required"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at" time_format:"2006-01-02 15:04:05"`
	UpdateAt  time.Time `gorm:"autoUpdateTime" json:"updated_at" time_format:"2006-01-02 15:04:05"`
}

func VerifyPassword(password, hashPass string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashPass), []byte(password))
}
