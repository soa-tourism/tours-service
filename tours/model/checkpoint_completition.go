package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CheckpointCompletition struct {
	Id               int64     `gorm:"primaryKey"`
	TourExecutionId  int64     `json:"TourExecutionId"`
	CheckpointId     int64     `json:"CheckpointId"`
	CompletitionTime time.Time `json:"CompletitionTime"`
}

func (chCompl *CheckpointCompletition) BeforeCreate(scope *gorm.DB) error {
	if err := chCompl.Validate(); err != nil {
		return err
	}

	uid := uuid.New()
	chCompl.Id = int64(uid.ID())
	chCompl.CompletitionTime = time.Now().UTC()

	return nil
}

func (chCompl *CheckpointCompletition) Validate() error {

	return nil
}

func NewCheckpointCompletition(tourExecutionId, checkpointId int64) *CheckpointCompletition {
	return &CheckpointCompletition{
		TourExecutionId:  tourExecutionId,
		CheckpointId:     checkpointId,
		CompletitionTime: time.Now().UTC(),
	}
}
