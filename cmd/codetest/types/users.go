package types

import (
	"time"
)

// hello world
type World struct {
	Hello string `json:"hello"`
}

// Defines values all entities should have.
type Model struct {
	ID        uint       `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

// Defines a user.
type User struct {
	Model
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

// roles in a company.
type Role struct {
	Model
	Name string `json:"name"`
}

// users can have one or more roles.
type UserRoles struct {
	Model
	Role
	Users []User `json:"users"`
}

// a company with it's members.
type Company struct {
	Model
	Name    string      `gorm:"column:company_name" json:"company"`
	Members []UserRoles `json:"members"`
}
