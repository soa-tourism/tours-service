package dto

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"tours/model"
)

type TourDto struct {
	Id          primitive.ObjectID `json:"Id"`
	AuthorId    int64              `json:"AuthorId"`
	Name        string             `json:"Name"`
	Description string             `json:"Description"`
	Difficulty  string             `json:"Difficulty"`
	Status      string             `json:"Status"`
	Price       float64            `json:"Price"`
	Tags        []string           `json:"Tags"`
	Equipment   []model.Equipment  `json:"Equipment"`
	Checkpoints []model.Checkpoint `json:"Checkpoints"`
}

func (dto *TourDto) MapToModel() *model.Tour {
	equipment := make([]model.Equipment, len(dto.Equipment))
	for i, equip := range dto.Equipment {
		equipment[i] = model.Equipment{
			ID:          equip.ID,
			Name:        equip.Name,
			Description: equip.Description,
		}
	}

	checkpoints := make([]model.Checkpoint, len(dto.Checkpoints))
	for i, ch := range dto.Checkpoints {
		checkpoints[i] = model.Checkpoint{
			ID:                    ch.ID,
			TourID:                ch.TourID,
			AuthorID:              ch.AuthorID,
			Longitude:             ch.Longitude,
			Latitude:              ch.Latitude,
			Name:                  ch.Name,
			Description:           ch.Description,
			Pictures:              ch.Pictures,
			RequiredTimeInSeconds: ch.RequiredTimeInSeconds,
			IsSecretPrerequisite:  ch.IsSecretPrerequisite,
			EncounterID:           ch.EncounterID,
			CheckpointSecret:      ch.CheckpointSecret,
		}
	}

	return &model.Tour{
		ID:          dto.Id,
		AuthorID:    dto.AuthorId,
		Name:        dto.Name,
		Description: dto.Description,
		Difficulty:  model.ParseDifficulty(dto.Difficulty),
		Status:      model.ParseStatus(dto.Status),
		Price:       dto.Price,
		Tags:        dto.Tags,
		Equipment:   equipment,
		Checkpoints: checkpoints,
	}
}

func MapFromTour(tour model.Tour) TourDto {
	return TourDto{
		Id:          tour.ID,
		AuthorId:    tour.AuthorID,
		Name:        tour.Name,
		Description: tour.Description,
		Difficulty:  tour.Difficulty.String(),
		Status:      tour.Status.String(),
		Price:       tour.Price,
		Tags:        tour.Tags,
		Equipment:   tour.Equipment,
		Checkpoints: tour.Checkpoints,
	}
}
