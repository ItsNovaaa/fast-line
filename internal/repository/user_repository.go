package repository

import (
	"context"
	"database/sql"
	"fast-line/internal/models"
	"fast-line/internal/repository/queries"
)

type UserRepository struct {
    queries *queries.Queries
}

func NewUserRepository(db *sql.DB) *UserRepository {
    return &UserRepository{
        queries: queries.New(db),
    }
}

func (r *UserRepository) CreateUser(req *models.CreateUserRequest, hashedPassword string) (*models.User, error) {
    user, err := r.queries.CreateUser(context.Background(), queries.CreateUserParams{
        Email:        req.Email,
        Password: hashedPassword,
        FirstName:    req.FirstName,
        LastName:     req.LastName,
        Phone:        &req.Phone,
    })
    if err != nil {
        return nil, err
    }
    
    return &models.User{
        ID:        user.ID,
        Email:     user.Email,
        FirstName: user.FirstName,
        LastName:  user.LastName,
        Phone:     user.Phone,
        IsActive:  user.IsActive,
        CreatedAt: user.CreatedAt,
        UpdatedAt: user.UpdatedAt,
    }, nil
}

func (r *UserRepository) GetUserByEmail(email string) (*models.User, error) {
    user, err := r.queries.GetUserByEmail(context.Background(), email)
    if err != nil {
        return nil, err
    }
    
    return &models.User{
        ID:           user.ID,
        Email:        user.Email,
        Password: user.Password,
        FirstName:    user.FirstName,
        LastName:     user.LastName,
        Phone:        user.Phone,
        IsActive:     user.IsActive,
        CreatedAt:    user.CreatedAt,
        UpdatedAt:    user.UpdatedAt,
    }, nil
}