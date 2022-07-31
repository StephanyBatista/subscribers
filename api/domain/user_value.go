package domain

type UserValue struct {
	Id   string `gorm:"not null;size:25"`
	Name string `gorm:"size:100; not null"`
}
