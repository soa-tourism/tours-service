package model

import (
	"errors"
	"github.com/google/uuid"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Tour struct {
	Id          int64          `gorm:"primaryKey"`
	AuthorId    int64          `json:"AuthorId"`
	Name        string         `json:"Name"`
	Description string         `json:"Description"`
	Difficulty  Difficulty     `json:"Difficulty"`
	Status      Status         `json:"Status"`
	Price       float64        `json:"Price"`
	Tags        pq.StringArray `json:"Tags" gorm:"type:text[]"`
	Equipment   []Equipment    `json:"Equipment" gorm:"many2many:tour_equipment"`
}

func (t *Tour) BeforeCreate(scope *gorm.DB) error {
	if err := t.Validate(); err != nil {
		return err
	}

	uid := uuid.New()
	t.Id = int64(uid.ID())

	return nil
}

func (t *Tour) Validate() error {
	if t.Name == "" {
		return errors.New("invalid name")
	}
	if t.AuthorId == 0 {
		return errors.New("invalid author")
	}
	if t.Price < 0 {
		return errors.New("invalid price")
	}

	return nil
}
