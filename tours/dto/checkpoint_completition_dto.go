package dto

import (
	"time"
	"tours/model"
)

type CheckpointCompletitionDto struct {
	Id               int64     `gorm:"primaryKey"`
	TourExecutionId  int64     `json:"TourExecutionId"`
	CheckpointId     int64     `json:"CheckpointId"`
	CompletitionTime time.Time `json:"CompletitionTime"`
}

func (dto *CheckpointCompletitionDto) MapToModel() *model.CheckpointCompletition {
	return &model.CheckpointCompletition{
		Id:               dto.Id,
		TourExecutionId:  dto.TourExecutionId,
		CheckpointId:     dto.CheckpointId,
		CompletitionTime: dto.CompletitionTime,
	}
}

func CheckpointCompletitionDtoFromModel(chCompletition model.CheckpointCompletition) CheckpointCompletitionDto {
	return CheckpointCompletitionDto{
		Id:               chCompletition.Id,
		TourExecutionId:  chCompletition.TourExecutionId,
		CheckpointId:     chCompletition.CheckpointId,
		CompletitionTime: chCompletition.CompletitionTime,
	}
}
