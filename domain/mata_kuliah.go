package domain

import (
	"time"

	"gorm.io/gorm"
)

type MataKuliah struct {
	ID         uint           `gorm:"primarykey;autoIncrement" json:"id"`
	CreatedAt  time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt  time.Time      `gorm:"autoUpdateTime:milli" json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	Nama       string         `gorm:"not null" json:"nama"`
	Sks        int            `gorm:"not null" json:"sks"`
	Dosen      string         `gorm:"not null" json:"dosen"`
	Mahasiswas []Mahasiswa    `gorm:"many2many:mahasiswa_matkuls"`
}
