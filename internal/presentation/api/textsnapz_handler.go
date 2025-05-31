package api

import (
	"github.com/FindZach/TextSnapz/internal/application"
	"github.com/FindZach/TextSnapz/internal/domain"
	"github.com/gin-gonic/gin"
	"net/http"
)

// TextSnapHandler handles HTTP requests for snaps.
type TextSnapHandler struct {
	service *application.TextSnapzService
}

// NewTextSnapHandler creates a new handler instance.
func NewTextSnapHandler(service *application.TextSnapzService) *TextSnapHandler {
	return &TextSnapHandler{service: service}
}

// CreateSnap handles POST /api/snaps, creating a new text snap with optional expiration.
func (h *TextSnapHandler) CreateSnap(c *gin.Context) {
	type Request struct {
		Title     string   `json:"title"`
		Content   string   `json:"content" binding:"required"`
		Syntax    string   `json:"syntax"`
		ExpiresIn *string  `json:"expires_in"` // e.g., "1h", "24h"
		Tags      []string `json:"tags"`
	}
	var req Request
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ip := c.ClientIP()
	id, err := h.service.CreateSnap(c.Request.Context(), req.Title, req.Content, req.Syntax, ip, req.ExpiresIn, req.Tags)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"id": id})
}

// GetSnap handles GET /api/snaps/:id to retrieve a snap.
func (h *TextSnapHandler) GetSnap(c *gin.Context) {
	id := c.Param("id")
	snap, err := h.service.GetSnap(c.Request.Context(), id)
	if err == domain.ErrSnapNotFound {
		c.JSON(http.StatusNotFound, gin.H{"error": "snap not found"})
		return
	}
	if err == domain.ErrSnapExpired {
		c.JSON(http.StatusGone, gin.H{"error": "snap expired"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, snap)
}
