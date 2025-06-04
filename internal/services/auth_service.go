package services

import (
    "fast-line/internal/models"
    "fast-line/internal/repository"
    "fast-line/internal/utils"
    "errors"
)

type AuthService struct {
    userRepo  *repository.UserRepository
    jwtSecret string
}

func NewAuthService(userRepo *repository.UserRepository, jwtSecret string) *AuthService {
    return &AuthService{
        userRepo:  userRepo,
        jwtSecret: jwtSecret,
    }
}

func (s *AuthService) Register(req *models.CreateUserRequest) (*models.AuthResponse, error) {
    // Check if user already exists
    existingUser, _ := s.userRepo.GetUserByEmail(req.Email)
    if existingUser != nil {
        return nil, errors.New("user with this email already exists")
    }
    
    // Hash password
    hashedPassword, err := utils.HashPassword(req.Password)
    if err != nil {
        return nil, err
    }
    
    // Create user
    user, err := s.userRepo.CreateUser(req, hashedPassword)
    if err != nil {
        return nil, err
    }
    
    // Generate token
    token, err := utils.GenerateToken(user.ID, user.Email, s.jwtSecret)
    if err != nil {
        return nil, err
    }
    
    return &models.AuthResponse{
        User:  user,
        Token: token,
    }, nil
}

func (s *AuthService) Login(req *models.LoginRequest) (*models.AuthResponse, error) {
    // Get user by email
    user, err := s.userRepo.GetUserByEmail(req.Email)
    if err != nil {
        return nil, errors.New("invalid email or password")
    }
    
    // Check password
    if err := utils.CheckPassword(req.Password, user.Password); err != nil {
        return nil, errors.New("invalid email or password")
    }
    
    // Generate token
    token, err := utils.GenerateToken(user.ID, user.Email, s.jwtSecret)
    if err != nil {
        return nil, err
    }
    
    // Remove password hash from response
    user.Password = ""
    
    return &models.AuthResponse{
        User:  user,
        Token: token,
    }, nil
}
