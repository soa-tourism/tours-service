package dto

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"tours/model"
)

type TouristPositionDto struct {
	Id        primitive.ObjectID `json:"Id"`
	CreatorId int64              `json:"CreatorId"`
	Longitude float64            `json:"Longitude"`
	Latitude  float64            `json:"Latitude"`
}

func (dto *TouristPositionDto) MapToModel() *model.TouristPosition {
	return &model.TouristPosition{
		ID:        dto.Id,
		CreatorID: dto.CreatorId,
		Longitude: dto.Longitude,
		Latitude:  dto.Latitude,
	}
}

func TouristPositionDtoFromModel(pos model.TouristPosition) TouristPositionDto {
	return TouristPositionDto{
		Id:        pos.ID,
		CreatorId: pos.CreatorID,
		Longitude: pos.Longitude,
		Latitude:  pos.Latitude,
	}
}
