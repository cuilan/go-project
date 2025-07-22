package models

// User - 数据库中的用户模型
type User struct {
	ID   int64  `sql:"primaryKey"`
	Name string `sql:"size:255;not null"`
}
