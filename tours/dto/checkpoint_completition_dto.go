package dto

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
	"tours/model"
)

type CheckpointCompletitionDto struct {
	Id               primitive.ObjectID `json:"Id"`
	TourExecutionId  primitive.ObjectID `json:"TourExecutionId"`
	CheckpointId     primitive.ObjectID `json:"CheckpointId"`
	CompletitionTime time.Time          `json:"CompletitionTime"`
}

func (dto *CheckpointCompletitionDto) MapToModel() *model.CheckpointCompletition {
	return &model.CheckpointCompletition{
		ID:               dto.Id,
		TourExecutionID:  dto.TourExecutionId,
		CheckpointID:     dto.CheckpointId,
		CompletitionTime: dto.CompletitionTime,
	}
}

func CheckpointCompletitionDtoFromModel(chCompletition model.CheckpointCompletition) CheckpointCompletitionDto {
	return CheckpointCompletitionDto{
		Id:               chCompletition.ID,
		TourExecutionId:  chCompletition.TourExecutionID,
		CheckpointId:     chCompletition.CheckpointID,
		CompletitionTime: chCompletition.CompletitionTime,
	}
}
