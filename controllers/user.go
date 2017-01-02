package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mefellows/follow/config"
	"github.com/mefellows/follow/models"
)

// UserController wraps up student related routes.
type UserController struct {
	config *config.Config
}

// NewUserController returns a new controller.
func NewUserController(config *config.Config) *UserController {
	return &UserController{
		config: config,
	}
}

// Get all users by user Route.
func (s *UserController) Get(c *gin.Context) {
	var users []models.User
	s.config.DB.Find(&users)

	var err error
	if err == nil {
		c.JSON(200, users)
	} else {
		log.Println("[DEBUG] unable to list users")
		c.JSON(http.StatusNotFound, gin.H{"status": "file not found"})
	}
}

// GetSingle gets a user by id.
func (s *UserController) GetSingle(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	var users models.User
	s.config.DB.Where("id = ?", id).Last(&users)

	var err error
	if err == nil {
		c.JSON(200, users)
	} else {
		log.Println("[DEBUG] unable to list user:", id)
		c.JSON(http.StatusNotFound, gin.H{"status": "file not found"})
	}
}
