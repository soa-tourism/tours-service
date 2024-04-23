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

type TourReviewRepository struct {
	Collection *mongo.Collection
}

func (repo *TourReviewRepository) FindAllByTourist(id int64) ([]model.TourReview, error) {
	ctx := context.TODO()
	filter := bson.M{"touristid": id}
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

	var tourReviews []model.TourReview
	if err := cursor.All(ctx, &tourReviews); err != nil {
		return nil, err
	}
	return tourReviews, nil
}

func (repo *TourReviewRepository) FindAllByAuthor(id int64) ([]model.TourReview, error) {
	ctx := context.TODO()
	filter := bson.M{"authorid": id}
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

	var tourReviews []model.TourReview
	if err := cursor.All(ctx, &tourReviews); err != nil {
		return nil, err
	}
	return tourReviews, nil
}

func (repo *TourReviewRepository) FindAllByTour(id primitive.ObjectID) ([]model.TourReview, error) {
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

	var tourReviews []model.TourReview
	if err := cursor.All(ctx, &tourReviews); err != nil {
		return nil, err
	}
	return tourReviews, nil
}

func (repo *TourReviewRepository) FindById(id primitive.ObjectID) (*model.TourReview, error) {
	ctx := context.TODO()
	filter := bson.M{"_id": id}
	var tourReview model.TourReview
	err := repo.Collection.FindOne(ctx, filter).Decode(&tourReview)
	if err != nil {
		return nil, err
	}
	return &tourReview, nil
}

func (repo *TourReviewRepository) Create(tourReview *model.TourReview) (model.TourReview, error) {
	ctx := context.TODO()
	result, err := repo.Collection.InsertOne(ctx, tourReview)
	if err != nil {
		return model.TourReview{}, err
	}

	insertedID, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return model.TourReview{}, errors.New("failed to get inserted ID")
	}
	tourReview.ID = insertedID

	return *tourReview, nil
}

func (repo *TourReviewRepository) GetAverageRating(id primitive.ObjectID) (float64, error) {
	ctx := context.TODO()
	pipeline := bson.A{
		bson.D{{"$match", bson.D{{"tourid", id}}}},
		bson.D{{"$group", bson.D{{"_id", nil}, {"average_rating", bson.D{{"$avg", "$rating"}}}}}},
	}

	cursor, err := repo.Collection.Aggregate(ctx, pipeline)
	if err != nil {
		return 0, err
	}
	defer func(cursor mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {
			fmt.Printf("Error closing cursor: %v\n", err)
		}
	}(*cursor, ctx)

	var result struct {
		AverageRating float64 `bson:"average_rating"`
	}
	if cursor.Next(ctx) {
		if err := cursor.Decode(&result); err != nil {
			return 0, err
		}
		return result.AverageRating, nil
	}

	return 0, nil
}
