package model

import "time"

type Book struct {
	ID   int    `xorm:"id"`
	Name string `xorm:"name"`
	// Author    Author    `xorm:"author"` // has one
	CreatedAt time.Time `xorm:"created_at"`
	UpdatedAt time.Time `xorm:"updated_at"`
}

type BookFromJson struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	// Author Author `json:"author"`
}
