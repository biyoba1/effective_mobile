package repository

import (
	"errors"
	"github.com/biyoba1/effective_mobile/initializer"
	"github.com/biyoba1/effective_mobile/internal/models"
	"gorm.io/gorm"
)

type SongsLyricsPostgres struct {
	db *gorm.DB
}

func NewLyricsPostgres(db *gorm.DB) *SongsLyricsPostgres {
	return &SongsLyricsPostgres{db: db}
}

func (s *SongsLyricsPostgres) CreateLyric(songId int, lyric models.SongLyric) (int, error) {
	song := models.Song{}
	err := initializer.DB.First(&song, songId).Error
	if err != nil {
		return 0, errors.New("песня с таким айди не существует")
	}

	lyric.SongID = uint(songId)
	id := initializer.DB.Create(&lyric)
	if id.Error != nil {
		return 0, id.Error
	}

	err = initializer.DB.Model(&song).Association("Lyrics").Append(&lyric)
	if err != nil {
		return 0, err
	}

	return int(id.RowsAffected), nil
}

func (s *SongsLyricsPostgres) GetSongLyrics(songId int, pagination *models.Pagination) ([]models.SongLyric, error) {
	var lyrics []models.SongLyric
	db := initializer.DB.Model(&models.SongLyric{})

	db = db.Where("song_id = ?", songId)
	db = db.Limit(pagination.Limit).Offset(pagination.Offset)

	err := db.Find(&lyrics).Error

	if err != nil {
		return nil, err
	}

	return lyrics, nil
}

func (s *SongsLyricsPostgres) UpdateLyric(songId int, lyricId int, lyric models.SongLyric) error {
	db := s.db.Model(&models.SongLyric{}).
		Where("song_id = ? AND id = ?", songId, lyricId).
		Updates(lyric)

	return db.Error
}

func (s *SongsLyricsPostgres) DeleteLyric(songId int, lyricId int) error {
	db := s.db.Where("song_id = ? AND id = ?", songId, lyricId).
		Delete(&models.SongLyric{})

	return db.Error
}
