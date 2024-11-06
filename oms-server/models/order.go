package models

type Order struct {
	ID          uint   `gorm:"primaryKey" json:"-"`
	Username    string `gorm:"not null"`
	ProductName string `gorm:"not null"`
}
