package models

import "gorm.io/gorm"

type Song struct {
	gorm.Model
	GroupName  string      `gorm:"not null"`
	SongName   string      `gorm:"not null"`
	SongDetail SongDetail  `gorm:"foreignkey:SongID;references:ID"`
	Lyrics     []SongLyric `gorm:"many2many:songs_lyrics"`
}
