package userAggregates

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

type User struct {
	CreatedAt time.Time

	ID      uuid.UUID `gorm:"primaryKey;type:varchar(36)"`
	Account Account   `gorm:"foreignKey:UserID"`
}

func Create(username string, password string, nickName string) *User {
	userId := uuid.NewV4()
	user := User{
		ID: userId,
		Account: Account{
			ID:       uuid.NewV4(),
			UserName: username,
			Password: password,
			NickName: nickName,
		},
	}
	return &user
}
