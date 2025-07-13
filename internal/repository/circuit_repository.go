package repository

import (
	"time"
	"context"
	"database/sql"
	"fast-line/internal/models"
	"fast-line/internal/repository/queries"
	"github.com/google/uuid"
	"errors"
	"fmt"
)

type CircuitRepository struct {
    queries *queries.Queries
}

func NewCircuitRepository(db *sql.DB) *CircuitRepository {
    return &CircuitRepository{
        queries: queries.New(db),
    }
}

func (r *CircuitRepository) CreateCircuit(req *models.CreateCircuitRequest) (*models.Circuit, error) {
	circuit, err := r.queries.CreateCircuit(context.Background(), queries.CreateCircuitParams{
        Name:        req.Name,
        CircuitName: req.CircuitName,
        StartDate:   req.StartDate,
        EndDate:     req.EndDate,
        Status:      req.Status,
    })
	
    if err != nil {
        return nil, err
    }
    
    return &models.Circuit{
        ID:          circuit.ID,
        Name:        circuit.Name,
        CircuitName: circuit.CircuitName,
        StartDate:   circuit.StartDate,
        EndDate:     circuit.EndDate,
        Status:      circuit.Status,
        CreatedAt:   time.Now(),
        UpdatedAt:   time.Now(),
    }, nil
}

func (r *CircuitRepository) UpdateCircuit(id uuid.UUID,req *models.UpdateCircuitRequest) (*models.Circuit, error) {
	_,err := r.GetCircuitByID(id)
	if err != nil {
		return nil, err
	}
	circuit, err := r.queries.UpdateCircuit(context.Background(),queries.UpdateCircuitParams{
		ID:          id,
		Name:        req.Name,
		CircuitName: req.CircuitName,
		StartDate:   req.StartDate,
		EndDate:     req.EndDate,
		Status:     req.Status,
	})

	if err != nil {
		return nil, err
	}
	
	return &models.Circuit{
		ID:          circuit.ID,
		Name:        circuit.Name,
		CircuitName: circuit.CircuitName,
		StartDate:   circuit.StartDate,
		EndDate:     circuit.EndDate,
		Status:      circuit.Status,
		CreatedAt:   circuit.CreatedAt,
		UpdatedAt:   circuit.UpdatedAt,
	}, nil
}

func (r *CircuitRepository) GetCircuitByID(id uuid.UUID) (*models.Circuit, error) {
	circuit, err := r.queries.GetCircuitByID(context.Background(), id)
	if err != nil {
		return nil, err
	}
	return &models.Circuit{
		ID:          circuit.ID,
		Name:        circuit.Name,
		CircuitName: circuit.CircuitName,
		StartDate:   circuit.StartDate,
		EndDate:     circuit.EndDate,
		Status:      circuit.Status,
		CreatedAt:   circuit.CreatedAt,
		UpdatedAt:   circuit.UpdatedAt,
	}, nil
}

func (r *CircuitRepository) ListCircuits(ctx context.Context, limit, offset int) ([]*models.Circuit, error) {
	circuits, err := r.queries.ListCircuits(context.Background(), queries.ListCircuitsParams{
		Limit:  int32(limit),
		Offset: int32(offset),

	})
	if err != nil {
		return nil, err
	}
	
	var circuitsList []*models.Circuit
	for _, circuit := range circuits {
		circuitsList = append(circuitsList, &models.Circuit{
			ID:          circuit.ID,
			Name:        circuit.Name,
			CircuitName: circuit.CircuitName,
			StartDate:   circuit.StartDate,
			EndDate:     circuit.EndDate,
			Status:      circuit.Status,
			CreatedAt:   circuit.CreatedAt,
			UpdatedAt:   circuit.UpdatedAt,
		})
	}
	
	return circuitsList, nil
}


func (r *CircuitRepository) Delete(ctx context.Context, id uuid.UUID) error {
    // Step 1: Check if the circuit exists. This is still a crucial step.
    _, err := r.GetCircuitByID(id)
    if err != nil {
        // If the error is sql.ErrNoRows, it means we can't find it.
        if errors.Is(err, sql.ErrNoRows) {
            return fmt.Errorf("cannot delete: circuit with id %s not found", id)
        }
        // For any other error (DB down, etc.), return it directly.
        return err
    }

    // Step 2: Call the sqlc-generated function, which now performs the soft delete (UPDATE).
    err = r.queries.DeleteCircuit(ctx, id)
    if err != nil {
        // Handle any errors from the UPDATE command itself.
        return err
    }

    return nil
}