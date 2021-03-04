package model

import "time"

/*
User Domain Model
*/
type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

/*
Users Domain Model
*/
type Users []User
