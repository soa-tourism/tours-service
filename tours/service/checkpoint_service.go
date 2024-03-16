package service

import (
	"fmt"
	"tours/model"
	"tours/repo"
)

type CheckpointService struct {
	CheckpointRepo *repo.CheckpointRepository
}

func (service *CheckpointService) FindAllCheckpoints() ([]model.Checkpoint, error) {
	checkpoints, err := service.CheckpointRepo.FindAll()
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve checkpoints: %v", err)
	}
	return checkpoints, nil
}

func (service *CheckpointService) FindCheckpoint(id int64) (*model.Checkpoint, error) {
	checkpoint, err := service.CheckpointRepo.FindById(id)
	if err != nil {
		return nil, fmt.Errorf(" checkpoint with ID %d not found", id)
	}
	return &checkpoint, nil
}

func (service *CheckpointService) CreateCheckpoint(checkpoint *model.Checkpoint) (*model.Checkpoint, error) {
	createdCheckpoint, err := service.CheckpointRepo.Create(checkpoint)
	if err != nil {
		return nil, fmt.Errorf("failed to create checkpoint: %v", err)
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

func (service *CheckpointService) DeleteCheckpoint(id int64) error {
	err := service.CheckpointRepo.Delete(id)
	if err != nil {
		return fmt.Errorf("failed to delete checkpoint with ID %d: %v", id, err)
	}
	return nil
}

func (service *CheckpointService) FindCheckpointsByTour(id int64) ([]model.Checkpoint, error) {
	checkpoints, err := service.CheckpointRepo.FindByTour(id)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve checkpoints by tour: %v", err)
	}
	return checkpoints, nil
}
