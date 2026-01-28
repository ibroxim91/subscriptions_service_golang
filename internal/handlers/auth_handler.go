package handlers

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

type AuthHandler struct{}

func NewAuthHandler() *AuthHandler {
    return &AuthHandler{}
}

// Login godoc
// @Summary Get auth token
// @Tags auth
// @Accept json
// @Produce json
// @Param credentials body LoginRequest true "Login credentials"
// @Success 200 {object} TokenResponse
// @Failure 401 {object} ErrorResponse
// @Router /login [post]
func (h *AuthHandler) Login(c *gin.Context) {
    var req LoginRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
        return
    }

    // Static login/password check
    if req.Username == "admin" && req.Password == "password" {
        c.JSON(http.StatusOK, TokenResponse{Token: "test-token"})
        return
    }

    c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
}

// LoginRequest represents login payload
type LoginRequest struct {
    Username string `json:"username" example:"admin"`
    Password string `json:"password" example:"password"`
}

// TokenResponse represents token response
type TokenResponse struct {
    Token string `json:"token" example:"test-token"`
}

// ErrorResponse represents error message
type ErrorResponse struct {
    Error string `json:"error" example:"invalid credentials"`
}
