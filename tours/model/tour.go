package model

import (
	"errors"
	"github.com/google/uuid"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Tour struct {
	Id          int64          `gorm:"primaryKey"`
	AuthorId    int64          `json:"author_id"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Difficulty  Difficulty     `json:"difficulty"`
	Status      TourStatus     `json:"status"`
	Price       float64        `json:"price"`
	Tags        pq.StringArray `json:"tags" gorm:"type:text[]"`
	Equipment   []Equipment    `json:"equipment" gorm:"many2many:tour_equipment"`
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

type TourStatus int

const (
	Draft TourStatus = iota
	Published
	Archived
)

type Difficulty int

const (
	Easy Difficulty = iota
	Medium
	Hard
)
