package model

import (
	"gorm.io/plugin/soft_delete"
	"time"
)

type Model struct {
	Id        uint64
	CreatedAt time.Time             `gorm:"autoCreateTime"`
	UpdatedAt time.Time             `gorm:"autoUpdateTime"`
	IsDeleted soft_delete.DeletedAt `gorm:"softDelete:flag"`
}
