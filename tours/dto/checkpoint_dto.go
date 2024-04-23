package dto

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"tours/model"
)

type CheckpointDto struct {
	Id                    primitive.ObjectID     `json:"Id"`
	TourId                primitive.ObjectID     `json:"TourId"`
	AuthorId              int64                  `json:"AuthorId"`
	Longitude             float64                `json:"Longitude"`
	Latitude              float64                `json:"Latitude"`
	Name                  string                 `json:"Name"`
	Description           string                 `json:"Description"`
	Pictures              []string               `json:"Pictures"`
	RequiredTimeInSeconds float64                `json:"RequiredTimeInSeconds"`
	IsSecretPrerequisite  bool                   `json:"IsSecretPrerequisite"`
	EncounterId           int64                  `json:"EncounterId"`
	CheckpointSecret      model.CheckpointSecret `json:"CheckpointSecret"`
}

func (dto *CheckpointDto) MapToModel() *model.Checkpoint {
	return &model.Checkpoint{
		ID:                    dto.Id,
		TourID:                dto.TourId,
		AuthorID:              dto.AuthorId,
		Longitude:             dto.Longitude,
		Latitude:              dto.Latitude,
		Name:                  dto.Name,
		Description:           dto.Description,
		Pictures:              dto.Pictures,
		RequiredTimeInSeconds: dto.RequiredTimeInSeconds,
		IsSecretPrerequisite:  dto.IsSecretPrerequisite,
		EncounterID:           dto.EncounterId,
		CheckpointSecret:      dto.CheckpointSecret,
	}
}

func CheckpointDtoFromModel(ch model.Checkpoint) CheckpointDto {
	return CheckpointDto{
		Id:                    ch.ID,
		TourId:                ch.TourID,
		AuthorId:              ch.AuthorID,
		Longitude:             ch.Longitude,
		Latitude:              ch.Latitude,
		Name:                  ch.Name,
		Description:           ch.Description,
		Pictures:              ch.Pictures,
		RequiredTimeInSeconds: ch.RequiredTimeInSeconds,
		IsSecretPrerequisite:  ch.IsSecretPrerequisite,
		EncounterId:           ch.EncounterID,
		CheckpointSecret:      ch.CheckpointSecret,
	}
}
