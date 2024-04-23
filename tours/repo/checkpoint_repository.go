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

type CheckpointRepository struct {
	Collection *mongo.Collection
}

func (repo *CheckpointRepository) FindAll() ([]model.Checkpoint, error) {
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

	var checkpoints []model.Checkpoint
	if err := cursor.All(ctx, &checkpoints); err != nil {
		return nil, err
	}
	return checkpoints, nil
}

func (repo *CheckpointRepository) FindById(id primitive.ObjectID) (model.Checkpoint, error) {
	ctx := context.TODO()
	filter := bson.M{"_id": id}
	var checkpoint model.Checkpoint
	err := repo.Collection.FindOne(ctx, filter).Decode(&checkpoint)
	if err != nil {
		return model.Checkpoint{}, err
	}
	return checkpoint, nil
}

func (repo *CheckpointRepository) Create(checkpoint *model.Checkpoint) (model.Checkpoint, error) {
	ctx := context.TODO()
	result, err := repo.Collection.InsertOne(ctx, checkpoint)
	if err != nil {
		return model.Checkpoint{}, err
	}

	insertedID, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return model.Checkpoint{}, errors.New("failed to get inserted ID")
	}
	checkpoint.ID = insertedID

	return *checkpoint, nil
}

func (repo *CheckpointRepository) Update(checkpoint *model.Checkpoint) (model.Checkpoint, error) {
	ctx := context.TODO()
	filter := bson.M{"_id": checkpoint.ID}
	update := bson.M{"$set": checkpoint}
	_, err := repo.Collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return model.Checkpoint{}, err
	}
	return *checkpoint, nil
}

func (repo *CheckpointRepository) Delete(id primitive.ObjectID) error {
	ctx := context.TODO()
	filter := bson.M{"_id": id}
	_, err := repo.Collection.DeleteOne(ctx, filter)
	return err
}

func (repo *CheckpointRepository) FindByTour(id primitive.ObjectID) ([]model.Checkpoint, error) {
	ctx := context.TODO()
	filter := bson.M{"tourid": id}
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

	var checkpoints []model.Checkpoint
	if err := cursor.All(ctx, &checkpoints); err != nil {
		return nil, err
	}
	return checkpoints, nil
}

func (repo *CheckpointRepository) UpdateCheckpointSecret(checkpoint *model.Checkpoint, secret *model.CheckpointSecret) (model.Checkpoint, error) {
	err := checkpoint.UpdateCheckpointSecret(secret.Description, secret.Pictures)
	if err != nil {
		return model.Checkpoint{}, err
	}
	ctx := context.TODO()
	filter := bson.M{"_id": checkpoint.ID}
	update := bson.M{"$set": checkpoint}
	_, err = repo.Collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return model.Checkpoint{}, err
	}
	return *checkpoint, nil
}
