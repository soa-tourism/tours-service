package repo

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"tours/model"
)

type TourRepository struct {
	Collection *mongo.Collection
}

func (repo *TourRepository) FindAll() ([]model.Tour, error) {
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

	var tours []model.Tour
	if err := cursor.All(ctx, &tours); err != nil {
		return nil, err
	}
	return tours, nil
}

func (repo *TourRepository) FindById(id primitive.ObjectID) (*model.Tour, error) {
	ctx := context.TODO()
	filter := bson.M{"_id": id}
	var tour model.Tour
	err := repo.Collection.FindOne(ctx, filter).Decode(&tour)
	if err != nil {
		return nil, err
	}
	return &tour, nil
}

func (repo *TourRepository) Create(tour *model.Tour) (*model.Tour, error) {
	ctx := context.TODO()
	result, err := repo.Collection.InsertOne(ctx, tour)
	if err != nil {
		return nil, err
	}

	insertedID, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, errors.New("failed to get inserted ID")
	}
	tour.ID = insertedID

	return tour, nil
}

func (repo *TourRepository) Update(tour *model.Tour) (*model.Tour, error) {
	ctx := context.TODO()
	filter := bson.M{"_id": tour.ID}
	update := bson.M{"$set": tour}
	_, err := repo.Collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}
	return tour, nil
}

func (repo *TourRepository) Delete(id primitive.ObjectID) error {
	ctx := context.TODO()
	filter := bson.M{"_id": id}
	_, err := repo.Collection.DeleteOne(ctx, filter)
	return err
}

func (repo *TourRepository) FindByAuthor(authorID int64) ([]model.Tour, error) {
	ctx := context.TODO()
	filter := bson.M{"authorid": authorID}
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

	var tours []model.Tour
	if err := cursor.All(ctx, &tours); err != nil {
		return nil, err
	}
	return tours, nil
}

func (repo *TourRepository) FindAllPublished() ([]model.Tour, error) {
	ctx := context.TODO()
	filter := bson.M{"status": model.Published}
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

	var tours []model.Tour
	if err := cursor.All(ctx, &tours); err != nil {
		return nil, err
	}
	return tours, nil
}

func (repo *TourRepository) AddEquipment(tourID primitive.ObjectID, equipmentID primitive.ObjectID) error {
	ctx := context.TODO()
	update := bson.M{"$addToSet": bson.M{"equipment": equipmentID}}
	filter := bson.M{"_id": tourID}
	_, err := repo.Collection.UpdateOne(ctx, filter, update)
	return err
}

func (repo *TourRepository) RemoveEquipment(tourID primitive.ObjectID, equipmentID primitive.ObjectID) error {
	ctx := context.TODO()
	update := bson.M{"$pull": bson.M{"equipment": bson.M{"_id": equipmentID}}}
	filter := bson.M{"_id": tourID}
	_, err := repo.Collection.UpdateOne(ctx, filter, update)
	return err
}
