package main

import (
	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
	"tours/handler"
	"tours/model"
	"tours/repo"
	"tours/service"
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
	)
	if err != nil {
		log.Fatalf("Error migrating models: %v", err)
	}

	return database
}

func startServer(database *gorm.DB) {
	equipmentHandler := initEquipment(database)
	tourHandler := initTour(database)

	router := mux.NewRouter().StrictSlash(true)

	initEquipmentHandler(equipmentHandler, router)
	initTourHandler(tourHandler, router)

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
