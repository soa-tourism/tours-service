package model

import (
	"database/sql/driver"
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"strings"
	"time"
)

type TourReview struct {
	Id         int64         `gorm:"primaryKey"`
	Rating     int           `json:"Rating"`
	Comment    string        `json:"Comment"`
	TouristId  int64         `json:"TouristId"`
	TourId     int64         `json:"TourId"`
	TourDate   time.Time     `json:"TourDate"`
	ReviewDate time.Time     `json:"ReviewDate"`
	Images     ImageListType `json:"ImageNames" gorm:"type:text[]"`
}

func (tr *TourReview) BeforeCreate(scope *gorm.DB) error {
	if err := tr.Validate(); err != nil {
		return err
	}

	uid := uuid.New()
	tr.Id = int64(uid.ID())

	return nil
}

func (tr *TourReview) Validate() error {
	if tr.Rating == 0 || tr.Rating > 5 {
		return errors.New("invalid rating")
	}
	if tr.TouristId == 0 {
		return errors.New("invalid tourist")
	}
	if tr.TourId == 0 {
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
