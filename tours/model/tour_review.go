package model

import (
	"database/sql/driver"
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gorm.io/gorm"
	"strings"
	"time"
)

type TourReview struct {
	ID         primitive.ObjectID `json:"Id" bson:"_id,omitempty"`
	Rating     int                `json:"Rating"`
	Comment    string             `json:"Comment"`
	TouristID  int64              `json:"TouristId"`
	TourID     primitive.ObjectID `json:"TourId"`
	TourDate   time.Time          `json:"TourDate"`
	ReviewDate time.Time          `json:"ReviewDate"`
	Images     ImageListType      `json:"ImageNames"`
}

func (tr *TourReview) BeforeCreate(scope *gorm.DB) error {
	if err := tr.Validate(); err != nil {
		return err
	}

	return nil
}

func (tr *TourReview) Validate() error {
	if tr.Rating == 0 || tr.Rating > 5 {
		return errors.New("invalid rating")
	}
	if tr.TouristID == 0 {
		return errors.New("invalid tourist")
	}
	if tr.TourID == primitive.NilObjectID {
		return errors.New("invalid tour")
	}
	return nil
}

type ImageListType []string

func (l *ImageListType) Scan(value interface{}) error {
	strVal, ok := value.(string)
	if !ok {
		return gorm.ErrInvalidData
	}
	*l = strings.Split(strings.Trim(strVal, "{}"), ",")
	return nil
}

func (l ImageListType) Value() (driver.Value, error) {
	return "{" + strings.Join(l, ",") + "}", nil
}
