package model

import (
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"gorm.io/gorm"
)

type Tour struct {
	ID          primitive.ObjectID `json:"Id" bson:"_id,omitempty"`
	AuthorID    int64              `json:"AuthorId"`
	Name        string             `json:"Name"`
	Description string             `json:"Description"`
	Difficulty  Difficulty         `json:"Difficulty"`
	Status      Status             `json:"Status"`
	Price       float64            `json:"Price"`
	Tags        []string           `json:"Tags"`
	Equipment   []Equipment        `json:"Equipment"`
	Checkpoints []Checkpoint       `json:"Checkpoints"`
}

func (t *Tour) BeforeCreate(scope *gorm.DB) error {
	if err := t.Validate(); err != nil {
		return err
	}

	if t.Equipment == nil {
		t.Equipment = []Equipment{}
	}

	if t.Checkpoints == nil {
		t.Checkpoints = []Checkpoint{}
	}

	return nil
}

func (t *Tour) Validate() error {
	if t.Name == "" {
		return errors.New("invalid name")
	}
	if t.AuthorID == 0 {
		return errors.New("invalid author")
	}
	if t.Price < 0 {
		return errors.New("invalid price")
	}

	return nil
}
