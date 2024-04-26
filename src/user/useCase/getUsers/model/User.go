package model

type User struct {
	Id        int     `json:"user_id"`
	Firstname string  `json:"first_name" gorm:"column:first_name"`
	Lastname  string  `json:"last_name" gorm:"column:last_name"`
	Balance   float32 `json:"balance"`
}
