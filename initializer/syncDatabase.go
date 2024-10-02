package initializer

import "github.com/biyoba1/effective_mobile/internal/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.Song{}, &models.SongLyric{}, &models.SongDetail{})
}
