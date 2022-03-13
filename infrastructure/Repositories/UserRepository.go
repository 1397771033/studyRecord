package Repositories

import (
	"StudentRecord/domain/aggregates/userAggregates"
	"StudentRecord/infrastructure/DataBase"
	_ "gorm.io/gorm"
)

var db = DataBase.NewDb()

type UserRepository struct {
}

func (r *UserRepository) Add(user *userAggregates.User) {
	db.Create(user)
}
func (r *UserRepository) Update(user *userAggregates.User) {

}
