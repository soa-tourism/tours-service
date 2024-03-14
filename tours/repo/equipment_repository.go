package repo

import (
	"tours/model"

	"gorm.io/gorm"
)

type EquipmentRepository struct {
	DB *gorm.DB
}

func (repo *EquipmentRepository) FindAll() ([]model.Equipment, error) {
	var equipments []model.Equipment
	dbResult := repo.DB.Find(&equipments)
	if dbResult.Error != nil {
		return nil, dbResult.Error
	}
	return equipments, nil
}

func (repo *EquipmentRepository) FindById(id int64) (model.Equipment, error) {
	equipment := model.Equipment{}
	dbResult := repo.DB.First(&equipment, "id = ?", id)
	if dbResult.Error != nil {
		return equipment, dbResult.Error
	}
	return equipment, nil
}

func (repo *EquipmentRepository) Create(equipment *model.Equipment) (*model.Equipment, error) {
	dbResult := repo.DB.Create(equipment)
	if dbResult.Error != nil {
		return nil, dbResult.Error
	}
	return equipment, nil
}

func (repo *EquipmentRepository) Update(equipment *model.Equipment) (*model.Equipment, error) {
	dbResult := repo.DB.Save(equipment)
	if dbResult.Error != nil {
		return nil, dbResult.Error
	}
	return equipment, nil
}

func (repo *EquipmentRepository) Delete(id int64) error {
	dbResult := repo.DB.Delete(&model.Equipment{}, id)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	return nil
}

func (repo *EquipmentRepository) GetAvailable(ids []int64) ([]model.Equipment, error) {
	var dbResult []model.Equipment

	if len(ids) == 0 {
		err := repo.DB.Find(&dbResult).Error
		if err != nil {
			return nil, err
		}
		return dbResult, nil
	}

	err := repo.DB.Where("id NOT IN (?)", ids).Find(&dbResult).Error
	if err != nil {
		return nil, err
	}

	return dbResult, nil
}
