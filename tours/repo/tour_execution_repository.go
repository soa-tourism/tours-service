package repo

import (
	"tours/model"

	"gorm.io/gorm"
)

type TourExecutionRepository struct {
	DB *gorm.DB
}

func (repo *TourExecutionRepository) FindAll() ([]model.TourExecution, error) {
	var TourExecutions []model.TourExecution
	dbResult := repo.DB.Preload("Tour").Find(&TourExecutions)
	if dbResult.Error != nil {
		return nil, dbResult.Error
	}
	return TourExecutions, nil
}

func (repo *TourExecutionRepository) FindById(id int64) (model.TourExecution, error) {
	TourExecution := model.TourExecution{}
	dbResult := repo.DB.Preload("Tour").First(&TourExecution, "id = ?", id)
	if dbResult.Error != nil {
		return TourExecution, dbResult.Error
	}
	return TourExecution, nil
}

func (repo *TourExecutionRepository) Create(TourExecution *model.TourExecution) (model.TourExecution, error) {
	dbResult := repo.DB.Create(TourExecution)
	if dbResult.Error != nil {
		return model.TourExecution{}, dbResult.Error
	}
	return *TourExecution, nil
}

func (repo *TourExecutionRepository) Update(TourExecution *model.TourExecution) (model.TourExecution, error) {
	dbResult := repo.DB.Save(TourExecution)
	if dbResult.Error != nil {
		return model.TourExecution{}, dbResult.Error
	}
	return *TourExecution, nil
}

func (repo *TourExecutionRepository) Delete(id int64) error {
	dbResult := repo.DB.Delete(&model.TourExecution{}, id)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	return nil
}

//// TODO
// func (repo *TourExecutionRepository) FindByTouristAndTour(tourId, touristId int64) (model.TourExecution, error) {
// 	var exactTourExecution model.TourExecution
// 	dbResult := repo.DB.Preload("CompletedCheckpoints").
// 		Preload("Tour", func(db *gorm.DB) *gorm.DB {
// 			return db.Preload("Checkpoints")
// 		}).
// 		Where("tour_id = ? AND tourist_id = ?", tourId, touristId).
// 		First(&exactTourExecution)
// 	if dbResult.Error != nil {
// 		return model.TourExecution{}, dbResult.Error
// 	}
// 	return exactTourExecution, nil
// }
