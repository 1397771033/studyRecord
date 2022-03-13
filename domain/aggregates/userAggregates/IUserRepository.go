package userAggregates

type IUserRepository interface {
	Add(user *User)
	Update(user *User)
}
