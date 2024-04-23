package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"tours/dto"
	"tours/service"
)

type PublishedTourHandler struct {
	PublishedTourService *service.PublishedTourService
}

func (handler *PublishedTourHandler) GetAllPublished(writer http.ResponseWriter, req *http.Request) {
	tours, err := handler.PublishedTourService.FindAll()
	writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	tourDtos := make([]dto.PublishedTourDto, len(tours))
	for i, tour := range tours {
		tourDtos[i] = dto.MapToPublishedTour(tour)
	}

	err = json.NewEncoder(writer).Encode(tourDtos)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (handler *PublishedTourHandler) Get(writer http.ResponseWriter, req *http.Request) {
	idStr := mux.Vars(req)["id"]
	id, err := primitive.ObjectIDFromHex(idStr)

	tour, err := handler.PublishedTourService.Find(id)
	writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	err = json.NewEncoder(writer).Encode(dto.MapToPublishedTour(*tour))
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
}
