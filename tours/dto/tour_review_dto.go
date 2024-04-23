package dto

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
	"tours/model"
)

type TourReviewDto struct {
	Id         primitive.ObjectID `json:"Id"`
	Rating     int                `json:"Rating"`
	Comment    string             `json:"Comment"`
	TouristId  int64              `json:"TouristId"`
	TourId     primitive.ObjectID `json:"TourId"`
	TourDate   time.Time          `json:"TourDate"`
	ReviewDate time.Time          `json:"ReviewDate"`
	Images     [][]byte           `json:"Images"`
}

func MapToModel(dto *TourReviewDto) *model.TourReview {
	tourReview := &model.TourReview{
		ID:         dto.Id,
		Rating:     dto.Rating,
		Comment:    dto.Comment,
		TouristID:  dto.TouristId,
		TourID:     dto.TourId,
		TourDate:   dto.TourDate,
		ReviewDate: dto.ReviewDate,
	}

	return tourReview
}
