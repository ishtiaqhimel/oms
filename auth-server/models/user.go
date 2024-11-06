package models

// User represents the user for this application
// swagger: model
type User struct {
	// the id for the user
	//
	// required: false
	ID uint `gorm:"primaryKey" json:"-"`
	// the name for the user, must be unique
	//
	// required: true
	// unique: true
	Username string `gorm:"unique;not null"`
	// the password for the user
	//
	// required: true
	Password string `gorm:"not null"`
}
