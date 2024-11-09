package models

// Order represents an order placed by an user
// swagger: model
type Order struct {
	// the id for the order
	//
	// required: false
	ID uint `gorm:"primaryKey"`
	// the name of the user who placed the order
	//
	// required: true
	Username string `gorm:"not null"`
	// the product name that has been ordered
	//
	// required: true
	ProductName string `gorm:"not null"`
}
