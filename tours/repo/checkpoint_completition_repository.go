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

type CheckpointCompletitionRepository struct {
	Collection *mongo.Collection
}

func (repo *CheckpointCompletitionRepository) FindAll() ([]model.CheckpointCompletition, error) {
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

	var completions []model.CheckpointCompletition
	if err := cursor.All(ctx, &completions); err != nil {
		return nil, err
	}

	return completions, nil
}

func (repo *CheckpointCompletitionRepository) FindById(id primitive.ObjectID) (*model.CheckpointCompletition, error) {
	ctx := context.TODO()
	filter := bson.M{"_id": id}

	var completion model.CheckpointCompletition
	err := repo.Collection.FindOne(ctx, filter).Decode(&completion)
	if err != nil {
		return nil, err
	}

	return &completion, nil
}

func (repo *CheckpointCompletitionRepository) Create(completion *model.CheckpointCompletition) error {
	ctx := context.TODO()
	result, err := repo.Collection.InsertOne(ctx, completion)
	if err != nil {
		return err
	}

	insertedID, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return errors.New("failed to get inserted ID")
	}
	completion.ID = insertedID

	return nil
}

func (repo *CheckpointCompletitionRepository) Update(completion *model.CheckpointCompletition) error {
	ctx := context.TODO()
	filter := bson.M{"_id": completion.ID}
	update := bson.M{"$set": completion}

	_, err := repo.Collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (repo *CheckpointCompletitionRepository) Delete(id primitive.ObjectID) error {
	ctx := context.TODO()
	filter := bson.M{"_id": id}

	_, err := repo.Collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	return nil
}

func (repo *CheckpointCompletitionRepository) FindByExecution(id primitive.ObjectID) ([]model.CheckpointCompletition, error) {
	ctx := context.TODO()
	filter := bson.M{"tourexecutionid": id}

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

	var completions []model.CheckpointCompletition
	if err := cursor.All(ctx, &completions); err != nil {
		return nil, err
	}

	return completions, nil
}
