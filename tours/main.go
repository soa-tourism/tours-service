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

	dsn := "user=postgres password=super dbname=soa-tours host=localhost port=5432 sslmode=disable"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		print(err)
		return nil
	}

	err = database.AutoMigrate(
		&model.Equipment{},
		&model.Tour{},
		&model.PublicCheckpoint{},
		&model.Checkpoint{},
		&model.TouristPosition{},
	)
	if err != nil {
		log.Fatalf("Error migrating models: %v", err)
	}

	return database
}

func startServer(database *gorm.DB) {
	equipmentHandler := initEquipment(database)
	tourHandler := initTour(database)
	publicCheckpointHandler := initPublicCheckpoint(database)
	checkpointHandler := initCheckpoint(database)
	touristPositionHandler := initTouristPosition(database)

	router := mux.NewRouter().StrictSlash(true)

	initEquipmentHandler(equipmentHandler, router)
	initTourHandler(tourHandler, router)
	initPublicCheckpointHandler(publicCheckpointHandler, router)
	initCheckpointHandler(checkpointHandler, router)
	initTouristPositionHandler(touristPositionHandler, router)

	println("Server starting...")
	log.Fatal(http.ListenAndServe(":8081", router))

}

func initEquipmentHandler(equipmentHandler *handler.EquipmentHandler, router *mux.Router) {
	v1 := router.PathPrefix("/v1/tours").Subrouter()
	v1.HandleFunc("/equipment", equipmentHandler.GetAll).Methods("GET")
	v1.HandleFunc("/equipment", equipmentHandler.Create).Methods("POST")
	v1.HandleFunc("/equipment/{id}", equipmentHandler.Get).Methods("GET")
	v1.HandleFunc("/equipment/{id}", equipmentHandler.Update).Methods("PUT")
	v1.HandleFunc("/equipment/{id}", equipmentHandler.Delete).Methods("DELETE")
	v1.HandleFunc("/{id}/equipment/available", equipmentHandler.GetAvailable).Methods("GET")
}

func initTourHandler(tourHandler *handler.TourHandler, router *mux.Router) {
	v1 := router.PathPrefix("/v1/tours").Subrouter()
	v1.HandleFunc("", tourHandler.GetAll).Methods("GET")
	v1.HandleFunc("", tourHandler.Create).Methods("POST")
	v1.HandleFunc("/{id}", tourHandler.Get).Methods("GET")
	v1.HandleFunc("/{id}", tourHandler.Update).Methods("PUT")
	v1.HandleFunc("/{id}", tourHandler.Delete).Methods("DELETE")
	v1.HandleFunc("/{id}", tourHandler.Delete).Methods("DELETE")
	v1.HandleFunc("/author/{id}", tourHandler.GetByAuthor).Methods("GET")
	v1.HandleFunc("/{id}/equipment/{equipmentId}", tourHandler.AddEquipment).Methods("POST")
	v1.HandleFunc("/{id}/equipment/{equipmentId}", tourHandler.RemoveEquipment).Methods("DELETE")
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
	//v1.HandleFunc("/encounter/{checkpointId}/{encounterId}/{}", checkpointHandler.Update).Methods("PUT")
}

func initTouristPositionHandler(touristPositionHandler *handler.TouristPositionHandler, router *mux.Router) {
	v1 := router.PathPrefix("/v1/position").Subrouter()
	v1.HandleFunc("", touristPositionHandler.GetAll).Methods("GET")
	v1.HandleFunc("/{id}", touristPositionHandler.Get).Methods("GET")
	v1.HandleFunc("", touristPositionHandler.Create).Methods("POST")
	v1.HandleFunc("/{id}", touristPositionHandler.Update).Methods("PUT")
	v1.HandleFunc("/{id}", touristPositionHandler.Delete).Methods("DELETE")
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
