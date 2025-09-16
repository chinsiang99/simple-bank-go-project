package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/chinsiang99/simple-bank-go-project/internal/models"
	"github.com/chinsiang99/simple-bank-go-project/internal/services"
	"github.com/gin-gonic/gin"
)

// -------------------------------
// Interface
// -------------------------------

type IUserHandler interface {
	CreateUser(c *gin.Context)
	GetUserByID(c *gin.Context)
	GetAllUsers(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
}

// -------------------------------
// Implementation
// -------------------------------

type userHandler struct {
	service services.IUserService
}

// Constructor (DI)
func NewUserHandler(service services.IUserService) IUserHandler {
	return &userHandler{service: service}
}

// -------------------------------
// Handler methods with Swagger Docs
// -------------------------------

// CreateUser godoc
// @Summary Create a new user
// @Description Create a new user with name, email, and password
// @Tags users
// @Accept json
// @Produce json
// @Param user body models.CreateUserRequest true "User info"
// @Success 201 {object} models.UserResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /users [post]
func (h *userHandler) CreateUser(c *gin.Context) {
	var req models.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password, // TODO: hash this before saving
	}

	if err := h.service.CreateUser(&user); err != nil {
		c.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create user"})
		return
	}

	resp := models.UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt.Format(time.RFC3339),
		UpdatedAt: user.UpdatedAt.Format(time.RFC3339),
	}

	c.JSON(http.StatusCreated, resp)
}

// GetUserByID godoc
// @Summary Get a user by ID
// @Description Get details of a user by their ID
// @Tags users
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} models.UserResponse
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /users/{id} [get]
func (h *userHandler) GetUserByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	user, err := h.service.GetUserByID(uint(id))
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	resp := models.UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt.Format(time.RFC3339),
		UpdatedAt: user.UpdatedAt.Format(time.RFC3339),
	}

	c.JSON(http.StatusOK, resp)
}

// GetAllUsers godoc
// @Summary Get all users
// @Description Get a list of all users
// @Tags users
// @Produce json
// @Success 200 {array} models.UserResponse
// @Failure 500 {object} map[string]string
// @Router /users [get]
func (h *userHandler) GetAllUsers(c *gin.Context) {
	users, err := h.service.GetAllUsers()
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch users"})
		return
	}

	var resp []models.UserResponse
	for _, user := range users {
		resp = append(resp, models.UserResponse{
			ID:        user.ID,
			Name:      user.Name,
			Email:     user.Email,
			CreatedAt: user.CreatedAt.Format(time.RFC3339),
			UpdatedAt: user.UpdatedAt.Format(time.RFC3339),
		})
	}

	c.JSON(http.StatusOK, resp)
}

// UpdateUser godoc
// @Summary Update a user
// @Description Update user details by ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param user body models.UpdateUserRequest true "Updated user info"
// @Success 200 {object} models.UserResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /users/{id} [put]
func (h *userHandler) UpdateUser(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	var req models.UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{
		ID:       uint(id),
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}

	if err := h.service.UpdateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update user"})
		return
	}

	resp := models.UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt.Format(time.RFC3339),
		UpdatedAt: user.UpdatedAt.Format(time.RFC3339),
	}

	c.JSON(http.StatusOK, resp)
}

// DeleteUser godoc
// @Summary Delete a user
// @Description Delete a user by their ID
// @Tags users
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /users/{id} [delete]
func (h *userHandler) DeleteUser(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	if err := h.service.DeleteUser(uint(id)); err != nil {
		c.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user deleted"})
}
