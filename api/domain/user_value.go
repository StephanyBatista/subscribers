package domain

type UserValue struct {
	Id   string `gorm:"size:25;"`
	Name string `gorm:"size:100;"`
}
