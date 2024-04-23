package dto

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"tours/model"
)

type PublishedTourDto struct {
	Id          primitive.ObjectID `json:"Id"`
	AuthorId    int64              `json:"AuthorId"`
	Name        string             `json:"Name"`
	Description string             `json:"Description"`
	Difficulty  string             `json:"Difficulty"`
	Price       float64            `json:"Price"`
	Tags        []string           `json:"Tags"`
	Equipment   []model.Equipment  `json:"Equipment"`
}

func MapToPublishedTour(tour model.Tour) PublishedTourDto {
	publishedTourDto := PublishedTourDto{
		Id:          tour.ID,
		AuthorId:    tour.AuthorID,
		Name:        tour.Name,
		Description: tour.Description,
		Difficulty:  tour.Difficulty.String(),
		Price:       tour.Price,
		Tags:        tour.Tags,
		Equipment:   tour.Equipment,
	}
	return publishedTourDto
}
