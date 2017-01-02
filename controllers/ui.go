package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mefellows/follow/config"
	"github.com/mefellows/follow/models"
)

// UIController wraps up student related routes.
type UIController struct {
	config *config.Config
}

// NewUIController returns a new controller.
func NewUIController(config *config.Config) *UIController {
	return &UIController{
		config: config,
	}
}

// Get UI Route.
func (s *UIController) Get(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	var locations []models.Location
	s.config.DB.Where("user_id = ?", id).Find(&locations)

	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title":            "Heatmap of current user",
		"LatStart":         locations[0].Location.Lat,
		"LngStart":         locations[0].Location.Lng,
		"Locations":        locations,
		"GoogleMapsAPIKey": s.config.GoogleMapsAPIKey,
	})
}
