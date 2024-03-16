package model

import (
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Checkpoint struct {
	Id                    int64 `gorm:"primaryKey"`
	TourId                int64 `json:"TourId"`
	Tour                  Tour
	AuthorId              int64            `json:"AuthorId"`
	Longitude             float64          `json:"Longitude"`
	Latitude              float64          `json:"Latitude"`
	Name                  string           `json:"Name"`
	Description           string           `json:"Description"`
	Pictures              []string         `json:"Pictures" gorm:"type:text[]"`
	RequiredTimeInSeconds float64          `json:"RequiredTimeInSeconds"`
	CheckpointSecret      CheckpointSecret `json:"CheckpointSecret"`
	IsSecretPrerequisite  bool             `json:"IsSecretPrerequisite"`
	EncounterId           int64            `json:"EncounterId"`
}

func (ch *Checkpoint) BeforeCreate(scope *gorm.DB) error {
	if err := ch.Validate(); err != nil {
		return err
	}

	uid := uuid.New()
	ch.Id = int64(uid.ID())

	return nil
}

func (ch *Checkpoint) Validate() error {
	if ch.TourId == 0 {
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
	return c.AuthorId == int64(userId)
}
