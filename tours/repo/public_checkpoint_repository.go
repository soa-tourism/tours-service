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

type PublicCheckpointRepository struct {
	Collection *mongo.Collection
}

func (repo *PublicCheckpointRepository) FindAll() ([]model.PublicCheckpoint, error) {
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

	var checkpoints []model.PublicCheckpoint
	if err := cursor.All(ctx, &checkpoints); err != nil {
		return nil, err
	}
	return checkpoints, nil
}

func (repo *PublicCheckpointRepository) FindById(id primitive.ObjectID) (model.PublicCheckpoint, error) {
	ctx := context.TODO()
	filter := bson.M{"_id": id}
	var checkpoint model.PublicCheckpoint
	err := repo.Collection.FindOne(ctx, filter).Decode(&checkpoint)
	if err != nil {
		return model.PublicCheckpoint{}, err
	}
	return checkpoint, nil
}

func (repo *PublicCheckpointRepository) Create(checkpoint *model.PublicCheckpoint) (model.PublicCheckpoint, error) {
	ctx := context.TODO()
	result, err := repo.Collection.InsertOne(ctx, checkpoint)
	if err != nil {
		return model.PublicCheckpoint{}, err
	}

	insertedID, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return model.PublicCheckpoint{}, errors.New("failed to get inserted ID")
	}
	checkpoint.ID = insertedID

	return *checkpoint, nil
}

func (repo *PublicCheckpointRepository) Update(checkpoint *model.PublicCheckpoint) (model.PublicCheckpoint, error) {
	ctx := context.TODO()
	filter := bson.M{"_id": checkpoint.ID}
	update := bson.M{"$set": checkpoint}
	_, err := repo.Collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return model.PublicCheckpoint{}, err
	}
	return *checkpoint, nil
}

func (repo *PublicCheckpointRepository) Delete(id primitive.ObjectID) error {
	ctx := context.TODO()
	filter := bson.M{"_id": id}
	_, err := repo.Collection.DeleteOne(ctx, filter)
	return err
}
