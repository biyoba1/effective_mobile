package services

import (
	"github.com/biyoba1/effective_mobile/internal/models"
	"github.com/biyoba1/effective_mobile/internal/repository"
)

type SongStruct struct {
	repo repository.SongService
}

func NewSongService(repo repository.SongService) *SongStruct {
	return &SongStruct{repo: repo}
}

func (s *SongStruct) CreateSong(song models.Song) (int, error) {
	return s.repo.CreateSong(song)
}

func (s *SongStruct) GetSong(songId int) (models.Song, error) {
	return s.repo.GetSong(songId)
}

func (s *SongStruct) GetSongs(filter map[string]string, pagination *models.Pagination) ([]models.Song, error) {
	return s.repo.GetSongs(filter, pagination)
}

func (s *SongStruct) UpdateSong(songId int, song models.Song) error {
	return s.repo.UpdateSong(songId, song)
}

func (s *SongStruct) DeleteSong(songId int) error {
	return s.repo.DeleteSong(songId)
}
