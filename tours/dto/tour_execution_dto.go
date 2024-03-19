package dto

import (
	"time"

	"tours/model"
)

type TourExecutionDto struct {
	Id                   int64 `gorm:"primaryKey"`
	TouristId            int64 `json:"TouristId"`
	TourId               int64 `json:"TourId"`
	Tour                 model.Tour
	Start                time.Time                      `json:"Start"`
	LastActivity         time.Time                      `json:"LastActivity"`
	ExecutionStatus      model.ExecutionStatus          `json:"ExecutionStatus"`
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
		ExecutionStatus:      model.ParseExecutionStatus(dto.ExecutionStatus.String()),
		CompletedCheckpoints: completedCheckpoints,
	}
}

func CheckpointCompletitionFromModel(dto model.TourExecution) TourExecutionDto {
	completedCheckpoints := make([]model.CheckpointCompletition, len(dto.CompletedCheckpoints))
	for i, checkpoint := range dto.CompletedCheckpoints {
		completedCheckpoints[i] = model.CheckpointCompletition{
			Id:               checkpoint.Id,
			TourExecutionId:  checkpoint.TourExecutionId,
			CheckpointId:     checkpoint.CheckpointId,
			CompletitionTime: checkpoint.CompletitionTime,
		}
	}

	return TourExecutionDto{
		Id:                   dto.Id,
		TouristId:            dto.TouristId,
		TourId:               dto.TourId,
		Start:                dto.Start,
		LastActivity:         dto.LastActivity,
		ExecutionStatus:      dto.ExecutionStatus,
		CompletedCheckpoints: completedCheckpoints,
	}
}
