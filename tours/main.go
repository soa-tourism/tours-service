package main

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"tours/handler"
	"tours/repo"
	"tours/service"

	"github.com/gorilla/mux"
)

func initDB() *mongo.Database {
	//mongoURI := os.Getenv("MONGO_DB_URI")
	mongoURI := "mongodb://root:pass@mongo:27017/"
	clientOptions := options.Client().ApplyURI(mongoURI)

	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatalf("Error connecting to MongoDB: %v", err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatalf("Error pinging MongoDB: %v", err)
	}

	databaseName := "tours_db"
	db := client.Database(databaseName)

	collectionNames := []string{"tours", "equipment", "tour_reviews", "public_checkpoints", "tourist_positions", "tour_executions"}

	for _, collectionName := range collectionNames {
		if err := createCollection(db, collectionName); err != nil {
			log.Fatalf("Error creating collection %s: %v", collectionName, err)
		}
	}
	return db
}

func createCollection(db *mongo.Database, collectionName string) error {
	ctx := context.Background()
	collectionExists, err := db.Collection(collectionName).EstimatedDocumentCount(ctx)
	if err != nil {
		return err
	}

	if collectionExists == 0 {
		err := db.CreateCollection(ctx, collectionName)
		if err != nil {
			return err
		}
	}

	return nil
}

func startServer(database *mongo.Database) {
	equipmentHandler, equipmentRepo := initEquipment(database)
	tourHandler, tourRepo := initTour(database, equipmentRepo)
	publishedTourHandler := initPublishedTour(database)
	tourReviewHandler := initTourReview(database)
	imageHandler := handler.NewImageHandler()
	publicCheckpointHandler := initPublicCheckpoint(database)
	checkpointHandler := initCheckpoint(database, tourRepo)
	touristPositionHandler := initTouristPosition(database)
	tourExecutionHandler := initTourExecution(database)

	router := mux.NewRouter().StrictSlash(true)

	initEquipmentHandler(equipmentHandler, router)
	initPublicCheckpointHandler(publicCheckpointHandler, router)
	initCheckpointHandler(checkpointHandler, router)
	initPublishedTourHandler(publishedTourHandler, router)
	initTourReviewHandler(tourReviewHandler, router)
	initTourHandler(tourHandler, router)
	initImageHandler(imageHandler, router)
	initTouristPositionHandler(touristPositionHandler, router)
	initTourExecutionHandler(tourExecutionHandler, router)

	println("Server starting...")
	log.Fatal(http.ListenAndServe(":8083", router))
}

func initEquipmentHandler(equipmentHandler *handler.EquipmentHandler, router *mux.Router) {
	v1 := router.PathPrefix("/v1/tours").Subrouter()
	v1.HandleFunc("/{id}/equipment/available", equipmentHandler.GetAvailable).Methods("GET")
	v1.HandleFunc("/equipment/{id}", equipmentHandler.Get).Methods("GET")
	v1.HandleFunc("/equipment/{id}", equipmentHandler.Update).Methods("PUT")
	v1.HandleFunc("/equipment/{id}", equipmentHandler.Delete).Methods("DELETE")
	v1.HandleFunc("/equipment", equipmentHandler.GetAll).Methods("GET")
	v1.HandleFunc("/equipment", equipmentHandler.Create).Methods("POST")
}

func initTourHandler(tourHandler *handler.TourHandler, router *mux.Router) {
	v1 := router.PathPrefix("/v1/tours").Subrouter()
	v1.HandleFunc("/{id}/equipment/{equipmentId}", tourHandler.AddEquipment).Methods("POST")
	v1.HandleFunc("/{id}/equipment/{equipmentId}", tourHandler.RemoveEquipment).Methods("DELETE")
	v1.HandleFunc("/author/{id}", tourHandler.GetByAuthor).Methods("GET")
	v1.HandleFunc("/{id}", tourHandler.Get).Methods("GET")
	v1.HandleFunc("/{id}", tourHandler.Update).Methods("PUT")
	v1.HandleFunc("/{id}", tourHandler.Delete).Methods("DELETE")
	v1.HandleFunc("", tourHandler.GetAll).Methods("GET")
	v1.HandleFunc("", tourHandler.Create).Methods("POST")
}

func initPublishedTourHandler(publishedTourHandler *handler.PublishedTourHandler, router *mux.Router) {
	v1 := router.PathPrefix("/v1/tours/published").Subrouter()
	v1.HandleFunc("/{id}", publishedTourHandler.Get).Methods("GET")
	v1.HandleFunc("", publishedTourHandler.GetAllPublished).Methods("GET")
}

func initTourReviewHandler(tourReviewHandler *handler.TourReviewHandler, router *mux.Router) {
	v1 := router.PathPrefix("/v1/tours/reviews").Subrouter()
	v1.HandleFunc("/tourist/{id}", tourReviewHandler.GetAllByTourist).Methods("GET")
	v1.HandleFunc("/author/{id}", tourReviewHandler.GetAllByAuthor).Methods("GET")
	v1.HandleFunc("/tour/{id}", tourReviewHandler.GetAllByTour).Methods("GET")
	v1.HandleFunc("/average/{id}", tourReviewHandler.GetAverageRating).Methods("GET")
	v1.HandleFunc("/{id}", tourReviewHandler.Get).Methods("GET")
	v1.HandleFunc("", tourReviewHandler.Create).Methods("POST")
}

func initImageHandler(imageHandler *handler.ImageHandler, router *mux.Router) {
	v1 := router.PathPrefix("/v1/images").Subrouter()
	v1.HandleFunc("/{image}", imageHandler.ServeImage).Methods("GET")
}

func initPublicCheckpointHandler(publicCheckpointHandler *handler.PublicCheckpointHandler, router *mux.Router) {
	v1 := router.PathPrefix("/v1/publicCheckpoint").Subrouter()
	v1.HandleFunc("", publicCheckpointHandler.GetAll).Methods("GET")
	v1.HandleFunc("/details/{id}", publicCheckpointHandler.Get).Methods("GET")
	v1.HandleFunc("", publicCheckpointHandler.Create).Methods("POST")
	v1.HandleFunc("/{id}", publicCheckpointHandler.Update).Methods("PUT")
	v1.HandleFunc("/{id}", publicCheckpointHandler.Delete).Methods("DELETE")
}

func initCheckpointHandler(checkpointHandler *handler.CheckpointHandler, router *mux.Router) {
	v1 := router.PathPrefix("/v1/checkpoint").Subrouter()
	v1.HandleFunc("/", checkpointHandler.GetAll).Methods("GET")
	v1.HandleFunc("/details/{id}", checkpointHandler.Get).Methods("GET")
	v1.HandleFunc("", checkpointHandler.Create).Methods("POST")
	v1.HandleFunc("/{id}", checkpointHandler.Update).Methods("PUT")
	v1.HandleFunc("/{id}", checkpointHandler.Delete).Methods("DELETE")
	v1.HandleFunc("/{id}", checkpointHandler.GetByTour).Methods("GET")
	v1.HandleFunc("/createSecret/{id}", checkpointHandler.UpdateCheckpointSecret).Methods("PUT")
	v1.HandleFunc("/encounter/{checkpointId}/{encounterId}/{isSecretPrerequisite}", checkpointHandler.UpdateCheckpointEncounter).Methods("PUT")
}

func initTouristPositionHandler(touristPositionHandler *handler.TouristPositionHandler, router *mux.Router) {
	v1 := router.PathPrefix("/v1/position").Subrouter()
	v1.HandleFunc("", touristPositionHandler.GetAll).Methods("GET")
	v1.HandleFunc("/{id}", touristPositionHandler.Get).Methods("GET")
	v1.HandleFunc("", touristPositionHandler.Create).Methods("POST")
	v1.HandleFunc("/{id}", touristPositionHandler.Update).Methods("PUT")
	v1.HandleFunc("/{id}", touristPositionHandler.Delete).Methods("DELETE")
}

func initTourExecutionHandler(tourExecutionHandler *handler.TourExecutionHandler, router *mux.Router) {
	v1 := router.PathPrefix("/v1/execution").Subrouter()
	v1.HandleFunc("", tourExecutionHandler.GetAll).Methods("GET")
	v1.HandleFunc("/{id}", tourExecutionHandler.Get).Methods("GET")
	v1.HandleFunc("", tourExecutionHandler.Create).Methods("POST")
	v1.HandleFunc("/{id}", tourExecutionHandler.Update).Methods("PUT")
	v1.HandleFunc("/{id}", tourExecutionHandler.Delete).Methods("DELETE")
	v1.HandleFunc("/all/{tourId}/{touristId}", tourExecutionHandler.GetByTouristAndTour).Methods("GET")
	v1.HandleFunc("/{tourId}/{touristId}", tourExecutionHandler.GetActiveByTouristAndTour).Methods("GET")
}

func main() {
	database := initDB()
	if database == nil {
		print("FAILED TO CONNECT TO DB")
		return
	}

	startServer(database)
}

func initEquipment(db *mongo.Database) (*handler.EquipmentHandler, *repo.EquipmentRepository) {
	equipmentRepo := &repo.EquipmentRepository{Collection: db.Collection("equipment")}
	equipmentService := &service.EquipmentService{EquipmentRepo: equipmentRepo}
	equipmentHandler := &handler.EquipmentHandler{EquipmentService: equipmentService}

	return equipmentHandler, equipmentRepo
}

func initTour(db *mongo.Database, equipmentRepo *repo.EquipmentRepository) (*handler.TourHandler, *repo.TourRepository) {
	tourRepo := &repo.TourRepository{Collection: db.Collection("tours")}
	tourService := &service.TourService{TourRepo: tourRepo, EquipmentRepo: equipmentRepo}
	tourHandler := &handler.TourHandler{TourService: tourService}

	return tourHandler, tourRepo
}

func initPublishedTour(db *mongo.Database) *handler.PublishedTourHandler {
	publishedTourRepo := &repo.TourRepository{Collection: db.Collection("published_tours")}
	publishedTourService := &service.PublishedTourService{TourRepo: publishedTourRepo}
	publishedTourHandler := &handler.PublishedTourHandler{PublishedTourService: publishedTourService}

	return publishedTourHandler
}

func initTourReview(db *mongo.Database) *handler.TourReviewHandler {
	tourReviewRepo := &repo.TourReviewRepository{Collection: db.Collection("tour_reviews")}
	tourReviewService := &service.TourReviewService{TourReviewRepo: tourReviewRepo}
	tourReviewHandler := &handler.TourReviewHandler{TourReviewService: tourReviewService}

	return tourReviewHandler
}

func initPublicCheckpoint(db *mongo.Database) *handler.PublicCheckpointHandler {
	publicCheckpointRepo := &repo.PublicCheckpointRepository{Collection: db.Collection("public_checkpoints")}
	publicCheckpointService := &service.PublicCheckpointService{CheckpointRepo: publicCheckpointRepo}
	publicCheckpointHandler := &handler.PublicCheckpointHandler{CheckpointService: publicCheckpointService}
	return publicCheckpointHandler
}

func initCheckpoint(db *mongo.Database, tourRepo *repo.TourRepository) *handler.CheckpointHandler {
	checkpointRepo := &repo.CheckpointRepository{Collection: db.Collection("checkpoints")}
	checkpointService := &service.CheckpointService{CheckpointRepo: checkpointRepo, TourRepo: tourRepo}
	checkpointHandler := &handler.CheckpointHandler{CheckpointService: checkpointService}
	return checkpointHandler
}

func initTouristPosition(db *mongo.Database) *handler.TouristPositionHandler {
	touristPositionRepo := &repo.TouristPositionRepository{Collection: db.Collection("tourist_positions")}
	touristPositionService := &service.TouristPositionService{TouristPositionRepo: touristPositionRepo}
	touristPositionHandler := &handler.TouristPositionHandler{TouristPositionService: touristPositionService}
	return touristPositionHandler
}

func initTourExecution(db *mongo.Database) *handler.TourExecutionHandler {
	tourExecutionRepo := &repo.TourExecutionRepository{Collection: db.Collection("tour_executions")}
	tourExecutionService := &service.TourExecutionService{TourExecutionRepo: tourExecutionRepo}
	tourExecutionHandler := &handler.TourExecutionHandler{TourExecutionService: tourExecutionService}

	return tourExecutionHandler
}
