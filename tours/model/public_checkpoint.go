package model

import (
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

type PublicCheckpoint struct {
	ID          primitive.ObjectID `json:"Id" bson:"_id,omitempty"`
	Longitude   float64            `json:"Longitude"`
	Latitude    float64            `json:"Latitude"`
	Name        string             `json:"Name"`
	Description string             `json:"Description"`
	Pictures    pq.StringArray     `json:"Pictures" gorm:"type:text[]"`
}

func (pch *PublicCheckpoint) BeforeCreate(scope *gorm.DB) error {
	if err := pch.Validate(); err != nil {
		return err
	}

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
