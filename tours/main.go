package main

import (
	"context"
	"log"
	"net"
	"tours/model"
	"tours/proto/tours"
	"tours/repo"
	"tours/service"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
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
	tourService      *service.TourService
	tourRepo         *repo.TourRepository
	//publishedTourService *service.TourService
	//publishedTourRepo *repo.TourRepository
	//tourReviewService *service.TourReviewService
	//tourReviewRepo *repo.TourReviewRepository
	//publicCheckpointService *service.PublicCheckpointService
	//publicCheckpointRepo *repo.PublicCheckpointRepository
	//checkpointService *service.CheckpointService
	//checkpointRepo *repo.CheckpointRepository
	// touristPositionService *service.TouristPositionService
	//touristPositionRepo *repo.TouristPositionRepository
	//tourExecutionService *service.TourExecutionService
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
	tourRepo := &repo.TourRepository{Collection: database.Collection("tours")}
	//publishedTourRepo := &repo.TourRepository{Collection: database.Collection("published_tours")}
	//tourReviewRepo := &repo.TourReviewRepository{Collection: database.Collection("tour_reviews")}
	//publicCheckpointRepo := &repo.PublicCheckpointRepository{Collection: database.Collection("public_checkpoints")}
	//checkpointRepo := &repo.CheckpointRepository{Collection: database.Collection("checkpoints")}
	//touristPositionRepo := &repo.TouristPositionRepository{Collection: database.Collection("tourist_positions")}
	//tourExecutionRepo := &repo.TourExecutionRepository{Collection: database.Collection("tour_executions")}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	tours.RegisterTourServer(grpcServer, Server{
		equipmentService: service.NewEquipmentService(equipmentRepo), equipmentRepo: equipmentRepo,
		tourService: service.NewTourService(tourRepo, equipmentRepo), tourRepo: tourRepo,
	})
	reflection.Register(grpcServer)
	grpcServer.Serve(lis)
}

// Equipment
func convertEquipment(e model.Equipment) *tours.EquipmentResponse {
	response := &tours.EquipmentResponse{
		Id:          e.ID.Hex(),
		Name:        e.Name,
		Description: e.Description,
	}
	return response
}

// TODO
// 	rpc GetAvailableEquipment (EquipmentIds) returns (EquipmentsResponse) {}
//  rpc GetAllEquipment (Page) returns (PagedEquipmentsResponse) {}
// 	rpc GetEquipment (Id) returns (EquipmentResponse) {}
// 	rpc CreateEquipment (EquipmentResponse) returns (EquipmentResponse) {}
// 	rpc UpdateEquipment (UpdateEquipmentId) returns (EquipmentResponse) {}
// 	rpc DeleteEquipment (Id) returns (EquipmentResponse) {}

// Tour
func convertTour(tour model.Tour) *tours.TourResponse {
	// equipments
	var equipments []*tours.EquipmentResponse
	for _, e := range tour.Equipment {
		var equipment = convertEquipment(e)
		equipments = append(equipments, equipment)
	}
	//checkpoints
	var checkpoints []*tours.CheckpointResponse
	for _, ch := range tour.Checkpoints {
		var checkpoint = convertCheckpoint(ch)
		checkpoints = append(checkpoints, checkpoint)
	}

	response := &tours.TourResponse{
		Id:          tour.ID.Hex(),
		AuthorId:    tour.AuthorID,
		Name:        tour.Name,
		Description: tour.Description,
		Difficulty:  tour.Difficulty.String(),
		Status:      tour.Status.String(),
		Price:       tour.Price,
		Tags:        tour.Tags,
		Equipment:   equipments,
		Checkpoints: checkpoints,
	}
	return response
}
func (s Server) GetAllTour(ctx context.Context, request *tours.Page) (*tours.PagedToursResponse, error) {
	all, err := s.tourService.FindAllTours()
	if err != nil {
		return nil, status.Error(codes.NotFound, "Tours not found")
	}
	var responses []*tours.TourResponse
	for _, t := range all {
		var response = convertTour(t)
		responses = append(responses, response)
	}
	pagedResponse := &tours.PagedToursResponse{
		Results:     responses,
		TotalCounts: int32(len(responses)),
	}
	return pagedResponse, nil
}
func (s Server) GetTour(ctx context.Context, request *tours.Id) (*tours.TourResponse, error) {
	objectID, err := primitive.ObjectIDFromHex(request.Id)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid tour ID format")
	}

	found, err := s.tourService.FindTour(objectID)
	if err != nil {
		return nil, status.Error(codes.NotFound, "Tour not found")
	}
	var response = convertTour(*found)
	return response, nil
}
func (s Server) CreateTour(ctx context.Context, request *tours.TourResponse) (*tours.TourResponse, error) {
	tour := &model.Tour{
		AuthorID:    request.AuthorId,
		Name:        request.Name,
		Description: request.Description,
		Difficulty:  model.ParseDifficulty(request.Difficulty),
		Status:      model.ParseStatus(request.Status),
		Price:       request.Price,
		Tags:        request.Tags,
		Equipment:   []model.Equipment{},
		Checkpoints: []model.Checkpoint{},
	}
	created, err := s.tourService.CreateTour(tour)
	if err != nil {
		return nil, status.Error(codes.NotFound, "Tour not created")
	}
	var response = convertTour(*created)
	return response, nil
}
func (s Server) UpdateTour(ctx context.Context, request *tours.UpdateTourId) (*tours.TourResponse, error) {
	objectID, err := primitive.ObjectIDFromHex(request.Id)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid tour ID format")
	}
	tour := &model.Tour{
		ID:          objectID,
		AuthorID:    request.Tour.AuthorId,
		Name:        request.Tour.Name,
		Description: request.Tour.Description,
		Difficulty:  model.ParseDifficulty(request.Tour.Difficulty),
		Status:      model.ParseStatus(request.Tour.Status),
		Price:       request.Tour.Price,
		Tags:        request.Tour.Tags,
		Equipment:   []model.Equipment{},
		Checkpoints: []model.Checkpoint{},
	}
	updated, err := s.tourService.UpdateTour(tour)
	if err != nil {
		return nil, status.Error(codes.NotFound, "Tour not created")
	}
	var response = convertTour(*updated)
	return response, nil
}
func (s Server) DeleteTour(ctx context.Context, request *tours.Id) (*tours.Blank, error) {
	objectID, err := primitive.ObjectIDFromHex(request.Id)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid tour ID format")
	}
	s.tourService.DeleteTour(objectID)
	response := &tours.Blank{}
	return response, nil
}
func (s Server) GetTourByAuthorId(ctx context.Context, request *tours.TourByAuthorId) (*tours.ToursResponse, error) {
	found, err := s.tourService.FindToursByAuthor(request.Id)
	if err != nil {
		return nil, status.Error(codes.NotFound, "Tour not found")
	}
	var responses []*tours.TourResponse
	for _, t := range found {
		response := convertTour(t)
		responses = append(responses, response)
	}

	toursResponse := &tours.ToursResponse{
		Tours: responses,
	}
	return toursResponse, nil
}
func (s Server) AddTourEquipment(ctx context.Context, request *tours.TourEquipment) (*tours.Blank, error) {
	tourID, err := primitive.ObjectIDFromHex(request.Id)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid tour ID format")
	}
	equipmentID, err := primitive.ObjectIDFromHex(request.EquipmentId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid tour ID format")
	}

	s.tourService.AddEquipment(tourID, equipmentID)
	response := &tours.Blank{}
	return response, nil
}
func (s Server) DeleteTourEquipment(ctx context.Context, request *tours.TourEquipment) (*tours.Blank, error) {
	tourID, err := primitive.ObjectIDFromHex(request.Id)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid tour ID format")
	}
	equipmentID, err := primitive.ObjectIDFromHex(request.EquipmentId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid tour ID format")
	}

	s.tourService.RemoveEquipment(tourID, equipmentID)
	response := &tours.Blank{}
	return response, nil
}

// PublishedTour
// TourReview
// PublicCheckpoint
// Checkpoint
func convertCheckpoint(ch model.Checkpoint) *tours.CheckpointResponse {
	response := &tours.CheckpointResponse{
		Id:                    ch.ID.Hex(),
		TourId:                ch.TourID.Hex(),
		AuthorId:              ch.AuthorID,
		Longitude:             ch.Longitude,
		Latitude:              ch.Latitude,
		Name:                  ch.Name,
		Description:           ch.Description,
		Pictures:              ch.Pictures,
		RequiredTimeInSeconds: ch.RequiredTimeInSeconds,
		IsSecretPrerequisite:  ch.IsSecretPrerequisite,
		EncounterId:           ch.EncounterID,
		//checkpointSecret: e.CheckpointSecret,
	}
	return response
}

// TouristPosition
// TourExecution

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
