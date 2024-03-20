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

func (repo *TourRepository) FindById(id int64) (*model.Tour, error) {
	tour := model.Tour{}
	dbResult := repo.DB.Preload("Equipment").Preload("Checkpoints").First(&tour, id)
	if dbResult.Error != nil {
		return &tour, dbResult.Error
	}
	return &tour, nil
}

func (repo *TourRepository) Create(tour *model.Tour) (model.Tour, error) {
	dbResult := repo.DB.Create(tour)
	if dbResult.Error != nil {
		return model.Tour{}, dbResult.Error
	}
	return *tour, nil
}

func (repo *TourRepository) Update(tour *model.Tour) (model.Tour, error) {
	tour.Equipment = nil
	tour.Checkpoints = nil
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

func (repo *TourRepository) AddEquipment(tourId int64, equipmentId int64) error {
	if err := repo.DB.Create(
		&model.TourEquipment{
			TourId:      tourId,
			EquipmentId: equipmentId,
		}).Error; err != nil {
		return err
	}

	return nil
}

func (repo *TourRepository) RemoveEquipment(tourId int64, equipmentId int64) error {
	if err := repo.DB.Delete(
		&model.TourEquipment{},
		"tour_id = ?  AND equipment_id = ?",
		tourId,
		equipmentId,
	).Error; err != nil {
		return err
	}

	return nil
}
