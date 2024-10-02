package repository

import (
	"errors"
	"github.com/biyoba1/effective_mobile/initializer"
	"github.com/biyoba1/effective_mobile/internal/models"
	"gorm.io/gorm"
)

type SongsDetailsPostgres struct {
	db *gorm.DB
}

func NewDetailsPostgres(db *gorm.DB) *SongsDetailsPostgres {
	return &SongsDetailsPostgres{db: db}
}

func (s *SongsDetailsPostgres) CreateDetails(songId int, details models.SongDetail) (int, error) {
	details.SongID = uint(songId)

	var existingDetails models.SongDetail
	err := initializer.DB.Where("song_id = ?", songId).First(&existingDetails).Error
	if err == nil {
		return 0, errors.New("запись для этой песни уже существует")
	}

	id := initializer.DB.Create(&details)
	if id.Error != nil {
		return 0, id.Error
	}

	song := models.Song{}
	err = initializer.DB.First(&song, songId).Error
	if err != nil {
		return 0, err
	}

	err = initializer.DB.Model(&song).Association("SongDetail").Append(&details)
	if err != nil {
		return 0, err
	}

	return int(id.RowsAffected), nil
}

func (s *SongsDetailsPostgres) GetSongDetails(songId int) (models.SongDetail, error) {
	var details models.SongDetail
	db := initializer.DB.Model(&models.SongDetail{})

	db = db.Where("song_id = ?", songId)

	err := db.Find(&details).Error

	if err != nil {
		return models.SongDetail{}, err
	}

	return details, nil
}

func (s *SongsDetailsPostgres) UpdateSongDetails(songId int, details models.SongDetail) error {
	db := initializer.DB.Model(&models.SongDetail{})
	err := db.Where("id = ?", songId).Updates(details).Error
	if err != nil {
		return err
	}
	return nil
}

func (s *SongsDetailsPostgres) DeleteDetails(songId int) error {
	db := s.db.Where("song_id = ?", songId).
		Delete(&models.SongDetail{})

	return db.Error
}
