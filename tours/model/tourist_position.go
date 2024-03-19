package model

import (
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TouristPosition struct {
	Id        int64   `gorm:"primaryKey"`
	CreatorId int64   `json:"CreatorId"`
	Longitude float64 `json:"Longitude"`
	Latitude  float64 `json:"Latitude"`
}

func (pos *TouristPosition) BeforeCreate(scope *gorm.DB) error {
	if err := pos.Validate(); err != nil {
		return err
	}

	uid := uuid.New()
	pos.Id = int64(uid.ID())

	return nil
}

func (pos *TouristPosition) Validate() error {
	if pos.CreatorId == 0 {
		return errors.New("invalid tourist position creatorId")
	}
	if pos.Longitude < 0 || pos.Longitude > 180 {
		return errors.New("invalid tourist position longitude")
	}
	if pos.Latitude < -90 || pos.Latitude > 90 {
		return errors.New("invalid tourist position latitude")
	}

	return nil
}
