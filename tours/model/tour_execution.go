package model

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TourExecution struct {
	Id                   int64 `gorm:"primaryKey"`
	TouristId            int64 `json:"TouristId"`
	TourId               int64 `json:"TourId"`
	Tour                 Tour
	Start                time.Time                `json:"Start"`
	LastActivity         time.Time                `json:"LastActivity"`
	ExecutionStatus      ExecutionStatus          `json:"ExecutionStatus"`
	CompletedCheckpoints []CheckpointCompletition `json:"CompletedCheckpoints" gorm:"foreignKey:TourExecutionId"` //
	// Changes              []abstractions.DomainEvent `json:"-" gorm:"type:jsonb"`
	// Version              int64                      `json:"-"`
}

func (t *TourExecution) BeforeCreate(scope *gorm.DB) error {
	if err := t.Validate(); err != nil {
		return err
	}

	if t.CompletedCheckpoints == nil {
		t.CompletedCheckpoints = []CheckpointCompletition{}
	}

	uid := uuid.New()
	t.Id = int64(uid.ID())

	return nil
}

func (execution *TourExecution) Validate() error {
	if execution.Start.After(execution.LastActivity) {
		return errors.New("invalid start time")
	}

	return nil
}

//// TODO event-sourcing for TourExectution
