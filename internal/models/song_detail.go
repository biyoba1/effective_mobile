package models

import "gorm.io/gorm"

type SongDetail struct {
	gorm.Model
	SongID      uint `gorm:"primary_key;foreignkey:SongID;references:ID"`
	ReleaseDate string
	Link        string
}
