package userAggregates

import (
	uuid "github.com/satori/go.uuid"
)

type Account struct {
	ID uuid.UUID `gorm:"primaryKey;type:varchar(36)"`
	// 用户名
	UserName string
	// 密码
	Password string
	// 昵称
	NickName string
	UserID   uuid.UUID `gorm:"type:varchar(36)"`
}
