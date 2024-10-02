package services

import (
	"github.com/biyoba1/effective_mobile/internal/models"
	"github.com/biyoba1/effective_mobile/internal/repository"
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

//delete {Song}
//update {SondDetail}
//create {Song}
//GetSongs {Song}
//GetSongLyrics {SongLyric}

type Service struct {
	SongService
	SongDetailsService
	SongLyricService
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		SongService:        NewSongService(repos.SongService),
		SongLyricService:   NewSongLyricsService(repos.SongLyricService),
		SongDetailsService: NewSongDetailsService(repos.SongDetailsService),
	}
}
