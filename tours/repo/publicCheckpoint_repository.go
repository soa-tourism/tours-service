package repo

import (
	"tours/model"

	"gorm.io/gorm"
)

type PublicCheckpointRepository struct {
	DB *gorm.DB
}

func (repo *PublicCheckpointRepository) FindAll() ([]model.PublicCheckpoint, error) {
	var checkpoints []model.PublicCheckpoint
	dbResult := repo.DB.Find(&checkpoints)
	if dbResult.Error != nil {
		return nil, dbResult.Error
	}
	return checkpoints, nil
}

func (repo *PublicCheckpointRepository) FindById(id int64) (model.PublicCheckpoint, error) {
	checkpoint := model.PublicCheckpoint{}
	dbResult := repo.DB.First(&checkpoint, "id = ?", id)
	if dbResult.Error != nil {
		return checkpoint, dbResult.Error
	}
	return checkpoint, nil
}

func (repo *PublicCheckpointRepository) Create(checkpoint *model.PublicCheckpoint) (model.PublicCheckpoint, error) {
	dbResult := repo.DB.Create(checkpoint)
	if dbResult.Error != nil {
		return model.PublicCheckpoint{}, dbResult.Error
	}
	return *checkpoint, nil
}

func (repo *PublicCheckpointRepository) Update(checkpoint *model.PublicCheckpoint) (model.PublicCheckpoint, error) {
	dbResult := repo.DB.Save(checkpoint)
	if dbResult.Error != nil {
		return model.PublicCheckpoint{}, dbResult.Error
	}
	return *checkpoint, nil
}

func (repo *PublicCheckpointRepository) Delete(id int64) error {
	dbResult := repo.DB.Delete(&model.PublicCheckpoint{}, id)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	return nil
}
