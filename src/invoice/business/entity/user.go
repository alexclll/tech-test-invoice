package entity

type User struct {
	Id        int
	Firstname string `gorm:"column:first_name"`
	Lastname  string `gorm:"column:last_name"`
	Balance   float32
}

func (user *User) AddToBalance(amount float32) {
	user.Balance = user.Balance + amount
}
