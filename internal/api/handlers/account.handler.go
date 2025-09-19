package handlers

import (
	"database/sql"
	"net/http"

	"github.com/chinsiang99/simple-bank-go-project/internal/dto"
	"github.com/chinsiang99/simple-bank-go-project/internal/services"
	"github.com/gin-gonic/gin"
)

// -------------------------------
// Interface
// -------------------------------

type IAccountHandler interface {
	CreateAccount(c *gin.Context)
	GetAccount(c *gin.Context)
	// GetUserByID(c *gin.Context)
	// GetAllUsers(c *gin.Context)
	// UpdateUser(c *gin.Context)
	// DeleteUser(c *gin.Context)
}

// -------------------------------
// Implementation
// -------------------------------

type accountHandler struct {
	service services.IAccountService
}

// Constructor (DI)
func NewAccountHandler(service services.IAccountService) IAccountHandler {
	return &accountHandler{service: service}
}

// -------------------------------
// Handler methods with Swagger Docs
// -------------------------------

// CreateAccount godoc
// @Summary Create a new account
// @Description Create a new bank account with owner and currency
// @Tags accounts
// @Accept json
// @Produce json
// @Param account body dto.CreateAccountRequest true "Account info"
// @Success 201 {object} db.Account
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /accounts [post]
func (handler *accountHandler) CreateAccount(ctx *gin.Context) {
	var req dto.CreateAccountRequest
	// var req models.CreateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	account, err := handler.service.CreateAccount(ctx, req)
	if err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create account"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"data": account})
}

// GetAccount godoc
// @Summary Get account by ID
// @Description Retrieve account details by account ID
// @Tags accounts
// @Accept json
// @Produce json
// @Param id path int true "Account ID"
// @Success 200 {object} db.Account
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /accounts/{id} [get]
func (handler *accountHandler) GetAccount(ctx *gin.Context) {
	var req dto.GetAccountRequest
	// var req models.CreateUserRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	account, err := handler.service.GetAccount(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "account not found"})
			return
		}
		ctx.Error(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get account"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"data": account})
}

// // GetUserByID godoc
// // @Summary Get a user by ID
// // @Description Get details of a user by their ID
// // @Tags users
// // @Produce json
// // @Param id path int true "User ID"
// // @Success 200 {object} models.UserResponse
// // @Failure 400 {object} map[string]string
// // @Failure 404 {object} map[string]string
// // @Router /users/{id} [get]
// func (h *userHandler) GetUserByID(c *gin.Context) {
// 	idStr := c.Param("id")
// 	id, err := strconv.Atoi(idStr)
// 	if err != nil {
// 		c.Error(err)
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
// 		return
// 	}

// 	user, err := h.service.GetUserByID(uint(id))
// 	if err != nil {
// 		c.Error(err)
// 		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
// 		return
// 	}

// 	resp := models.UserResponse{
// 		ID:        user.ID,
// 		Name:      user.Name,
// 		Email:     user.Email,
// 		CreatedAt: user.CreatedAt.Format(time.RFC3339),
// 		UpdatedAt: user.UpdatedAt.Format(time.RFC3339),
// 	}

// 	c.JSON(http.StatusOK, resp)
// }

// // GetAllUsers godoc
// // @Summary Get all users
// // @Description Get a list of all users
// // @Tags users
// // @Produce json
// // @Success 200 {array} models.UserResponse
// // @Failure 500 {object} map[string]string
// // @Router /users [get]
// func (h *userHandler) GetAllUsers(c *gin.Context) {
// 	users, err := h.service.GetAllUsers()
// 	if err != nil {
// 		c.Error(err)
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch users"})
// 		return
// 	}

// 	var resp []models.UserResponse
// 	for _, user := range users {
// 		resp = append(resp, models.UserResponse{
// 			ID:        user.ID,
// 			Name:      user.Name,
// 			Email:     user.Email,
// 			CreatedAt: user.CreatedAt.Format(time.RFC3339),
// 			UpdatedAt: user.UpdatedAt.Format(time.RFC3339),
// 		})
// 	}

// 	c.JSON(http.StatusOK, resp)
// }

// // UpdateUser godoc
// // @Summary Update a user
// // @Description Update user details by ID
// // @Tags users
// // @Accept json
// // @Produce json
// // @Param id path int true "User ID"
// // @Param user body models.UpdateUserRequest true "Updated user info"
// // @Success 200 {object} models.UserResponse
// // @Failure 400 {object} map[string]string
// // @Failure 500 {object} map[string]string
// // @Router /users/{id} [put]
// func (h *userHandler) UpdateUser(c *gin.Context) {
// 	idStr := c.Param("id")
// 	id, err := strconv.Atoi(idStr)
// 	if err != nil {
// 		c.Error(err)
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
// 		return
// 	}

// 	var req models.UpdateUserRequest
// 	if err := c.ShouldBindJSON(&req); err != nil {
// 		c.Error(err)
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	user := models.User{
// 		ID:       uint(id),
// 		Name:     req.Name,
// 		Email:    req.Email,
// 		Password: req.Password,
// 	}

// 	if err := h.service.UpdateUser(&user); err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update user"})
// 		return
// 	}

// 	resp := models.UserResponse{
// 		ID:        user.ID,
// 		Name:      user.Name,
// 		Email:     user.Email,
// 		CreatedAt: user.CreatedAt.Format(time.RFC3339),
// 		UpdatedAt: user.UpdatedAt.Format(time.RFC3339),
// 	}

// 	c.JSON(http.StatusOK, resp)
// }

// // DeleteUser godoc
// // @Summary Delete a user
// // @Description Delete a user by their ID
// // @Tags users
// // @Produce json
// // @Param id path int true "User ID"
// // @Success 200 {object} map[string]string
// // @Failure 400 {object} map[string]string
// // @Failure 500 {object} map[string]string
// // @Router /users/{id} [delete]
// func (h *userHandler) DeleteUser(c *gin.Context) {
// 	idStr := c.Param("id")
// 	id, err := strconv.Atoi(idStr)
// 	if err != nil {
// 		c.Error(err)
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
// 		return
// 	}

// 	if err := h.service.DeleteUser(uint(id)); err != nil {
// 		c.Error(err)
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete user"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"message": "user deleted"})
// }
