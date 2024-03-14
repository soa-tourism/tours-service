package service

import (
	"fmt"
	"tours/model"
	"tours/repo"
)

type EquipmentService struct {
	EquipmentRepo *repo.EquipmentRepository
}

func (service *EquipmentService) FindAllEquipment() ([]model.Equipment, error) {
	equipments, err := service.EquipmentRepo.FindAll()
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve equipment: %v", err)
	}
	return equipments, nil
}

func (service *EquipmentService) FindEquipment(id int64) (*model.Equipment, error) {
	equipment, err := service.EquipmentRepo.FindById(id)
	if err != nil {
		return nil, fmt.Errorf("equipment with ID %d not found", id)
	}
	return &equipment, nil
}

func (service *EquipmentService) CreateEquipment(equipment *model.Equipment) (*model.Equipment, error) {
	equip, err := service.EquipmentRepo.Create(equipment)
	if err != nil {
		return equip, fmt.Errorf("failed to create equipment: %v", err)
	}
	return equip, nil
}

func (service *EquipmentService) UpdateEquipment(equipment *model.Equipment) (*model.Equipment, error) {
	equip, err := service.EquipmentRepo.Update(equipment)
	if err != nil {
		return equip, fmt.Errorf("failed to update equipment: %v", err)
	}
	return equip, nil
}

func (service *EquipmentService) DeleteEquipment(id int64) error {
	err := service.EquipmentRepo.Delete(id)
	if err != nil {
		return fmt.Errorf("failed to delete equipment with ID %d: %v", id, err)
	}
	return nil
}

func (service *EquipmentService) GetAvailableEquipment(tourId int64, equipmentIds []int64) ([]model.Equipment, error) {
	availableEquipment, err := service.EquipmentRepo.GetAvailable(equipmentIds)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve equipment: %v", err)
	}
	return availableEquipment, nil
}
