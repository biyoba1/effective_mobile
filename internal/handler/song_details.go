package handler

import (
	"github.com/biyoba1/effective_mobile/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) CreateDetails(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id")
		return
	}

	var input models.SongDetail
	err = c.Bind(&input)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid body")
		return
	}

	detailsId, err := h.services.SongDetailsService.CreateDetails(id, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, detailsId)
}

func (h *Handler) GetDetails(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id")
		return
	}

	songs, err := h.services.SongDetailsService.GetSongDetails(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, songs)
}

func (h *Handler) DeleteDetails(c *gin.Context) {
	songId, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid song_id")
		return
	}

	err = h.services.SongDetailsService.DeleteDetails(songId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}

func (h *Handler) UpdateDetails(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id")
		return
	}

	var input models.SongDetail
	err = c.Bind(&input)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid body")
		return
	}

	err = h.services.SongDetailsService.UpdateSongDetails(id, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}
