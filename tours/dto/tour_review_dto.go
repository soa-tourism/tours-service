package dto

import (
	"time"
	"tours/model"
)

type TourReviewDto struct {
	Id         int64     `gorm:"primaryKey"`
	Rating     int       `json:"Rating"`
	Comment    string    `json:"Comment"`
	TouristId  int64     `json:"TouristId"`
	TourId     int64     `json:"TourId"`
	TourDate   time.Time `json:"TourDate"`
	ReviewDate time.Time `json:"ReviewDate"`
	Images     [][]byte  `json:"Images"`
}

func MapToModel(dto *TourReviewDto) *model.TourReview {
	tourReview := &model.TourReview{
		Id:         dto.Id,
		Rating:     dto.Rating,
		Comment:    dto.Comment,
		TouristId:  dto.TouristId,
		TourId:     dto.TourId,
		TourDate:   dto.TourDate,
		ReviewDate: dto.ReviewDate,
	}

	return tourReview
}
