package model

import (
	"time"
)

type User struct {
	Model
	Username string
	Password string
	Tel      string
	Email    string
	Gender   bool
	Birth    time.Time
}

func (User) TableName() string {
	return "user"
}
