package dto

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"tours/model"
)

type PublicCheckpointDto struct {
	Id          primitive.ObjectID `json:"Id"`
	Longitude   float64            `json:"Longitude"`
	Latitude    float64            `json:"Latitude"`
	Name        string             `json:"Name"`
	Description string             `json:"Description"`
	Pictures    []string           `json:"Pictures"`
}

func (dto *PublicCheckpointDto) MapToModel() *model.PublicCheckpoint {
	return &model.PublicCheckpoint{
		ID:          dto.Id,
		Longitude:   dto.Longitude,
		Latitude:    dto.Latitude,
		Name:        dto.Name,
		Description: dto.Description,
		Pictures:    dto.Pictures,
	}
}

func PublicCheckpointDtoFromModel(pch model.PublicCheckpoint) PublicCheckpointDto {
	return PublicCheckpointDto{
		Id:          pch.ID,
		Longitude:   pch.Longitude,
		Latitude:    pch.Latitude,
		Name:        pch.Name,
		Description: pch.Description,
		Pictures:    pch.Pictures,
	}
}
