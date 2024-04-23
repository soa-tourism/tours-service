package repo

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"tours/model"
)

type EquipmentRepository struct {
	Collection *mongo.Collection
}

func (repo *EquipmentRepository) FindAll() ([]model.Equipment, error) {
	ctx := context.TODO()
	cursor, err := repo.Collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var equipment []model.Equipment
	if err := cursor.All(ctx, &equipment); err != nil {
		return nil, err
	}
	return equipment, nil
}

func (repo *EquipmentRepository) FindById(id primitive.ObjectID) (*model.Equipment, error) {
	ctx := context.TODO()
	filter := bson.M{"_id": id}
	var equipment model.Equipment
	err := repo.Collection.FindOne(ctx, filter).Decode(&equipment)
	if err != nil {
		return nil, err
	}
	return &equipment, nil
}

func (repo *EquipmentRepository) Create(equipment *model.Equipment) (*model.Equipment, error) {
	ctx := context.TODO()
	result, err := repo.Collection.InsertOne(ctx, equipment)
	if err != nil {
		return nil, err
	}

	insertedID, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, errors.New("failed to get inserted ID")
	}
	equipment.ID = insertedID

	return equipment, nil
}

func (repo *EquipmentRepository) Update(equipment *model.Equipment) (*model.Equipment, error) {
	ctx := context.TODO()
	filter := bson.M{"_id": equipment.ID}
	update := bson.D{{"$set", equipment}}
	_, err := repo.Collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}
	return equipment, nil
}

func (repo *EquipmentRepository) Delete(id primitive.ObjectID) error {
	ctx := context.TODO()
	filter := bson.M{"_id": id}
	_, err := repo.Collection.DeleteOne(ctx, filter)
	return err
}

func (repo *EquipmentRepository) GetAvailable(ids []primitive.ObjectID) ([]model.Equipment, error) {
	ctx := context.TODO()

	filter := bson.M{"_id": bson.M{"$in": ids}}
	findOptions := options.Find()

	cursor, err := repo.Collection.Find(ctx, filter, findOptions)
	if err != nil {
		return nil, err
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {

		}
	}(cursor, ctx)

	var equipment []model.Equipment
	if err := cursor.All(ctx, &equipment); err != nil {
		return nil, err
	}

	return equipment, nil
}
