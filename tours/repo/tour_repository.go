package repo

import (
	"tours/model"

	"gorm.io/gorm"
)

type TourRepository struct {
	DB *gorm.DB
}

func (repo *TourRepository) FindAll() ([]model.Tour, error) {
	var tours []model.Tour
	dbResult := repo.DB.Find(&tours)
	if dbResult.Error != nil {
		return nil, dbResult.Error
	}
	return tours, nil
}

func (repo *TourRepository) FindById(id int64) (model.Tour, error) {
	tour := model.Tour{}
	dbResult := repo.DB.First(&tour, "id = ?", id)
	if dbResult.Error != nil {
		return tour, dbResult.Error
	}
	return tour, nil
}

func (repo *TourRepository) Create(tour *model.Tour) (model.Tour, error) {
	dbResult := repo.DB.Create(tour)
	if dbResult.Error != nil {
		return model.Tour{}, dbResult.Error
	}
	return *tour, nil
}

func (repo *TourRepository) Update(tour *model.Tour) (model.Tour, error) {
	dbResult := repo.DB.Save(tour)
	if dbResult.Error != nil {
		return model.Tour{}, dbResult.Error
	}
	return *tour, nil
}

func (repo *TourRepository) Delete(id int64) error {
	dbResult := repo.DB.Delete(&model.Tour{}, id)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	return nil
}

func (repo *TourRepository) FindByAuthor(id int64) ([]model.Tour, error) {
	var tours []model.Tour
	dbResult := repo.DB.Where("author_id = ?", id).Find(&tours)
	if dbResult.Error != nil {
		return nil, dbResult.Error
	}
	return tours, nil
}
