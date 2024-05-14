package main

import (
	"context"
	"log"
	"net"
	"tours/proto/tours"
	"tours/repo"
	"tours/service"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
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

type Server struct {
	tours.UnimplementedTourServer
	equipmentService *service.EquipmentService
	equipmentRepo    *repo.EquipmentRepository
	//tourRepo *repo.TourRepository
	//publishedTourRepo *repo.TourRepository
	//tourReviewRepo *repo.TourReviewRepository
	//publicCheckpointRepo *repo.PublicCheckpointRepository
	//checkpointRepo *repo.CheckpointRepository
	//touristPositionRepo *repo.TouristPositionRepository
	//tourExecutionRepo *repo.TourExecutionRepository
}

func main() {
	database := initDB()
	if database == nil {
		print("FAILED TO CONNECT TO DB")
		return
	}

	lis, err := net.Listen("tcp", "0.0.0.0:8083")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	equipmentRepo := &repo.EquipmentRepository{Collection: database.Collection("equipment")}
	//tourRepo := &repo.TourRepository{Collection: database.Collection("tours")}
	//publishedTourRepo := &repo.TourRepository{Collection: database.Collection("published_tours")}
	//tourReviewRepo := &repo.TourReviewRepository{Collection: database.Collection("tour_reviews")}
	//publicCheckpointRepo := &repo.PublicCheckpointRepository{Collection: database.Collection("public_checkpoints")}
	//checkpointRepo := &repo.CheckpointRepository{Collection: database.Collection("checkpoints")}
	//touristPositionRepo := &repo.TouristPositionRepository{Collection: database.Collection("tourist_positions")}
	//tourExecutionRepo := &repo.TourExecutionRepository{Collection: database.Collection("tour_executions")}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	tours.RegisterTourServer(grpcServer, Server{
		equipmentService: service.NewEquipmentService(equipmentRepo), equipmentRepo: equipmentRepo})
	reflection.Register(grpcServer)
	grpcServer.Serve(lis)
}

// func startServer(database *mongo.Database) {
// 	equipmentHandler, equipmentRepo := initEquipment(database)
// 	tourHandler, tourRepo := initTour(database, equipmentRepo)
// 	publishedTourHandler := initPublishedTour(database)
// 	tourReviewHandler := initTourReview(database)
// 	imageHandler := handler.NewImageHandler()
// 	publicCheckpointHandler := initPublicCheckpoint(database)
// 	checkpointHandler := initCheckpoint(database, tourRepo)
// 	touristPositionHandler := initTouristPosition(database)
// 	tourExecutionHandler := initTourExecution(database)

// 	router := mux.NewRouter().StrictSlash(true)

// 	initEquipmentHandler(equipmentHandler, router)
// 	initPublicCheckpointHandler(publicCheckpointHandler, router)
// 	initCheckpointHandler(checkpointHandler, router)
// 	initPublishedTourHandler(publishedTourHandler, router)
// 	initTourReviewHandler(tourReviewHandler, router)
// 	initTourHandler(tourHandler, router)
// 	initImageHandler(imageHandler, router)
// 	initTouristPositionHandler(touristPositionHandler, router)
// 	initTourExecutionHandler(tourExecutionHandler, router)

// 	println("Server starting...")
// 	log.Fatal(http.ListenAndServe(":8083", router))
// }

// func initEquipmentHandler(equipmentHandler *handler.EquipmentHandler, router *mux.Router) {
// 	v1 := router.PathPrefix("/v1/tours").Subrouter()
// 	v1.HandleFunc("/{id}/equipment/available", equipmentHandler.GetAvailable).Methods("GET")
// 	v1.HandleFunc("/equipment/{id}", equipmentHandler.Get).Methods("GET")
// 	v1.HandleFunc("/equipment/{id}", equipmentHandler.Update).Methods("PUT")
// 	v1.HandleFunc("/equipment/{id}", equipmentHandler.Delete).Methods("DELETE")
// 	v1.HandleFunc("/equipment", equipmentHandler.GetAll).Methods("GET")
// 	v1.HandleFunc("/equipment", equipmentHandler.Create).Methods("POST")
// }

// func initTourHandler(tourHandler *handler.TourHandler, router *mux.Router) {
// 	v1 := router.PathPrefix("/v1/tours").Subrouter()
// 	v1.HandleFunc("/{id}/equipment/{equipmentId}", tourHandler.AddEquipment).Methods("POST")
// 	v1.HandleFunc("/{id}/equipment/{equipmentId}", tourHandler.RemoveEquipment).Methods("DELETE")
// 	v1.HandleFunc("/author/{id}", tourHandler.GetByAuthor).Methods("GET")
// 	v1.HandleFunc("/{id}", tourHandler.Get).Methods("GET")
// 	v1.HandleFunc("/{id}", tourHandler.Update).Methods("PUT")
// 	v1.HandleFunc("/{id}", tourHandler.Delete).Methods("DELETE")
// 	v1.HandleFunc("", tourHandler.GetAll).Methods("GET")
// 	v1.HandleFunc("", tourHandler.Create).Methods("POST")
// }

// func initPublishedTourHandler(publishedTourHandler *handler.PublishedTourHandler, router *mux.Router) {
// 	v1 := router.PathPrefix("/v1/tours/published").Subrouter()
// 	v1.HandleFunc("/{id}", publishedTourHandler.Get).Methods("GET")
// 	v1.HandleFunc("", publishedTourHandler.GetAllPublished).Methods("GET")
// }

// func initTourReviewHandler(tourReviewHandler *handler.TourReviewHandler, router *mux.Router) {
// 	v1 := router.PathPrefix("/v1/tours/reviews").Subrouter()
// 	v1.HandleFunc("/tourist/{id}", tourReviewHandler.GetAllByTourist).Methods("GET")
// 	v1.HandleFunc("/author/{id}", tourReviewHandler.GetAllByAuthor).Methods("GET")
// 	v1.HandleFunc("/tour/{id}", tourReviewHandler.GetAllByTour).Methods("GET")
// 	v1.HandleFunc("/average/{id}", tourReviewHandler.GetAverageRating).Methods("GET")
// 	v1.HandleFunc("/{id}", tourReviewHandler.Get).Methods("GET")
// 	v1.HandleFunc("", tourReviewHandler.Create).Methods("POST")
// }

// func initImageHandler(imageHandler *handler.ImageHandler, router *mux.Router) {
// 	v1 := router.PathPrefix("/v1/images").Subrouter()
// 	v1.HandleFunc("/{image}", imageHandler.ServeImage).Methods("GET")
// }

// func initPublicCheckpointHandler(publicCheckpointHandler *handler.PublicCheckpointHandler, router *mux.Router) {
// 	v1 := router.PathPrefix("/v1/publicCheckpoint").Subrouter()
// 	v1.HandleFunc("", publicCheckpointHandler.GetAll).Methods("GET")
// 	v1.HandleFunc("/details/{id}", publicCheckpointHandler.Get).Methods("GET")
// 	v1.HandleFunc("", publicCheckpointHandler.Create).Methods("POST")
// 	v1.HandleFunc("/{id}", publicCheckpointHandler.Update).Methods("PUT")
// 	v1.HandleFunc("/{id}", publicCheckpointHandler.Delete).Methods("DELETE")
// }

// func initCheckpointHandler(checkpointHandler *handler.CheckpointHandler, router *mux.Router) {
// 	v1 := router.PathPrefix("/v1/checkpoint").Subrouter()
// 	v1.HandleFunc("/", checkpointHandler.GetAll).Methods("GET")
// 	v1.HandleFunc("/details/{id}", checkpointHandler.Get).Methods("GET")
// 	v1.HandleFunc("", checkpointHandler.Create).Methods("POST")
// 	v1.HandleFunc("/{id}", checkpointHandler.Update).Methods("PUT")
// 	v1.HandleFunc("/{id}", checkpointHandler.Delete).Methods("DELETE")
// 	v1.HandleFunc("/{id}", checkpointHandler.GetByTour).Methods("GET")
// 	v1.HandleFunc("/createSecret/{id}", checkpointHandler.UpdateCheckpointSecret).Methods("PUT")
// 	v1.HandleFunc("/encounter/{checkpointId}/{encounterId}/{isSecretPrerequisite}", checkpointHandler.UpdateCheckpointEncounter).Methods("PUT")
// }

// func initTouristPositionHandler(touristPositionHandler *handler.TouristPositionHandler, router *mux.Router) {
// 	v1 := router.PathPrefix("/v1/position").Subrouter()
// 	v1.HandleFunc("", touristPositionHandler.GetAll).Methods("GET")
// 	v1.HandleFunc("/{id}", touristPositionHandler.Get).Methods("GET")
// 	v1.HandleFunc("", touristPositionHandler.Create).Methods("POST")
// 	v1.HandleFunc("/{id}", touristPositionHandler.Update).Methods("PUT")
// 	v1.HandleFunc("/{id}", touristPositionHandler.Delete).Methods("DELETE")
// }

// func initTourExecutionHandler(tourExecutionHandler *handler.TourExecutionHandler, router *mux.Router) {
// 	v1 := router.PathPrefix("/v1/execution").Subrouter()
// 	v1.HandleFunc("", tourExecutionHandler.GetAll).Methods("GET")
// 	v1.HandleFunc("/{id}", tourExecutionHandler.Get).Methods("GET")
// 	v1.HandleFunc("", tourExecutionHandler.Create).Methods("POST")
// 	v1.HandleFunc("/{id}", tourExecutionHandler.Update).Methods("PUT")
// 	v1.HandleFunc("/{id}", tourExecutionHandler.Delete).Methods("DELETE")
// 	v1.HandleFunc("/all/{tourId}/{touristId}", tourExecutionHandler.GetByTouristAndTour).Methods("GET")
// 	v1.HandleFunc("/{tourId}/{touristId}", tourExecutionHandler.GetActiveByTouristAndTour).Methods("GET")
// }
