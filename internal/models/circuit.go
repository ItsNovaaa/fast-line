package models

import (
    "time"
	"github.com/google/uuid"
)

type Circuit struct {
    ID          uuid.UUID `json:"id" db:"id"`
    Name        string    `json:"name" db:"name"`
    CircuitName string    `json:"circuit_name" db:"circuit_name"`
    StartDate   time.Time `json:"start_date" db:"start_date"`
    EndDate     time.Time `json:"end_date" db:"end_date"`
    Status      int32      `json:"status" db:"status"`
    CreatedAt   time.Time `json:"created_at" db:"created_at"`
    UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

type CreateCircuitRequest struct {
    Name        string    `json:"name" binding:"required,min=3,max=255"`
    CircuitName string    `json:"circuit_name" binding:"required,min=3,max=255"`
    StartDate   time.Time `json:"start_date" binding:"required"`
    EndDate     time.Time `json:"end_date" binding:"required"`
    Status      int32      `json:"status" binding:"required"`
}

type UpdateCircuitRequest struct {
    // ID          uuid.UUID `json:"id" binding:"required"`
    Name        string    `json:"name" binding:"required,min=3,max=255"`
    CircuitName string    `json:"circuit_name" binding:"required,min=3,max=255"`
    StartDate   time.Time `json:"start_date" binding:"required"`
    EndDate     time.Time `json:"end_date" binding:"required"`
    Status      int32      `json:"status" binding:"required"`

}