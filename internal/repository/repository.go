package repository

import (
	"github.com/biyoba1/effective_mobile/internal/models"
	"gorm.io/gorm"
)

type SongService interface {
	CreateSong(song models.Song) (int, error)
	GetSong(songId int) (models.Song, error)
	GetSongs(filter map[string]string, pagination *models.Pagination) ([]models.Song, error)
	UpdateSong(songId int, song models.Song) error
	DeleteSong(songId int) error
}

type SongDetailsService interface {
	CreateDetails(songId int, details models.SongDetail) (int, error)
	GetSongDetails(songId int) (models.SongDetail, error)
	UpdateSongDetails(songId int, details models.SongDetail) error
	DeleteDetails(songId int) error
}

type SongLyricService interface {
	CreateLyric(songId int, lyric models.SongLyric) (int, error)
	GetSongLyrics(songId int, pagination *models.Pagination) ([]models.SongLyric, error)
	UpdateLyric(songId int, lyricId int, lyric models.SongLyric) error
	DeleteLyric(songId int, lyricId int) error
}

type Repository struct {
	SongService
	SongDetailsService
	SongLyricService
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		SongService:        NewSongsPostgres(db),
		SongLyricService:   NewLyricsPostgres(db),
		SongDetailsService: NewDetailsPostgres(db),
	}
}
