package service

import (
	"errors"
	"fmt"
	"tours/model"
	"tours/repo"

	"gorm.io/gorm"
)

type TouristPositionService struct {
	TouristPositionRepo *repo.TouristPositionRepository
}

func (service *TouristPositionService) FindAllTouristPositions() ([]model.TouristPosition, error) {
	TouristPosition, err := service.TouristPositionRepo.FindAll()
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve tourist position: %v", err)
	}
	return TouristPosition, nil
}

func (service *TouristPositionService) FindTouristPosition(id int64) (*model.TouristPosition, error) {
	TouristPosition, err := service.TouristPositionRepo.FindById(id)
	if err != nil {
		return nil, fmt.Errorf("tourist position with ID %d not found", id)
	}
	return &TouristPosition, nil
}

func (service *TouristPositionService) CreateTouristPosition(touristPosition *model.TouristPosition) (*model.TouristPosition, error) {
	// Check if tourist position for that creatorId already exists
	existingPosition, err := service.TouristPositionRepo.FindByCreator(touristPosition.CreatorId)

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, fmt.Errorf("failed to find tourist position: %v", err)
	}

	if existingPosition.Id == 0 {
		// Tourist position for that creator id doesn't exist -> Create
		createdTouristPosition, err := service.TouristPositionRepo.Create(touristPosition)
		if err != nil {
			return nil, fmt.Errorf("failed to create tourist position: %v", err)
		}
		return &createdTouristPosition, nil
	} else {
		// Tourist position for that creator id exists -> Update
		existingPosition.Longitude = touristPosition.Longitude
		existingPosition.Latitude = touristPosition.Latitude

		updatedTouristPosition, err := service.TouristPositionRepo.Update(&existingPosition)
		if err != nil {
			return nil, fmt.Errorf("failed to update tourist position: %v", err)
		}
		return &updatedTouristPosition, nil
	}
}

func (service *TouristPositionService) UpdateTouristPosition(TouristPosition *model.TouristPosition) (*model.TouristPosition, error) {
	updatedTouristPosition, err := service.TouristPositionRepo.Update(TouristPosition)
	if err != nil {
		return nil, fmt.Errorf("failed to update TouristPosition: %v", err)
	}
	return &updatedTouristPosition, nil
}

func (service *TouristPositionService) DeleteTouristPosition(id int64) error {
	err := service.TouristPositionRepo.Delete(id)
	if err != nil {
		return fmt.Errorf("failed to delete tourist position with ID %d: %v", id, err)
	}
	return nil
}

func (service *TouristPositionService) FindPositionByCreator(id int64) (model.TouristPosition, error) {
	position, err := service.TouristPositionRepo.FindByCreator(id)
	if err != nil {
		return position, fmt.Errorf("failed to retrieve tourist position: %v", err)
	}
	return position, nil
}
