package handler

import (
	"github.com/biyoba1/effective_mobile/internal/models"
	"github.com/biyoba1/effective_mobile/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) CreateLyric(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id")
		return
	}

	var input models.SongLyric
	err = c.Bind(&input)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid body")
		return
	}

	lyricId, err := h.services.SongLyricService.CreateLyric(id, input)
	if err != nil {
		if err.Error() == "песня с таким айди не существует" {
			newErrorResponse(c, http.StatusNotFound, err.Error())
		} else {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
		}
		return
	}

	c.JSON(http.StatusCreated, lyricId)
}

func (h *Handler) GetLyrics(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id")
		return
	}

	page, _ := strconv.Atoi(c.Param("page"))

	err = middleware.PageValidate(page)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid page parameters")
		return
	}

	pagination := models.Pagination{
		Limit:  5,
		Offset: (page - 1) * 5,
		Page:   page,
	}

	songs, err := h.services.SongLyricService.GetSongLyrics(id, &pagination)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, songs)
}

func (h *Handler) UpdateLyric(c *gin.Context) {
	songId, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid song_id")
		return
	}

	lyricId, err := strconv.Atoi(c.Query("lyric_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id")
		return
	}

	var input models.SongLyric
	err = c.Bind(&input)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid body")
		return
	}

	err = h.services.SongLyricService.UpdateLyric(songId, lyricId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}

func (h *Handler) DeleteLyric(c *gin.Context) {
	songId, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid song_id")
		return
	}

	lyricId, err := strconv.Atoi(c.Query("lyric_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id")
		return
	}

	err = h.services.SongLyricService.DeleteLyric(songId, lyricId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}
