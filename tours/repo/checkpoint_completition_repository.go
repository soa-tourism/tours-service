package repo

import (
	"tours/model"

	"gorm.io/gorm"
)

type CheckpointCompletitionRepository struct {
	DB *gorm.DB
}

func (repo *CheckpointCompletitionRepository) FindAll() ([]model.CheckpointCompletition, error) {
	var CheckpointCompletitions []model.CheckpointCompletition
	dbResult := repo.DB.Find(&CheckpointCompletitions)
	if dbResult.Error != nil {
		return nil, dbResult.Error
	}
	return CheckpointCompletitions, nil
}

func (repo *CheckpointCompletitionRepository) FindById(id int64) (model.CheckpointCompletition, error) {
	CheckpointCompletition := model.CheckpointCompletition{}
	dbResult := repo.DB.First(&CheckpointCompletition, "id = ?", id)
	if dbResult.Error != nil {
		return CheckpointCompletition, dbResult.Error
	}
	return CheckpointCompletition, nil
}

func (repo *CheckpointCompletitionRepository) Create(CheckpointCompletition *model.CheckpointCompletition) (model.CheckpointCompletition, error) {
	dbResult := repo.DB.Create(CheckpointCompletition)
	if dbResult.Error != nil {
		return model.CheckpointCompletition{}, dbResult.Error
	}
	return *CheckpointCompletition, nil
}

func (repo *CheckpointCompletitionRepository) Update(CheckpointCompletition *model.CheckpointCompletition) (model.CheckpointCompletition, error) {
	dbResult := repo.DB.Save(CheckpointCompletition)
	if dbResult.Error != nil {
		return model.CheckpointCompletition{}, dbResult.Error
	}
	return *CheckpointCompletition, nil
}

func (repo *CheckpointCompletitionRepository) Delete(id int64) error {
	dbResult := repo.DB.Delete(&model.CheckpointCompletition{}, id)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	return nil
}

func (repo *CheckpointCompletitionRepository) FindByExecution(id int64) ([]model.CheckpointCompletition, error) {
	var CheckpointCompletitions []model.CheckpointCompletition
	dbResult := repo.DB.Where("tour_execution_id = ?", id).Find(&CheckpointCompletitions)
	if dbResult.Error != nil {
		return nil, dbResult.Error
	}
	return CheckpointCompletitions, nil
}
