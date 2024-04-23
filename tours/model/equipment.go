package model

import (
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gorm.io/gorm"
)

type Equipment struct {
	ID          primitive.ObjectID `json:"Id" bson:"_id,omitempty"`
	Name        string             `json:"Name"`
	Description string             `json:"Description"`
}

func (e *Equipment) BeforeCreate(scope *gorm.DB) error {
	if err := e.Validate(); err != nil {
		return err
	}

	return nil
}

func (e *Equipment) Validate() error {
	if e.Name == "" {
		return errors.New("invalid name")
	}
	return nil
}
