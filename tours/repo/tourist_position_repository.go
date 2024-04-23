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

type TouristPositionRepository struct {
	Collection *mongo.Collection
}

func (repo *TouristPositionRepository) FindAll() ([]model.TouristPosition, error) {
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

	var positions []model.TouristPosition
	if err := cursor.All(ctx, &positions); err != nil {
		return nil, err
	}
	return positions, nil
}

func (repo *TouristPositionRepository) FindById(id primitive.ObjectID) (model.TouristPosition, error) {
	ctx := context.TODO()
	filter := bson.M{"_id": id}
	var position model.TouristPosition
	err := repo.Collection.FindOne(ctx, filter).Decode(&position)
	if err != nil {
		return model.TouristPosition{}, err
	}
	return position, nil
}

func (repo *TouristPositionRepository) Create(position *model.TouristPosition) (model.TouristPosition, error) {
	ctx := context.TODO()
	result, err := repo.Collection.InsertOne(ctx, position)
	if err != nil {
		return model.TouristPosition{}, err
	}

	insertedID, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return model.TouristPosition{}, errors.New("failed to get inserted ID")
	}
	position.ID = insertedID

	return *position, nil
}

func (repo *TouristPositionRepository) Update(position *model.TouristPosition) (model.TouristPosition, error) {
	ctx := context.TODO()
	filter := bson.M{"_id": position.ID}
	update := bson.M{"$set": position}
	_, err := repo.Collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return model.TouristPosition{}, err
	}
	return *position, nil
}

func (repo *TouristPositionRepository) Delete(id primitive.ObjectID) error {
	ctx := context.TODO()
	filter := bson.M{"_id": id}
	_, err := repo.Collection.DeleteOne(ctx, filter)
	return err
}

func (repo *TouristPositionRepository) FindByCreator(id int64) (model.TouristPosition, error) {
	ctx := context.TODO()
	filter := bson.M{"creatorid": id}
	var position model.TouristPosition
	err := repo.Collection.FindOne(ctx, filter).Decode(&position)
	if err != nil {
		return model.TouristPosition{}, err
	}
	return position, nil
}
