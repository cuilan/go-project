package models

import "gorm.io/gorm"

// User - 数据库中的用户模型
type User struct {
	gorm.Model        // GORM 的基础模型，包含 ID, CreatedAt, UpdatedAt, DeletedAt
	Name       string `gorm:"size:255;not null"`
	Email      string `gorm:"size:255;unique;not null"`
}
