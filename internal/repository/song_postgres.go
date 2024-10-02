package repository

import (
	"github.com/biyoba1/effective_mobile/initializer"
	"github.com/biyoba1/effective_mobile/internal/models"
	"gorm.io/gorm"
)

type SongsPostgres struct {
	db *gorm.DB
}

func NewSongsPostgres(db *gorm.DB) *SongsPostgres {
	return &SongsPostgres{db: db}
}

func (s *SongsPostgres) CreateSong(song models.Song) (int, error) {
	id := initializer.DB.Create(&song)
	if id.Error != nil {
		return 0, id.Error
	}

	return int(id.RowsAffected), nil
}

func (s *SongsPostgres) GetSong(songId int) (models.Song, error) {
	var song models.Song

	err := initializer.DB.Preload("Lyrics").Preload("SongDetail").First(&song, songId).Error
	if err != nil {
		return models.Song{}, err
	}

	return song, nil
}

func (s *SongsPostgres) GetSongs(filter map[string]string, pagination *models.Pagination) ([]models.Song, error) {
	var songs []models.Song
	db := initializer.DB.Model(&models.Song{})

	//применение фильтров (сортировка по группе или песне)
	for key, value := range filter {
		if key == "GroupName" {
			db = db.Where("group_name = ?", value)
		} else {
			db = db.Where("song_name = ?", value)
		}
	}

	db = db.Limit(pagination.Limit).Offset(pagination.Offset)

	err := db.Preload("Lyrics").Preload("SongDetail").Find(&songs).Error

	if err != nil {
		return nil, err
	}

	return songs, nil
}

func (s *SongsPostgres) UpdateSong(songId int, song models.Song) error {
	db := initializer.DB.Model(&models.Song{})
	err := db.Where("id = ?", songId).Updates(song).Error
	if err != nil {
		return err
	}
	return nil
}

func (s *SongsPostgres) DeleteSong(songId int) error {
	var song models.Song
	err := initializer.DB.Delete(&song, songId).Error

	if err != nil {
		return err
	}

	return nil
}
