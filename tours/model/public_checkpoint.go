package model

import (
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PublicCheckpoint struct {
	Id          int64    `gorm:"primaryKey"`
	Longitude   float64  `json:"Longitude"`
	Latitude    float64  `json:"Latitude"`
	Name        string   `json:"Name"`
	Description string   `json:"Description"`
	Pictures    []string `json:"Pictures" gorm:"type:text[]"`
}

func (pch *PublicCheckpoint) BeforeCreate(scope *gorm.DB) error {
	if err := pch.Validate(); err != nil {
		return err
	}

	uid := uuid.New()
	pch.Id = int64(uid.ID())

	return nil
}

func (pch *PublicCheckpoint) Validate() error {
	if pch.Name == "" {
		return errors.New("invalid public checkpoint name")
	}
	if pch.Longitude < 0 || pch.Longitude > 180 {
		return errors.New("invalid public checkpoint longitude")
	}
	if pch.Latitude < -90 || pch.Latitude > 90 {
		return errors.New("invalid public checkpoint latitude")
	}
	if len(pch.Pictures) == 0 {
		return errors.New("invalid public checkpoint pictures")
	}

	return nil
}
