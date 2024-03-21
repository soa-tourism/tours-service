package repo

import (
	"tours/model"

	"gorm.io/gorm"
)

type TouristPositionRepository struct {
	DB *gorm.DB
}

func (repo *TouristPositionRepository) FindAll() ([]model.TouristPosition, error) {
	var positions []model.TouristPosition
	dbResult := repo.DB.Find(&positions)
	if dbResult.Error != nil {
		return nil, dbResult.Error
	}
	return positions, nil
}

func (repo *TouristPositionRepository) FindById(id int64) (model.TouristPosition, error) {
	positions := model.TouristPosition{}
	dbResult := repo.DB.First(&positions, "creator_id = ?", id)
	if dbResult.Error != nil {
		return positions, dbResult.Error
	}
	return positions, nil
}

func (repo *TouristPositionRepository) Create(positions *model.TouristPosition) (model.TouristPosition, error) {
	dbResult := repo.DB.Create(positions)
	if dbResult.Error != nil {
		return model.TouristPosition{}, dbResult.Error
	}
	return *positions, nil
}

func (repo *TouristPositionRepository) Update(positions *model.TouristPosition) (model.TouristPosition, error) {
	dbResult := repo.DB.Save(positions)
	if dbResult.Error != nil {
		return model.TouristPosition{}, dbResult.Error
	}
	return *positions, nil
}

func (repo *TouristPositionRepository) Delete(id int64) error {
	dbResult := repo.DB.Delete(&model.TouristPosition{}, id)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	return nil
}

func (repo *TouristPositionRepository) FindByCreator(id int64) (model.TouristPosition, error) {
	pos := model.TouristPosition{}
	dbResult := repo.DB.First(&pos, "creator_id = ?", id)
	if dbResult.Error != nil {
		return pos, dbResult.Error
	}
	return pos, nil
}
