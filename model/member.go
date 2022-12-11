package model

import (
	"time"
)

type Member struct {
	Model
	Username string
	Password string
	Tel      string
	Email    string
	Gender   bool
	Birth    time.Time
}

func (Member) TableName() string {
	return "member"
}
