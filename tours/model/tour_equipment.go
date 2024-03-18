package model

type TourEquipment struct {
	TourId      int64 `gorm:"primaryKey"`
	EquipmentId int64 `gorm:"primaryKey"`
}
