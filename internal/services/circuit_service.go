package services

import (
    "fast-line/internal/models"
    "fast-line/internal/repository"
	// "fast-line/internal/utils"
	"github.com/google/uuid"
	"errors"
	"context"
)

type CircuitService struct {
    circuitRepo *repository.CircuitRepository
}

func NewCircuitService(circuitRepo *repository.CircuitRepository) *CircuitService {
    return &CircuitService{
        circuitRepo: circuitRepo,
    }
}

func (s *CircuitService) Create(req *models.CreateCircuitRequest) (*models.Circuit, error) {

	circuit, err := s.circuitRepo.CreateCircuit(req)
    if err != nil {
        return nil, err
    }
    return circuit, nil
}

func (s *CircuitService) Update(id uuid.UUID,req *models.UpdateCircuitRequest) (*models.Circuit, error) {
	existingCircuit, _ := s.circuitRepo.GetCircuitByID(id)
	if existingCircuit == nil {
		return nil, errors.New("circuit with this id does not exist")
	}

	circuit, err := s.circuitRepo.UpdateCircuit(id,req)
    if err != nil {
        return nil, err
    }
    
    return circuit, nil

}

func (s *CircuitService) Get(id uuid.UUID) (*models.Circuit, error) {
	existingCircuit, _ := s.circuitRepo.GetCircuitByID(id)
	if existingCircuit == nil {
		return nil, errors.New("circuit with this id does not exist")
	}
	return existingCircuit, nil
}

func (s *CircuitService) List(limit, offset int) ([]*models.Circuit, error) {
	circuits, err := s.circuitRepo.ListCircuits(context.Background(), limit, offset)
	if err != nil {
		return nil, err
	}
	return circuits, nil
}

func (s *CircuitService) Delete(id uuid.UUID) error {
    // The service delegates the deletion task to the repository.
    // We pass a background context for the operation.
    return s.circuitRepo.Delete(context.Background(), id)
}