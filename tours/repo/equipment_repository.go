package repo

import (
	"tours/model"

	"gorm.io/gorm"
)

type EquipmentRepository struct {
	DatabaseConnection *gorm.DB
}

func (repo *EquipmentRepository) FindAll() ([]model.Equipment, error) {
	var equipments []model.Equipment
	dbResult := repo.DatabaseConnection.Find(&equipments)
	if dbResult.Error != nil {
		return nil, dbResult.Error
	}
	return equipments, nil
}

func (repo *EquipmentRepository) FindById(id int64) (model.Equipment, error) {
	equipment := model.Equipment{}
	dbResult := repo.DatabaseConnection.First(&equipment, "id = ?", id)
	if dbResult.Error != nil {
		return equipment, dbResult.Error
	}
	return equipment, nil
}

func (repo *EquipmentRepository) Create(equipment *model.Equipment) (*model.Equipment, error) {
	dbResult := repo.DatabaseConnection.Create(equipment)
	if dbResult.Error != nil {
		return nil, dbResult.Error
	}
	return equipment, nil
}

func (repo *EquipmentRepository) Update(equipment *model.Equipment) (*model.Equipment, error) {
	dbResult := repo.DatabaseConnection.Save(equipment)
	if dbResult.Error != nil {
		return nil, dbResult.Error
	}
	return equipment, nil
}

func (repo *EquipmentRepository) Delete(id int64) error {
	dbResult := repo.DatabaseConnection.Delete(&model.Equipment{}, id)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	return nil
}
