package dto

import "tours/model"

type CheckpointDto struct {
	Id                    int64                  `json:"Id"`
	TourId                int64                  `json:"TourId"`
	AuthorId              int64                  `json:"AuthorId"`
	Longitude             float64                `json:"Longitude"`
	Latitude              float64                `json:"Latitude"`
	Name                  string                 `json:"Name"`
	Description           string                 `json:"Description"`
	Pictures              []string               `json:"Pictures"`
	RequiredTimeInSeconds float64                `json:"RequiredTimeInSeconds"`
	CheckpointSecret      model.CheckpointSecret `json:"CheckpointSecret"`
	IsSecretPrerequisite  bool                   `json:"IsSecretPrerequisite"`
	EncounterId           int64                  `json:"EncounterId"`
}

func (dto *CheckpointDto) MapToModel() *model.Checkpoint {
	return &model.Checkpoint{
		Id:                    dto.Id,
		TourId:                dto.TourId,
		AuthorId:              dto.AuthorId,
		Longitude:             dto.Longitude,
		Latitude:              dto.Latitude,
		Name:                  dto.Name,
		Description:           dto.Description,
		Pictures:              dto.Pictures,
		RequiredTimeInSeconds: dto.RequiredTimeInSeconds,
		CheckpointSecret:      dto.CheckpointSecret,
		IsSecretPrerequisite:  dto.IsSecretPrerequisite,
		EncounterId:           dto.EncounterId,
	}
}

func CheckpointDtoFromModel(ch model.Checkpoint) CheckpointDto {
	return CheckpointDto{
		Id:                    ch.Id,
		TourId:                ch.TourId,
		AuthorId:              ch.AuthorId,
		Longitude:             ch.Longitude,
		Latitude:              ch.Latitude,
		Name:                  ch.Name,
		Description:           ch.Description,
		Pictures:              ch.Pictures,
		RequiredTimeInSeconds: ch.RequiredTimeInSeconds,
		CheckpointSecret:      ch.CheckpointSecret,
		IsSecretPrerequisite:  ch.IsSecretPrerequisite,
		EncounterId:           ch.EncounterId,
	}
}
