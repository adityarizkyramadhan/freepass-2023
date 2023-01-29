package domain

import (
	"time"

	"gorm.io/gorm"
)

type Mahasiswa struct {
	ID           uint           `gorm:"primarykey;autoIncrement" json:"id"`
	CreatedAt    time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt    time.Time      `gorm:"autoUpdateTime:milli" json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	Nim          string         `gorm:"unique" json:"nim"`
	Password     string         `gorm:"not null" json:"-"`
	Nama         string         `gorm:"not null" json:"nama"`
	ProgramStudi string         `gorm:"not null" json:"program_studi"`
	Fakultas     string         `gorm:"not null" json:"fakultas"`
	MataKuliahs  []MataKuliah   `gorm:"many2many:mahasiswa_matkuls"`
}
