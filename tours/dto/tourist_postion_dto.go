package dto

import "tours/model"

type TouristPositionDto struct {
	Id        int64   `json:"Id"`
	CreatorId int64   `json:"CreatorId"`
	Longitude float64 `json:"Longitude"`
	Latitude  float64 `json:"Latitude"`
}

func (dto *TouristPositionDto) MapToModel() *model.TouristPosition {
	return &model.TouristPosition{
		Id:        dto.Id,
		CreatorId: dto.CreatorId,
		Longitude: dto.Longitude,
		Latitude:  dto.Latitude,
	}
}

func TouristPositionDtoFromModel(pos model.TouristPosition) TouristPositionDto {
	return TouristPositionDto{
		Id:        pos.Id,
		CreatorId: pos.CreatorId,
		Longitude: pos.Longitude,
		Latitude:  pos.Latitude,
	}
}
