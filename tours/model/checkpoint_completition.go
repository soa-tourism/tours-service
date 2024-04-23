package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"

	"gorm.io/gorm"
)

type CheckpointCompletition struct {
	ID               primitive.ObjectID `json:"Id" bson:"_id,omitempty"`
	TourExecutionID  primitive.ObjectID `json:"TourExecutionId"`
	CheckpointID     primitive.ObjectID `json:"CheckpointId"`
	CompletitionTime time.Time          `json:"CompletitionTime"`
}

func (chCompl *CheckpointCompletition) BeforeCreate(scope *gorm.DB) error {
	if err := chCompl.Validate(); err != nil {
		return err
	}

	chCompl.CompletitionTime = time.Now().UTC()

	return nil
}

func (chCompl *CheckpointCompletition) Validate() error {

	return nil
}
