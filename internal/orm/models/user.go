package models

// User - 数据库中的用户模型
type User struct {
	Username string `json:"username" sql:"size:255;not null"`
	Password string `json:"password" sql:"size:255;not null"`
	Id       int64  `json:"id" sql:"primaryKey"`
}

// 如需特殊表名，请自定义
// func (u *User) TableName() string {
// 	return "t_user"
// }
