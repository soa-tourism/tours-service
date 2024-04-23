package service

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"tours/model"
	"tours/repo"
)

type CheckpointService struct {
	CheckpointRepo *repo.CheckpointRepository
	TourRepo       *repo.TourRepository
}

func (service *CheckpointService) FindAllCheckpoints() ([]model.Checkpoint, error) {
	checkpoints, err := service.CheckpointRepo.FindAll()
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve checkpoints: %v", err)
	}
	return checkpoints, nil
}

func (service *CheckpointService) FindCheckpoint(id primitive.ObjectID) (*model.Checkpoint, error) {
	checkpoint, err := service.CheckpointRepo.FindById(id)
	if err != nil {
		return nil, fmt.Errorf("checkpoint with ID %s not found", id.Hex())
	}
	return &checkpoint, nil
}

func (service *CheckpointService) CreateCheckpoint(checkpoint *model.Checkpoint) (*model.Checkpoint, error) {
	checkpoint.ID = primitive.NewObjectID()
	createdCheckpoint, err := service.CheckpointRepo.Create(checkpoint)
	if err != nil {
		return nil, fmt.Errorf("failed to create checkpoint: %v", err)
	}

	tourID := checkpoint.TourID
	if tourID.IsZero() {
		return nil, fmt.Errorf("checkpoint must belong to a tour")
	}
	tour, err := service.TourRepo.FindById(tourID)
	if err != nil {
		return nil, fmt.Errorf("failed to find tour: %v", err)
	}
	tour.Checkpoints = append(tour.Checkpoints, createdCheckpoint)
	_, err = service.TourRepo.Update(tour)
	if err != nil {
		return nil, fmt.Errorf("failed to update tour with new checkpoint: %v", err)
	}

	return &createdCheckpoint, nil
}

func (service *CheckpointService) UpdateCheckpoint(checkpoint *model.Checkpoint) (*model.Checkpoint, error) {
	updatedCheckpoint, err := service.CheckpointRepo.Update(checkpoint)
	if err != nil {
		return nil, fmt.Errorf("failed to update checkpoint: %v", err)
	}
	return &updatedCheckpoint, nil
}

func (service *CheckpointService) DeleteCheckpoint(id primitive.ObjectID) error {
	err := service.CheckpointRepo.Delete(id)
	if err != nil {
		return fmt.Errorf("failed to delete checkpoint with ID %s: %v", id.Hex(), err)
	}
	return nil
}

func (service *CheckpointService) FindCheckpointsByTour(id primitive.ObjectID) ([]model.Checkpoint, error) {
	checkpoints, err := service.CheckpointRepo.FindByTour(id)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve checkpoints by tour: %v", err)
	}
	return checkpoints, nil
}

func (service *CheckpointService) UpdateCheckpointSecret(id primitive.ObjectID, secret *model.CheckpointSecret) (*model.Checkpoint, error) {
	checkpoint1, err := service.CheckpointRepo.FindById(id)
	if err != nil {
		return nil, fmt.Errorf(" checkpoint with ID %d not found", id)
	}

	checkpoint1.UpdateCheckpointSecret(secret.Description, secret.Pictures)

	updatedCheckpoint, err := service.CheckpointRepo.Update(&checkpoint1)
	if err != nil {
		return nil, fmt.Errorf("failed to update checkpoint: %v", err)
	}
	return &updatedCheckpoint, nil
}
