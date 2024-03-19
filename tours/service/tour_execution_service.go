package service

import (
	"fmt"
	"tours/model"
	"tours/repo"
)

type TourExecutionService struct {
	TourExecutionRepo *repo.TourExecutionRepository
}

func (service *TourExecutionService) FindAllTourExecutions() ([]model.TourExecution, error) {
	tourExecutions, err := service.TourExecutionRepo.FindAll()
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve tour executions: %v", err)
	}
	return tourExecutions, nil
}

func (service *TourExecutionService) FindTourExecution(id int64) (*model.TourExecution, error) {
	tourExecution, err := service.TourExecutionRepo.FindById(id)
	if err != nil {
		return nil, fmt.Errorf("tour execution with ID %d not found", id)
	}
	return &tourExecution, nil
}

func (service *TourExecutionService) CreateTourExecution(tourExecution *model.TourExecution) (*model.TourExecution, error) {
	createdTourExecution, err := service.TourExecutionRepo.Create(tourExecution)
	if err != nil {
		return nil, fmt.Errorf("failed to create tour execution: %v", err)
	}
	return &createdTourExecution, nil
}

func (service *TourExecutionService) UpdateTourExecution(tourExecution *model.TourExecution) (*model.TourExecution, error) {
	updatedTourExecution, err := service.TourExecutionRepo.Update(tourExecution)
	if err != nil {
		return nil, fmt.Errorf("failed to update tour execution: %v", err)
	}
	return &updatedTourExecution, nil
}

func (service *TourExecutionService) DeleteTourExecution(id int64) error {
	err := service.TourExecutionRepo.Delete(id)
	if err != nil {
		return fmt.Errorf("failed to delete tour execution with ID %d: %v", id, err)
	}
	return nil
}
