package business

type User struct {
	Id        int
	Firstname string `gorm:"column:first_name"`
	Lastname  string `gorm:"column:last_name"`
	Balance   float32
}
