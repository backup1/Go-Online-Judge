package models

import (
	"time"
)

type User struct {
	ID        int       `gorm:"column:id" json:"uid"`
	Username  string    `gorm:"column:username" json:"username"`
	NickName  string    `gorm:"column:nickname" json:"nickname"`
	Password  string    `gorm:"column:password" json:"password"`
	Token     string    `gorm:"column:token" json:"token"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	LastLogin time.Time `gorm:"column:last_login" json:"last_login"`
}
