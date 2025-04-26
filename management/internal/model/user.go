package model

import "time"

type User struct {
	ID        uint `gorm:"primarykey"`
	Name      string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (User) TableName() string {
	return "users" // Replace with your desired table name, e.g., "user_table"
}
