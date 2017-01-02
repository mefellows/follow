package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mefellows/follow/config"
	"github.com/mefellows/follow/models"
)

// LocationController wraps up student related routes.
type LocationController struct {
	config *config.Config
}

// NewLocationController returns a new controller.
func NewLocationController(config *config.Config) *LocationController {
	return &LocationController{
		config: config,
	}
}

// Get all locations by user Route.
func (s *LocationController) Get(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	var locations []models.Location
	s.config.DB.Order("id DESC").Where("user_id = ?", id).Find(&locations)

	var err error
	if err == nil {
		c.JSON(200, locations)
	} else {
		log.Println("[DEBUG] unable to list locations for user:", id)
		c.JSON(http.StatusNotFound, gin.H{"status": "file not found"})
	}
}

// Create a location route.
func (s *LocationController) Create(c *gin.Context) {
	var location models.Location

	if c.Bind(&location) == nil {
		log.Println("[DEBUG] location", location)

		if err := s.config.DB.Debug().Save(&location).Error; err != nil {
			c.JSON(500, models.Error{Code: 500, Message: err.Error()})
			return
		}
		c.JSON(200, &location)
	} else {
		log.Println("[DEBUG] unable to create location")
		c.JSON(500, models.Error{Code: 500, Message: "unable to create location"})
	}
}

// DeleteItem route
func (s *LocationController) DeleteItem(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	s.config.DB.Delete(&models.Location{
		Model: models.Model{ID: id},
	})
	var err error
	if err == nil {
		c.JSON(200, gin.H{})
	} else {
		log.Println("[DEBUG] unable to find items for id", id)
		c.JSON(http.StatusNotFound, gin.H{"status": "file not found"})
	}
}
