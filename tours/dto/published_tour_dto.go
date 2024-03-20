package dto

import "tours/model"

type PublishedTourDto struct {
	Id          int64             `json:"Id"`
	AuthorId    int64             `json:"AuthorId"`
	Name        string            `json:"Name"`
	Description string            `json:"Description"`
	Difficulty  string            `json:"Difficulty"`
	Price       float64           `json:"Price"`
	Tags        []string          `json:"Tags"`
	Equipment   []model.Equipment `json:"Equipment"`
}

func MapToPublishedTour(tour model.Tour) PublishedTourDto {
	publishedTourDto := PublishedTourDto{
		Id:          tour.Id,
		AuthorId:    tour.AuthorId,
		Name:        tour.Name,
		Description: tour.Description,
		Difficulty:  tour.Difficulty.String(),
		Price:       tour.Price,
		Tags:        tour.Tags,
		Equipment:   tour.Equipment,
	}
	return publishedTourDto
}
