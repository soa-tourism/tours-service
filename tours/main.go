package main

import (
	"log"
	"net/http"
	"tours/handler"
	"tours/model"
	"tours/repo"
	"tours/service"

	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func initDB() *gorm.DB {

	dsn := "user=postgres password=super dbname=soa-tours host=tours_db port=5432 sslmode=disable"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		print(err)
		return nil
	}

	err = database.AutoMigrate(
		&model.Equipment{},
		&model.PublicCheckpoint{},
		&model.Tour{},
		&model.Checkpoint{},
		&model.TourReview{},
		&model.TouristPosition{},
		&model.TourExecution{},
		&model.CheckpointCompletition{},
	)
	if err != nil {
		log.Fatalf("Error migrating models: %v", err)
	}

	return database
}

func startServer(database *gorm.DB) {
	equipmentHandler := initEquipment(database)
	tourHandler := initTour(database)
	publishedTourHandler := initPublishedTour(database)
	tourReviewHandler := initTourReview(database)
	imageHandler := handler.NewImageHandler()
	publicCheckpointHandler := initPublicCheckpoint(database)
	checkpointHandler := initCheckpoint(database)
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
	log.Fatal(http.ListenAndServe(":81", router))
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

func initEquipment(database *gorm.DB) *handler.EquipmentHandler {
	equipmentRepo := &repo.EquipmentRepository{DB: database}
	equipmentService := &service.EquipmentService{EquipmentRepo: equipmentRepo}
	equipmentHandler := &handler.EquipmentHandler{EquipmentService: equipmentService}

	return equipmentHandler
}

func initTour(database *gorm.DB) *handler.TourHandler {
	tourRepo := &repo.TourRepository{DB: database}
	tourService := &service.TourService{TourRepo: tourRepo}
	tourHandler := &handler.TourHandler{TourService: tourService}

	return tourHandler
}

func initPublishedTour(database *gorm.DB) *handler.PublishedTourHandler {
	publishedTourRepo := &repo.TourRepository{DB: database}
	publishedTourService := &service.PublishedTourService{TourRepo: publishedTourRepo}
	publishedTourHandler := &handler.PublishedTourHandler{PublishedTourService: publishedTourService}

	return publishedTourHandler
}

func initTourReview(database *gorm.DB) *handler.TourReviewHandler {
	tourReviewRepo := &repo.TourReviewRepository{DB: database}
	tourReviewService := &service.TourReviewService{TourReviewRepo: tourReviewRepo}
	tourReviewHandler := &handler.TourReviewHandler{TourReviewService: tourReviewService}

	return tourReviewHandler
}

func initPublicCheckpoint(database *gorm.DB) *handler.PublicCheckpointHandler {
	publicCheckpointRepo := &repo.PublicCheckpointRepository{DB: database}
	publicCheckpointService := &service.PublicCheckpointService{CheckpointRepo: publicCheckpointRepo}
	publicCheckpointHandler := &handler.PublicCheckpointHandler{CheckpointService: publicCheckpointService}
	return publicCheckpointHandler
}

func initCheckpoint(database *gorm.DB) *handler.CheckpointHandler {
	checkpointRepo := &repo.CheckpointRepository{DB: database}
	checkpointService := &service.CheckpointService{CheckpointRepo: checkpointRepo}
	checkpointHandler := &handler.CheckpointHandler{CheckpointService: checkpointService}
	return checkpointHandler
}

func initTouristPosition(database *gorm.DB) *handler.TouristPositionHandler {
	touristPositionRepo := &repo.TouristPositionRepository{DB: database}
	touristPositionService := &service.TouristPositionService{TouristPositionRepo: touristPositionRepo}
	touristPositionHandler := &handler.TouristPositionHandler{TouristPositionService: touristPositionService}
	return touristPositionHandler
}

func initTourExecution(database *gorm.DB) *handler.TourExecutionHandler {
	tourExecutionRepo := &repo.TourExecutionRepository{DB: database}
	tourExecutionService := &service.TourExecutionService{TourExecutionRepo: tourExecutionRepo}
	tourExecutionHandler := &handler.TourExecutionHandler{TourExecutionService: tourExecutionService}

	return tourExecutionHandler
}
