package entity

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID             string `json:"id" gorm:"primaryKey"`
	Name           string `json:"name" binding:"required" gorm:"index"`
	Email          string `json:"email" binding:"required" gorm:"uniqueIndex"`
	PasswordDigest string `json:"-"`
}
