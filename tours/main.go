package main

import (
	"context"
	"log"
	"net"
	"strconv"
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
	checkpointService      *service.CheckpointService
	checkpointRepo         *repo.CheckpointRepository
	touristPositionService *service.TouristPositionService
	touristPositionRepo    *repo.TouristPositionRepository
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
	checkpointRepo := &repo.CheckpointRepository{Collection: database.Collection("checkpoints")}
	touristPositionRepo := &repo.TouristPositionRepository{Collection: database.Collection("tourist_positions")}
	//tourExecutionRepo := &repo.TourExecutionRepository{Collection: database.Collection("tour_executions")}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	tours.RegisterTourServer(grpcServer, Server{
		equipmentService: service.NewEquipmentService(equipmentRepo), equipmentRepo: equipmentRepo,
		tourService: service.NewTourService(tourRepo, equipmentRepo), tourRepo: tourRepo,
		checkpointService: &service.CheckpointService{
			CheckpointRepo: checkpointRepo,
			TourRepo:       tourRepo,
		}, checkpointRepo: checkpointRepo,
		touristPositionService: service.NewTouristPositionService(touristPositionRepo), touristPositionRepo: touristPositionRepo,
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
func (s Server) GetAvailableEquipment(ctx context.Context, request *tours.EquipmentIds) (*tours.EquipmentsResponse, error) {
	tourID, err := primitive.ObjectIDFromHex(request.Id)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid tour ID format")
	}
	var equipmentsIds = []primitive.ObjectID{}
	for _, id := range request.EquipmentIds {
		equipId, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return nil, status.Error(codes.InvalidArgument, "Invalid tour ID format")
		}
		equipmentsIds = append(equipmentsIds, equipId)
	}
	available, err := s.equipmentService.GetAvailableEquipment(tourID, equipmentsIds)
	if err != nil {
		return nil, status.Error(codes.NotFound, "Equipment not found")
	}

	var responses []*tours.EquipmentResponse
	for _, equipment := range available {
		response := convertEquipment(equipment)
		responses = append(responses, response)
	}

	equpmentResponse := &tours.EquipmentsResponse{
		EquipmentResponse: responses,
	}
	return equpmentResponse, nil
}
func (s Server) GetAllEquipment(ctx context.Context, request *tours.Page) (*tours.PagedEquipmentsResponse, error) {
	all, err := s.equipmentService.FindAllEquipment()
	if err != nil {
		return nil, status.Error(codes.NotFound, "Equipments not found")
	}
	var responses []*tours.EquipmentResponse
	for _, t := range all {
		var response = convertEquipment(t)
		responses = append(responses, response)
	}
	pagedResponse := &tours.PagedEquipmentsResponse{
		Results:    responses,
		TotalCount: int32(len(responses)),
	}
	return pagedResponse, nil
}
func (s Server) GetEquipment(ctx context.Context, request *tours.Id) (*tours.EquipmentResponse, error) {
	objectID, err := primitive.ObjectIDFromHex(request.Id)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid equipment ID format")
	}

	found, err := s.equipmentService.FindEquipment(objectID)
	if err != nil {
		return nil, status.Error(codes.NotFound, "Equipment not found")
	}
	var response = convertEquipment(*found)
	return response, nil
}
func (s Server) CreateEquipment(ctx context.Context, request *tours.EquipmentResponse) (*tours.EquipmentResponse, error) {
	equipment := &model.Equipment{
		Name:        request.Name,
		Description: request.Description,
	}
	created, err := s.equipmentService.CreateEquipment(equipment)
	if err != nil {
		return nil, status.Error(codes.NotFound, "Equipment not created")
	}
	var response = convertEquipment(*created)
	return response, nil
}
func (s Server) UpdateEquipment(ctx context.Context, request *tours.UpdateEquipmentId) (*tours.EquipmentResponse, error) {
	objectID, err := primitive.ObjectIDFromHex(request.Id)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid equipment ID format")
	}
	equipment := &model.Equipment{
		ID:          objectID,
		Name:        request.Equipment.Name,
		Description: request.Equipment.Description,
	}
	updated, err := s.equipmentService.UpdateEquipment(equipment)
	if err != nil {
		return nil, status.Error(codes.NotFound, "Equipment not updated")
	}
	var response = convertEquipment(*updated)
	return response, nil
}
func (s Server) DeleteEquipment(ctx context.Context, request *tours.Id) (*tours.Blank, error) {
	objectID, err := primitive.ObjectIDFromHex(request.Id)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid tour ID format")
	}
	err = s.equipmentService.DeleteEquipment(objectID)
	response := &tours.Blank{}
	return response, err
}

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
		Results:    responses,
		TotalCount: int32(len(responses)),
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
		return nil, status.Error(codes.NotFound, "Tour not updated")
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
func convertCheckpointResponse(ch tours.CheckpointResponse) *model.Checkpoint {
	objectID, err := primitive.ObjectIDFromHex(ch.Id)
	if err != nil {
		return nil
	}
	objectID2, err2 := primitive.ObjectIDFromHex(ch.TourId)
	if err2 != nil {
		return nil
	}
	response := &model.Checkpoint{
		ID:                    objectID,
		TourID:                objectID2,
		AuthorID:              ch.AuthorId,
		Longitude:             ch.Longitude,
		Latitude:              ch.Latitude,
		Name:                  ch.Name,
		Description:           ch.Description,
		Pictures:              ch.Pictures,
		RequiredTimeInSeconds: ch.RequiredTimeInSeconds,
		IsSecretPrerequisite:  ch.IsSecretPrerequisite,
		EncounterID:           ch.EncounterId,
		//checkpointSecret: e.CheckpointSecret,
	}
	return response
}
func (s Server) GetAllCheckpoints(ctx context.Context, request *tours.Page) (*tours.PagedCheckpoints, error) {
	all, err := s.checkpointService.FindAllCheckpoints()
	if err != nil {
		return nil, status.Error(codes.NotFound, "Checkpoint not found")
	}
	var responses []*tours.CheckpointResponse
	for _, t := range all {
		var response = convertCheckpoint(t)
		responses = append(responses, response)
	}
	pagedResponse := &tours.PagedCheckpoints{
		Results:    responses,
		TotalCount: int32(len(responses)),
	}
	return pagedResponse, nil
}
func (s Server) GetAllCheckpointsByTour(ctx context.Context, request *tours.PageWithId) (*tours.CheckpointsResponse, error) {
	objectID, err := primitive.ObjectIDFromHex(request.Id)
	if err != nil {
		return nil, status.Error(codes.NotFound, "bad objectID")
	}
	all, err := s.checkpointService.FindCheckpointsByTour(objectID)
	if err != nil {
		return nil, status.Error(codes.NotFound, "Checkpoint not found")
	}
	var responses []*tours.CheckpointResponse
	for _, t := range all {
		var response = convertCheckpoint(t)
		responses = append(responses, response)
	}
	pagedResponse := &tours.CheckpointsResponse{
		Checkpoints: responses,
	}
	return pagedResponse, nil
}
func (s Server) GetCheckpointById(ctx context.Context, request *tours.Id) (*tours.CheckpointResponse, error) {
	objectID, err := primitive.ObjectIDFromHex(request.Id)
	if err != nil {
		return nil, status.Error(codes.NotFound, "bad objectID")
	}
	all, err2 := s.checkpointService.FindCheckpoint(objectID)
	if err2 != nil {
		return nil, status.Error(codes.NotFound, "Checkpoint not found")
	}
	return convertCheckpoint(*all), nil
}
func (s Server) CreateCheckpoint(ctx context.Context, request *tours.CreateCheckpointRequest) (*tours.CheckpointResponse, error) {
	checkpoint := convertCheckpointResponse(*request.Checkpoint)
	all, err := s.checkpointService.CreateCheckpoint(checkpoint)
	if err != nil {
		return nil, status.Error(codes.NotFound, "Checkpoint not found")
	}
	return convertCheckpoint(*all), nil
}
func (s Server) UpdateCheckpoint(ctx context.Context, request *tours.UpdateCheckpointRequest) (*tours.CheckpointResponse, error) {
	checkpoint := convertCheckpointResponse(*request.Checkpoint)
	objectID, err := primitive.ObjectIDFromHex(request.Id)
	if err != nil {
		return nil, status.Error(codes.NotFound, "bad objectID")
	}
	checkpoint.ID = objectID
	all, err2 := s.checkpointService.UpdateCheckpoint(checkpoint)
	if err2 != nil {
		return nil, status.Error(codes.NotFound, "Equipments not found")
	}
	return convertCheckpoint(*all), nil
}
func (s Server) DeleteCheckpoint(ctx context.Context, request *tours.Id) (*tours.Blank, error) {
	objectID, err := primitive.ObjectIDFromHex(request.Id)
	if err != nil {
		return nil, status.Error(codes.NotFound, "bad objectID")
	}
	err2 := s.checkpointService.DeleteCheckpoint(objectID)
	if err2 != nil {
		return nil, status.Error(codes.NotFound, "Checkpoint not found")
	}
	return &tours.Blank{}, nil
}

// TouristPosition
func convertTouristPosition(pos model.TouristPosition) *tours.TouristPositionResponse {
	response := &tours.TouristPositionResponse{
		Id:        pos.ID.Hex(),
		CreatorId: pos.CreatorID,
		Longitude: float32(pos.Longitude),
		Latitude:  float32(pos.Latitude),
	}
	return response
}
func (s Server) GetAllTouristPostions(ctx context.Context, request *tours.Page) (*tours.PagedTouristPositionResponse, error) {
	all, err := s.touristPositionService.FindAllTouristPositions()
	if err != nil {
		return nil, status.Error(codes.NotFound, "Tourist positions not found")
	}
	var responses []*tours.TouristPositionResponse
	for _, t := range all {
		var response = convertTouristPosition(t)
		responses = append(responses, response)
	}
	pagedResponse := &tours.PagedTouristPositionResponse{
		Results:    responses,
		TotalCount: int32(len(responses)),
	}
	return pagedResponse, nil
}
func (s Server) GetTouristPostionByCreatorId(ctx context.Context, request *tours.TouristPositionByCreator) (*tours.TouristPositionResponse, error) {
	id, err := strconv.ParseInt(request.Id, 10, 32)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid ID format")
	}
	id64 := int64(id) // Convert id to int64
	found, err := s.touristPositionService.FindPositionByCreator(id64)
	if err != nil {
		return nil, status.Error(codes.NotFound, "Tourist positions not found")
	}
	var response = convertTouristPosition(found)
	return response, nil
}
func (s Server) CreateTouristPostion(ctx context.Context, request *tours.TouristPositionResponse) (*tours.TouristPositionResponse, error) {
	position := &model.TouristPosition{
		CreatorID: request.CreatorId,
		Longitude: float64(request.Longitude),
		Latitude:  float64(request.Latitude),
	}
	updated, err2 := s.touristPositionService.CreateTouristPosition(position)
	if err2 != nil {
		return nil, status.Error(codes.NotFound, "Position not created")
	}
	var response = convertTouristPosition(*updated)
	return response, nil
}
func (s Server) UpdateTouristPosition(ctx context.Context, request *tours.UpdateTouristPositionRequest) (*tours.TouristPositionResponse, error) {
	objectID, err := primitive.ObjectIDFromHex(request.Id)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid tourist position ID format")
	}
	position := &model.TouristPosition{
		ID:        objectID,
		CreatorID: request.Position.CreatorId,
		Longitude: float64(request.Position.Longitude),
		Latitude:  float64(request.Position.Latitude),
	}
	updated, err2 := s.touristPositionService.UpdateTouristPosition(position)
	if err2 != nil {
		return nil, status.Error(codes.NotFound, "Position not updated")
	}
	var response = convertTouristPosition(*updated)
	return response, nil
}
func (s Server) DeleteTouristPosition(ctx context.Context, request *tours.Id) (*tours.Blank, error) {
	objectID, err := primitive.ObjectIDFromHex(request.Id)
	if err != nil {
		return nil, status.Error(codes.NotFound, "bad position ID")
	}
	err2 := s.touristPositionService.DeleteTouristPosition(objectID)
	if err2 != nil {
		return nil, status.Error(codes.NotFound, "Tourist position not found")
	}
	return &tours.Blank{}, nil
}

// TourExecution

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
