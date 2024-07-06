package models

import "time"

type User struct {
	ID 				uint  			`json:"id" gorm:"primaryKey"`
	Name 			string			`json:"name"`
	Surname 	string 			`json:"surname"`
	CreatedAt	time.Time		`json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
}