package model

import (
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Equipment struct {
	Id          int64  `gorm:"primaryKey"`
	Name        string `json:"Name"`
	Description string `json:"Description"`
}

func (e *Equipment) BeforeCreate(scope *gorm.DB) error {
	if err := e.Validate(); err != nil {
		return err
	}

	uid := uuid.New()
	e.Id = int64(uid.ID())

	return nil
}

func (e *Equipment) Validate() error {
	if e.Name == "" {
		return errors.New("invalid name")
	}
	return nil
}
