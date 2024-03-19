package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"

	"github.com/lib/pq"
)

type CheckpointSecret struct {
	Description string         `json:"Description"`
	Pictures    pq.StringArray `json:"Pictures" gorm:"type:text[]"`
}

// Implementing the Scanner interface for CheckpointSecret
func (cs *CheckpointSecret) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	// Check if the value is a byte slice
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("Scan source is not []byte")
	}
	// Unmarshal the byte slice to CheckpointSecret
	return json.Unmarshal(bytes, cs)
}

// Implementing the Valuer interface for CheckpointSecret
func (cs CheckpointSecret) Value() (driver.Value, error) {
	if cs.Description == "" && len(cs.Pictures) == 0 {
		return nil, nil
	}
	// Marshal the CheckpointSecret to JSON
	return json.Marshal(cs)
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
