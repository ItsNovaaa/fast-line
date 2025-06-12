package repository

import (
	"time"
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
        Phone:        req.Phone,
    })
    if err != nil {
        return nil, err
    }
    
    return &models.User{
        ID:        user.ID,
        Email:     user.Email,
        FirstName: user.FirstName,
        LastName:  user.LastName,
        Phone:     &user.Phone,
		
    }, nil
}

func (r *UserRepository) GetUserByEmail(email string) (*models.User, error) {
    user, err := r.queries.GetUserByEmail(context.Background(), email)

	if err != nil {
		return nil, err
	}
	
	var updatedAt time.Time
	if user.UpdatedAt.Valid {
		updatedAt = user.UpdatedAt.Time
	}

	var createdAt time.Time
	if user.CreatedAt.Valid {
		createdAt = user.CreatedAt.Time
	}

	var isActive bool
	if user.IsActive.Valid {
		isActive = user.IsActive.Bool
	}

	// Buat struct User dengan data yang sudah diperbaiki
	return &models.User{
		ID:        user.ID,
		Email:     user.Email,
		Password: user.Password, // Sebaiknya jangan kembalikan password
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Phone:     &user.Phone, // BENAR: Ambil alamat untuk membuat pointer
		IsActive:  isActive,
		CreatedAt: createdAt, // Gunakan variabel yang sudah disiapkan
		UpdatedAt: updatedAt, // Gunakan variabel yang sudah disiapkan
	}, nil

}