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
	router.HandleFunc("/tours/equipment", equipmentHandler.GetAll).Methods("GET")
	router.HandleFunc("/tours/equipment", equipmentHandler.Create).Methods("POST")
	router.HandleFunc("/tours/equipment/{id}", equipmentHandler.Get).Methods("GET")
	router.HandleFunc("/tours/equipment/{id}", equipmentHandler.Update).Methods("PUT")
	router.HandleFunc("/tours/equipment/{id}", equipmentHandler.Delete).Methods("DELETE")
	router.HandleFunc("/tours/{id}/equipment/available", equipmentHandler.GetAvailable).Methods("GET")
}

func initTourHandler(tourHandler *handler.TourHandler, router *mux.Router) {
	router.HandleFunc("/tours/tour", tourHandler.GetAll).Methods("GET")
	router.HandleFunc("/tours/tour", tourHandler.Create).Methods("POST")
	router.HandleFunc("/tours/tour/{id}", tourHandler.Get).Methods("GET")
	router.HandleFunc("/tours/tour/{id}", tourHandler.Update).Methods("PUT")
	router.HandleFunc("/tours/tour/{id}", tourHandler.Delete).Methods("DELETE")
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
