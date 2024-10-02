package handler

import (
	"github.com/biyoba1/effective_mobile/internal/models"
	"github.com/biyoba1/effective_mobile/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) CreateSong(c *gin.Context) {
	var input models.Song
	err := c.Bind(&input)

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid body")
		return
	}

	id, err := h.services.CreateSong(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, id)
}

func (h *Handler) GetSong(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id")
		return
	}
	song, err := h.services.SongService.GetSong(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, song)
}

func (h *Handler) GetSongs(c *gin.Context) {
	var filter map[string]string
	err := c.BindJSON(&filter)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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

	songs, err := h.services.SongService.GetSongs(filter, &pagination)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, songs)
}

func (h *Handler) UpdateSong(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id")
		return
	}

	var input models.Song
	err = c.Bind(&input)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid body")
		return
	}

	err = h.services.SongService.UpdateSong(id, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}

func (h *Handler) DeleteSong(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id")
		return
	}
	err = h.services.SongService.DeleteSong(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}
