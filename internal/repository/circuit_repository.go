package repository

import (
	"time"
	"context"
	"database/sql"
	"fast-line/internal/models"
	"fast-line/internal/repository/queries"
	"github.com/google/uuid"
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

func (r *CircuitRepository) UpdateCircuit(req *models.UpdateCircuitRequest) (*models.Circuit, error) {
	

	circuit, err := r.queries.UpdateCircuit(context.Background(),queries.UpdateCircuitParams{
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