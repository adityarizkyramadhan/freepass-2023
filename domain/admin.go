package domain

import (
	"time"

	"gorm.io/gorm"
)

type Admin struct {
	ID        uint           `gorm:"primarykey;autoIncrement" json:"id"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime:milli" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	Username  string         `gorm:"unique" json:"username"`
	Password  string         `gorm:"not null" json:"-"`
	Nama      string         `gorm:"not null" json:"nama"`
}
