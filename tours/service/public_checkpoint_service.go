package service

import (
	"fmt"
	"tours/model"
	"tours/repo"
)

type PublicCheckpointService struct {
	CheckpointRepo *repo.PublicCheckpointRepository
}

func (service *PublicCheckpointService) FindAllPublicCheckpoints() ([]model.PublicCheckpoint, error) {
	checkpoints, err := service.CheckpointRepo.FindAll()
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve public checkpoints: %v", err)
	}
	return checkpoints, nil
}

func (service *PublicCheckpointService) FindPublicCheckpoint(id int64) (*model.PublicCheckpoint, error) {
	checkpoint, err := service.CheckpointRepo.FindById(id)
	if err != nil {
		return nil, fmt.Errorf("public checkpoint with ID %d not found", id)
	}
	return &checkpoint, nil
}

func (service *PublicCheckpointService) CreatePublicCheckpoint(checkpoint *model.PublicCheckpoint) (*model.PublicCheckpoint, error) {
	createdCheckpoint, err := service.CheckpointRepo.Create(checkpoint)
	if err != nil {
		return nil, fmt.Errorf("failed to create public checkpoint: %v", err)
	}
	return &createdCheckpoint, nil
}

func (service *PublicCheckpointService) UpdatePublicCheckpoint(checkpoint *model.PublicCheckpoint) (*model.PublicCheckpoint, error) {
	updatedCheckpoint, err := service.CheckpointRepo.Update(checkpoint)
	if err != nil {
		return nil, fmt.Errorf("failed to update public checkpoint: %v", err)
	}
	return &updatedCheckpoint, nil
}

func (service *PublicCheckpointService) DeletePublicCheckpoint(id int64) error {
	err := service.CheckpointRepo.Delete(id)
	if err != nil {
		return fmt.Errorf("failed to delete public checkpoint with ID %d: %v", id, err)
	}
	return nil
}
