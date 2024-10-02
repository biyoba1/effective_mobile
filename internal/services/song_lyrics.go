package services

import (
	"github.com/biyoba1/effective_mobile/internal/models"
	"github.com/biyoba1/effective_mobile/internal/repository"
)

type SongLyricsStruct struct {
	repo repository.SongLyricService
}

func NewSongLyricsService(repo repository.SongLyricService) *SongLyricsStruct {
	return &SongLyricsStruct{repo: repo}
}

func (s *SongLyricsStruct) CreateLyric(songId int, lyric models.SongLyric) (int, error) {
	return s.repo.CreateLyric(songId, lyric)
}

func (s *SongLyricsStruct) GetSongLyrics(songId int, pagination *models.Pagination) ([]models.SongLyric, error) {
	return s.repo.GetSongLyrics(songId, pagination)
}

func (s *SongLyricsStruct) UpdateLyric(songId int, lyricId int, lyric models.SongLyric) error {
	return s.repo.UpdateLyric(songId, lyricId, lyric)
}

func (s *SongLyricsStruct) DeleteLyric(songId int, lyricId int) error {
	return s.repo.DeleteLyric(songId, lyricId)
}
