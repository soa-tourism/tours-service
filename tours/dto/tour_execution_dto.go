package dto

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"

	"tours/model"
)

type TourExecutionDto struct {
	Id                   primitive.ObjectID             `json:"Id"`
	TouristId            int64                          `json:"TouristId"`
	TourId               primitive.ObjectID             `json:"TourId"`
	Start                time.Time                      `json:"Start"`
	LastActivity         time.Time                      `json:"LastActivity"`
	ExecutionStatus      string                         `json:"ExecutionStatus"`
	CompletedCheckpoints []model.CheckpointCompletition `json:"CompletedCheckpoints"`
}

func (dto *TourExecutionDto) MapToModel() *model.TourExecution {
	completedCheckpoints := make([]model.CheckpointCompletition, len(dto.CompletedCheckpoints))
	for i, checkpoint := range dto.CompletedCheckpoints {
		completedCheckpoints[i] = model.CheckpointCompletition{
			ID:               checkpoint.ID,
			TourExecutionID:  checkpoint.TourExecutionID,
			CheckpointID:     checkpoint.CheckpointID,
			CompletitionTime: checkpoint.CompletitionTime,
		}
	}

	return &model.TourExecution{
		ID:                   dto.Id,
		TouristID:            dto.TouristId,
		TourID:               dto.TourId,
		Start:                dto.Start,
		LastActivity:         dto.LastActivity,
		ExecutionStatus:      model.ParseExecutionStatus(dto.ExecutionStatus),
		CompletedCheckpoints: completedCheckpoints,
	}
}

func TourExecutionDtoFromModel(execution model.TourExecution) TourExecutionDto {
	completedCheckpoints := make([]model.CheckpointCompletition, len(execution.CompletedCheckpoints))
	for i, checkpoint := range execution.CompletedCheckpoints {
		completedCheckpoints[i] = model.CheckpointCompletition{
			ID:               checkpoint.ID,
			TourExecutionID:  checkpoint.TourExecutionID,
			CheckpointID:     checkpoint.CheckpointID,
			CompletitionTime: checkpoint.CompletitionTime,
		}
	}

	return TourExecutionDto{
		Id:                   execution.ID,
		TouristId:            execution.TouristID,
		TourId:               execution.TourID,
		Start:                execution.Start,
		LastActivity:         execution.LastActivity,
		ExecutionStatus:      execution.ExecutionStatus.String(),
		CompletedCheckpoints: completedCheckpoints,
	}
}
