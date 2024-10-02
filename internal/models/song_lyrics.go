package models

import "gorm.io/gorm"

type SongLyric struct {
	gorm.Model
	SongID uint `gorm:"foreignkey:SongID;references:ID"`
	Lyric  string
}
