package repo

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"tours/model"
)

type TourExecutionRepository struct {
	Collection *mongo.Collection
}

func (repo *TourExecutionRepository) FindAll() ([]model.TourExecution, error) {
	ctx := context.TODO()
	cursor, err := repo.Collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {
			fmt.Printf("Error closing cursor: %v\n", err)
		}
	}(cursor, ctx)

	var tourExecutions []model.TourExecution
	if err := cursor.All(ctx, &tourExecutions); err != nil {
		return nil, err
	}
	return tourExecutions, nil
}

func (repo *TourExecutionRepository) FindById(id primitive.ObjectID) (model.TourExecution, error) {
	ctx := context.TODO()
	filter := bson.M{"_id": id}
	var tourExecution model.TourExecution
	err := repo.Collection.FindOne(ctx, filter).Decode(&tourExecution)
	if err != nil {
		return model.TourExecution{}, err
	}
	return tourExecution, nil
}

func (repo *TourExecutionRepository) Create(tourExecution *model.TourExecution) (model.TourExecution, error) {
	ctx := context.TODO()
	result, err := repo.Collection.InsertOne(ctx, tourExecution)
	if err != nil {
		return model.TourExecution{}, err
	}

	insertedID, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return model.TourExecution{}, errors.New("failed to get inserted ID")
	}
	tourExecution.ID = insertedID

	return *tourExecution, nil
}

func (repo *TourExecutionRepository) Update(tourExecution *model.TourExecution) (model.TourExecution, error) {
	ctx := context.TODO()
	filter := bson.M{"_id": tourExecution.ID}
	update := bson.M{"$set": tourExecution}
	_, err := repo.Collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return model.TourExecution{}, err
	}
	return *tourExecution, nil
}

func (repo *TourExecutionRepository) Delete(id primitive.ObjectID) error {
	ctx := context.TODO()
	filter := bson.M{"_id": id}
	_, err := repo.Collection.DeleteOne(ctx, filter)
	return err
}

func (repo *TourExecutionRepository) FindByTouristAndTour(tourId primitive.ObjectID, touristId int64) ([]model.TourExecution, error) {
	ctx := context.TODO()
	filter := bson.M{"tourid": tourId, "touristid": touristId}
	cursor, err := repo.Collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {
			fmt.Printf("Error closing cursor: %v\n", err)
		}
	}(cursor, ctx)

	var executions []model.TourExecution
	if err := cursor.All(ctx, &executions); err != nil {
		return nil, err
	}
	return executions, nil
}

func (repo *TourExecutionRepository) FindActiveByTouristAndTour(tourId primitive.ObjectID, touristId int64) (model.TourExecution, error) {
	ctx := context.TODO()
	filter := bson.M{"executionstatus": 2, "tourid": tourId, "touristid": touristId}
	var exactTourExecution model.TourExecution
	err := repo.Collection.FindOne(ctx, filter).Decode(&exactTourExecution)
	if err != nil {
		return model.TourExecution{}, err
	}
	return exactTourExecution, nil
}
