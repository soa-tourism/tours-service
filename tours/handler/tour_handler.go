package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"tours/dto"
	"tours/service"
)

type TourHandler struct {
	TourService *service.TourService
}

func (handler *TourHandler) GetAll(writer http.ResponseWriter, req *http.Request) {
	tours, err := handler.TourService.FindAllTours()
	writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	tourDtos := make([]dto.TourDto, len(tours))
	for i, tour := range tours {
		tourDtos[i] = dto.MapFromTour(tour)
	}

	err = json.NewEncoder(writer).Encode(tourDtos)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (handler *TourHandler) Get(writer http.ResponseWriter, req *http.Request) {
	idStr := mux.Vars(req)["id"]
	id, err := strconv.ParseInt(idStr, 10, 64)

	tour, err := handler.TourService.FindTour(id)
	writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	err = json.NewEncoder(writer).Encode(dto.MapFromTour(*tour))
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (handler *TourHandler) Create(writer http.ResponseWriter, req *http.Request) {
	var tourDto dto.TourDto
	err := json.NewDecoder(req.Body).Decode(&tourDto)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	tour := tourDto.MapToModel()
	createdTour, err := handler.TourService.CreateTour(tour)
	if err != nil {
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}

	err = json.NewEncoder(writer).Encode(dto.MapFromTour(*createdTour))
	if err != nil {
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
}

func (handler *TourHandler) Update(writer http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	idStr := params["id"]
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	var tourDto dto.TourDto
	err = json.NewDecoder(req.Body).Decode(&tourDto)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	tour := tourDto.MapToModel()
	tour.Id = id
	updatedTour, err := handler.TourService.UpdateTour(tour)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(writer).Encode(dto.MapFromTour(*updatedTour))
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
}

func (handler *TourHandler) Delete(writer http.ResponseWriter, req *http.Request) {
	idStr := mux.Vars(req)["id"]
	id, err := strconv.ParseInt(idStr, 10, 64)

	err = handler.TourService.DeleteTour(id)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(http.StatusOK)
}

func (handler *TourHandler) GetByAuthor(writer http.ResponseWriter, req *http.Request) {
	idStr := mux.Vars(req)["id"]
	id, err := strconv.ParseInt(idStr, 10, 64)

	tours, err := handler.TourService.FindToursByAuthor(id)
	writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	tourDtos := make([]dto.TourDto, len(tours))
	for i, tour := range tours {
		tourDtos[i] = dto.MapFromTour(tour)
	}

	err = json.NewEncoder(writer).Encode(tourDtos)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (handler *TourHandler) AddEquipment(writer http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	idStr := params["id"]
	equipmentIdStr := params["equipmentId"]
	id, err := strconv.ParseInt(idStr, 10, 64)
	equipmentId, err := strconv.ParseInt(equipmentIdStr, 10, 64)

	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	err = handler.TourService.AddEquipment(id, equipmentId)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(http.StatusOK)
}

func (handler *TourHandler) RemoveEquipment(writer http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	idStr := params["id"]
	equipmentIdStr := params["equipmentId"]
	id, err := strconv.ParseInt(idStr, 10, 64)
	equipmentId, err := strconv.ParseInt(equipmentIdStr, 10, 64)

	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	err = handler.TourService.RemoveEquipment(id, equipmentId)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(http.StatusOK)
}
