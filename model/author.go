package model

import "time"

type Author struct {
	ID   int    `xorm:"id"`
	Name string `xorm:"name"`
	// Books     []Book    `xorm:"books"` // has many
	CreatedAt time.Time `xorm:"created_at"`
	UpdatedAt time.Time `xorm:"updated_at"`
}

type AuthorFromJson struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	// Books []Book `json:"books"`
}
