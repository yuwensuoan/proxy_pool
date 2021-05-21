package models

import (
	"gorm.io/gorm"
	"time"
)

type BaseModel struct {
	// gorm.Model
	ID	uint `gorm:"primarykey"`
}

type DatetimeField struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
