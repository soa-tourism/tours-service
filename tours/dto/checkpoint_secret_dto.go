package dto

import (
	"encoding/json"
	"tours/model"
)

type CheckpointSecretDto struct {
	Description string   `json:"Description"`
	Pictures    []string `json:"Pictures"`
}

func (dto *CheckpointSecretDto) MapToModel() (*model.CheckpointSecret, error) {
	// Marshal the Pictures field to JSON
	_, err := json.Marshal(dto.Pictures)
	if err != nil {
		return nil, err
	}

	// Create a new CheckpointSecret instance
	checkpointSecret := &model.CheckpointSecret{
		Description: dto.Description,
		Pictures:    dto.Pictures,
	}

	return checkpointSecret, nil
}

func CheckpointSecretDtoFromModel(secret model.CheckpointSecret) CheckpointSecretDto {
	return CheckpointSecretDto{
		Description: secret.Description,
		Pictures:    secret.Pictures,
	}
}
