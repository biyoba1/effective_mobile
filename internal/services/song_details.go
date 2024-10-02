package services

import (
	"github.com/biyoba1/effective_mobile/internal/models"
	"github.com/biyoba1/effective_mobile/internal/repository"
)

type SongDetailsStruct struct {
	repo repository.SongDetailsService
}

func NewSongDetailsService(repo repository.SongDetailsService) *SongDetailsStruct {
	return &SongDetailsStruct{repo: repo}
}

func (s *SongDetailsStruct) GetSongDetails(songId int) (models.SongDetail, error) {
	return s.repo.GetSongDetails(songId)
}

func (s *SongDetailsStruct) UpdateSongDetails(songId int, details models.SongDetail) error {
	return s.repo.UpdateSongDetails(songId, details)
}

func (s *SongDetailsStruct) CreateDetails(songId int, details models.SongDetail) (int, error) {
	return s.repo.CreateDetails(songId, details)
}

func (s *SongDetailsStruct) DeleteDetails(songId int) error {
	return s.repo.DeleteDetails(songId)
}
