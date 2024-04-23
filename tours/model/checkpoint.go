package model

import (
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Checkpoint struct {
	ID                    primitive.ObjectID `json:"Id" bson:"_id,omitempty"`
	TourID                primitive.ObjectID `json:"TourId"`
	AuthorID              int64              `json:"AuthorId"`
	Longitude             float64            `json:"Longitude"`
	Latitude              float64            `json:"Latitude"`
	Name                  string             `json:"Name"`
	Description           string             `json:"Description"`
	Pictures              pq.StringArray     `json:"Pictures" gorm:"type:text[]"`
	RequiredTimeInSeconds float64            `json:"RequiredTimeInSeconds"`
	IsSecretPrerequisite  bool               `json:"IsSecretPrerequisite"`
	EncounterID           int64              `json:"EncounterId"`
	CheckpointSecret      CheckpointSecret   `json:"CheckpointSecret" gorm:"type:json"`
}

func (ch *Checkpoint) BeforeCreate(scope *gorm.DB) error {
	if err := ch.Validate(); err != nil {
		return err
	}

	return nil
}

func (ch *Checkpoint) Validate() error {
	if ch.TourID == primitive.NilObjectID {
		return errors.New("invalid checkpoint tourId")
	}
	if ch.Longitude < 0 || ch.Longitude > 180 {
		return errors.New("invalid checkpoint longitude")
	}
	if ch.Latitude < -90 || ch.Latitude > 90 {
		return errors.New("invalid checkpoint latitude")
	}
	if ch.Name == "" {
		return errors.New("invalid checkpoint name")
	}
	if ch.RequiredTimeInSeconds <= 0 {
		return errors.New("invalid checkpoint RequiredTimeInSeconds")
	}
	if ch.Pictures == nil || len(ch.Pictures) == 0 {
		return errors.New("invalid checkpoint pictures")
	}

	return nil
}

func (c *Checkpoint) UpdateCheckpointSecret(description string, pictures []string) error {
	newSecret, err := NewCheckpointSecret(description, pictures)
	if err != nil {
		return err
	}
	c.CheckpointSecret = *newSecret

	return nil
}

func (c *Checkpoint) DeleteCheckpointSecret() {
	c.CheckpointSecret = CheckpointSecret{} // Reset the checkpoint secret
}

func (c *Checkpoint) IsAuthor(userId int) bool {
	return c.AuthorID == int64(userId)
}
