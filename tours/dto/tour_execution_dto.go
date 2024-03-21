package dto

import (
	"time"

	"tours/model"
)

type TourExecutionDto struct {
	Id                   int64                          `gorm:"primaryKey"`
	TouristId            int64                          `json:"TouristId"`
	TourId               int64                          `json:"TourId"`
	Start                time.Time                      `json:"Start"`
	LastActivity         time.Time                      `json:"LastActivity"`
	ExecutionStatus      string                         `json:"ExecutionStatus"`
	CompletedCheckpoints []model.CheckpointCompletition `json:"CompletedCheckpoints"`
}

func (dto *TourExecutionDto) MapToModel() *model.TourExecution {
	completedCheckpoints := make([]model.CheckpointCompletition, len(dto.CompletedCheckpoints))
	for i, checkpoint := range dto.CompletedCheckpoints {
		completedCheckpoints[i] = model.CheckpointCompletition{
			Id:               checkpoint.Id,
			TourExecutionId:  checkpoint.TourExecutionId,
			CheckpointId:     checkpoint.CheckpointId,
			CompletitionTime: checkpoint.CompletitionTime,
		}
	}

	return &model.TourExecution{
		Id:                   dto.Id,
		TouristId:            dto.TouristId,
		TourId:               dto.TourId,
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
			Id:               checkpoint.Id,
			TourExecutionId:  checkpoint.TourExecutionId,
			CheckpointId:     checkpoint.CheckpointId,
			CompletitionTime: checkpoint.CompletitionTime,
		}
	}

	return TourExecutionDto{
		Id:                   execution.Id,
		TouristId:            execution.TouristId,
		TourId:               execution.TourId,
		Start:                execution.Start,
		LastActivity:         execution.LastActivity,
		ExecutionStatus:      execution.ExecutionStatus.String(),
		CompletedCheckpoints: completedCheckpoints,
	}
}
