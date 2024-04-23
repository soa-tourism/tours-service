package model

import (
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"

	"gorm.io/gorm"
)

type TourExecution struct {
	ID                   primitive.ObjectID       `json:"Id" bson:"_id,omitempty"`
	TouristID            int64                    `json:"TouristId"`
	TourID               primitive.ObjectID       `json:"TourId"`
	Start                time.Time                `json:"Start"`
	LastActivity         time.Time                `json:"LastActivity"`
	ExecutionStatus      ExecutionStatus          `json:"ExecutionStatus"`
	CompletedCheckpoints []CheckpointCompletition `json:"CompletedCheckpoints"`
}

func (t *TourExecution) BeforeCreate(scope *gorm.DB) error {
	if err := t.Validate(); err != nil {
		return err
	}

	if t.CompletedCheckpoints == nil {
		t.CompletedCheckpoints = []CheckpointCompletition{}
	}

	return nil
}

func (execution *TourExecution) Validate() error {
	if execution.Start.After(execution.LastActivity) {
		return errors.New("invalid start time")
	}

	return nil
}
