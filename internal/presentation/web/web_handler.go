package web

import (
	"github.com/FindZach/TextSnapz/internal/application"
	"github.com/FindZach/TextSnapz/internal/domain"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

// TextSnapWebHandler handles web UI requests.
type TextSnapWebHandler struct {
	service *application.TextSnapzService
	tmpl    *template.Template
}

// NewTextSnapWebHandler creates a new web handler instance.
func NewTextSnapWebHandler(service *application.TextSnapzService) (*TextSnapWebHandler, error) {
	tmpl, err := template.ParseFiles("templates/index.html", "templates/snap.html", "templates/raw.html") // Added raw.html
	if err != nil {
		return nil, err
	}
	return &TextSnapWebHandler{service: service, tmpl: tmpl}, nil
}

// RegisterRoutes sets up web routes for the UI.
func (h *TextSnapWebHandler) RegisterRoutes(r *gin.Engine) {
	r.Static("/static", "./static")
	r.GET("/", h.Home)
	r.GET("/s/:id", h.GetSnap)
	r.GET("/s/:id/raw", h.GetRawSnap) // route for raw content
}

// Home handles GET / to render the homepage.
func (h *TextSnapWebHandler) Home(c *gin.Context) {
	c.Header("Content-Type", "text/html")
	if err := h.tmpl.ExecuteTemplate(c.Writer, "index.html", nil); err != nil {
		c.String(http.StatusInternalServerError, "Error rendering template")
	}
}

// GetSnap handles GET /s/:id to render a snap.
func (h *TextSnapWebHandler) GetSnap(c *gin.Context) {
	id := c.Param("id")
	snap, err := h.service.GetSnap(c.Request.Context(), id)
	if err == domain.ErrSnapNotFound {
		c.String(http.StatusNotFound, "Snap not found")
		return
	}
	if err == domain.ErrSnapExpired {
		c.String(http.StatusGone, "Snap expired")
		return
	}
	if err != nil {
		c.String(http.StatusInternalServerError, "Error retrieving snap")
		return
	}
	c.Header("Content-Type", "text/html")
	if err := h.tmpl.ExecuteTemplate(c.Writer, "snap.html", snap); err != nil {
		c.String(http.StatusInternalServerError, "Error rendering template")
	}
}

// GetRawSnap handles GET /s/:id/raw to render raw snap content.
func (h *TextSnapWebHandler) GetRawSnap(c *gin.Context) {
	id := c.Param("id")
	snap, err := h.service.GetSnap(c.Request.Context(), id)
	if err == domain.ErrSnapNotFound {
		c.String(http.StatusNotFound, "Snap not found")
		return
	}
	if err == domain.ErrSnapExpired {
		c.String(http.StatusGone, "Snap expired")
		return
	}
	if err != nil {
		c.String(http.StatusInternalServerError, "Error retrieving snap")
		return
	}
	c.Header("Content-Type", "text/html")
	if err := h.tmpl.ExecuteTemplate(c.Writer, "raw.html", snap); err != nil {
		c.String(http.StatusInternalServerError, "Error rendering template")
	}
}
