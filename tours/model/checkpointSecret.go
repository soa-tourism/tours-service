package model

import (
	"errors"
)

type CheckpointSecret struct {
	Description string   `json:"Description"`
	Pictures    []string `json:"Pictures" gorm:"type:text[]"`
}

func NewCheckpointSecret(description string, pictures []string) (*CheckpointSecret, error) {
	if description == "" {
		return nil, errors.New("invalid checkpoint secret description")
	}

	return &CheckpointSecret{
		Description: description,
		Pictures:    pictures,
	}, nil
}
