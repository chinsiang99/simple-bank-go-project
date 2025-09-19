package handlers

// "myapp/internal/services"

// type AuthHandler struct {
// 	authService services.IAuthService
// }

// func NewAuthHandler(services *services.ServiceManager) *AuthHandler {
// 	return &AuthHandler{authService: services.AuthService}
// }

// func (h *AuthHandler) Login(c *gin.Context) {
// 	var req struct {
// 		Username string `json:"username"`
// 		Password string `json:"password"`
// 	}
// 	if err := c.ShouldBindJSON(&req); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	token, err := h.authService.Login(c.Request.Context(), req.Username, req.Password)
// 	if err != nil {
// 		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"token": token})
// }
