package dto

import "tours/model"

type TourDto struct {
	Id          int64             `json:"Id"`
	AuthorId    int64             `json:"AuthorId"`
	Name        string            `json:"Name"`
	Description string            `json:"Description"`
	Difficulty  string            `json:"Difficulty"`
	Status      string            `json:"Status"`
	Price       float64           `json:"Price"`
	Tags        []string          `json:"Tags"`
	Equipment   []model.Equipment `json:"Equipment"`
}

func (dto *TourDto) MapToModel() *model.Tour {
	equipment := make([]model.Equipment, len(dto.Equipment))
	for i, equip := range dto.Equipment {
		equipment[i] = model.Equipment{
			Id:          equip.Id,
			Name:        equip.Name,
			Description: equip.Description,
		}
	}

	return &model.Tour{
		Id:          dto.Id,
		AuthorId:    dto.AuthorId,
		Name:        dto.Name,
		Description: dto.Description,
		Difficulty:  model.ParseDifficulty(dto.Difficulty),
		Status:      model.ParseStatus(dto.Status),
		Price:       dto.Price,
		Tags:        dto.Tags,
		Equipment:   equipment,
	}
}

func MapFromModel(tour model.Tour) TourDto {
	return TourDto{
		Id:          tour.Id,
		AuthorId:    tour.AuthorId,
		Name:        tour.Name,
		Description: tour.Description,
		Difficulty:  tour.Difficulty.String(),
		Status:      tour.Status.String(),
		Price:       tour.Price,
		Tags:        tour.Tags,
		Equipment:   tour.Equipment,
	}
}
