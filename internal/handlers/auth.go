package handlers

import (
    "net/http"
    "fast-line/internal/models"
    "fast-line/internal/services"
    "fast-line/internal/utils"
    
    "github.com/gin-gonic/gin"
)

type AuthHandler struct {
    authService *services.AuthService
}

func NewAuthHandler(authService *services.AuthService) *AuthHandler {
    return &AuthHandler{
        authService: authService,
    }
}

func (h *AuthHandler) Register(c *gin.Context) {
    var req models.CreateUserRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request data", err.Error())
        return
    }
    
    authResponse, err := h.authService.Register(&req)
    if err != nil {
        utils.ErrorResponse(c, http.StatusBadRequest, "Registration failed", err.Error())
        return
    }
    
    utils.SuccessResponse(c, http.StatusCreated, "User registered successfully", authResponse)
}

func (h *AuthHandler) Login(c *gin.Context) {
    var req models.LoginRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request data", err.Error())
        return
    }
    
    authResponse, err := h.authService.Login(&req)
    if err != nil {
        utils.ErrorResponse(c, http.StatusUnauthorized, "Login failed", err.Error())
        return
    }
    
    utils.SuccessResponse(c, http.StatusOK, "Login successful", authResponse)
}