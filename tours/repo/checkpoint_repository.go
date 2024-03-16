package repo

import (
	"tours/model"

	"gorm.io/gorm"
)

type CheckpointRepository struct {
	DB *gorm.DB
}

func (repo *CheckpointRepository) FindAll() ([]model.Checkpoint, error) {
	var checkpoints []model.Checkpoint
	dbResult := repo.DB.Find(&checkpoints)
	if dbResult.Error != nil {
		return nil, dbResult.Error
	}
	return checkpoints, nil
}

func (repo *CheckpointRepository) FindById(id int64) (model.Checkpoint, error) {
	checkpoint := model.Checkpoint{}
	dbResult := repo.DB.First(&checkpoint, "id = ?", id)
	if dbResult.Error != nil {
		return checkpoint, dbResult.Error
	}
	return checkpoint, nil
}

func (repo *CheckpointRepository) Create(checkpoint *model.Checkpoint) (model.Checkpoint, error) {
	dbResult := repo.DB.Create(checkpoint)
	if dbResult.Error != nil {
		return model.Checkpoint{}, dbResult.Error
	}
	return *checkpoint, nil
}

func (repo *CheckpointRepository) Update(checkpoint *model.Checkpoint) (model.Checkpoint, error) {
	dbResult := repo.DB.Save(checkpoint)
	if dbResult.Error != nil {
		return model.Checkpoint{}, dbResult.Error
	}
	return *checkpoint, nil
}

func (repo *CheckpointRepository) Delete(id int64) error {
	dbResult := repo.DB.Delete(&model.Checkpoint{}, id)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	return nil
}

func (repo *CheckpointRepository) FindByTour(id int64) ([]model.Checkpoint, error) {
	var checkpoints []model.Checkpoint
	dbResult := repo.DB.Where("tour_id = ?", id).Find(&checkpoints)
	if dbResult.Error != nil {
		return nil, dbResult.Error
	}
	return checkpoints, nil
}
