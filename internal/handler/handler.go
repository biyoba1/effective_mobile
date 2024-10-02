package handler

import (
	"github.com/biyoba1/effective_mobile/internal/services"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *services.Service
}

func NewHandler(services *services.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()

	songs := router.Group("/songs")
	{
		songs.POST("/", h.CreateSong)        // создание новой песни
		songs.GET("/page/:page", h.GetSongs) // получение списка песен с фильтрацией и пагинацией
		songs.GET("/:id", h.GetSong)         // получение информации о конкретной песне
		songs.PUT("/:id", h.UpdateSong)      // изменение информации о конкретной песне
		songs.DELETE("/:id", h.DeleteSong)   // удаление конкретной песни

		lyrics := songs.Group("/lyrics")
		{
			lyrics.POST("/:id", h.CreateLyric)             // добавление нового куплета к песне
			lyrics.GET("/page/:page", h.GetLyrics)         // получение текста песни с пагинацией по куплетам
			lyrics.PUT("/:id/:lyric_id", h.UpdateLyric)    // изменение конкретного куплета
			lyrics.DELETE("/:id/:lyric_id", h.DeleteLyric) // удаление конкретного куплета
		}

		details := songs.Group("/details")
		{
			details.POST("/:id", h.CreateDetails)   // добавление информации о конкретной песне с деталями
			details.GET("/:id", h.GetDetails)       // получение информации о конкретной песне с деталями
			details.PUT("/:id", h.UpdateDetails)    // изменение информации о конкретной песне с деталями
			details.DELETE("/:id", h.DeleteDetails) // удаление информации о конкретной песне с деталями
		}
	}

	return router
}
