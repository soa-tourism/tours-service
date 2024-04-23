package handler

import (
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"strconv"
	"tours/dto"
	"tours/model"
	"tours/service"

	"github.com/gorilla/mux"
)

type TourExecutionHandler struct {
	TourExecutionService *service.TourExecutionService
}

func (handler *TourExecutionHandler) GetAll(writer http.ResponseWriter, req *http.Request) {
	tourExecutions, err := handler.TourExecutionService.FindAllTourExecutions()
	writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	tourExecutionDtos := make([]dto.TourExecutionDto, len(tourExecutions))
	for i, exec := range tourExecutions {
		tourExecutionDtos[i] = dto.TourExecutionDtoFromModel(exec)
	}

	err = json.NewEncoder(writer).Encode(tourExecutionDtos)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (handler *TourExecutionHandler) Get(writer http.ResponseWriter, req *http.Request) {
	idStr := mux.Vars(req)["id"]
	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	tourExecution, err := handler.TourExecutionService.FindTourExecution(id)
	writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	err = json.NewEncoder(writer).Encode(dto.TourExecutionDtoFromModel(*tourExecution))
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (handler *TourExecutionHandler) Create(writer http.ResponseWriter, req *http.Request) {
	var executionDto dto.TourExecutionDto
	err := json.NewDecoder(req.Body).Decode(&executionDto)
	if err != nil {
		fmt.Println("Error while decoding!")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	tourExecution := executionDto.MapToModel()
	createdTourExecution, err := handler.TourExecutionService.CreateTourExecution(tourExecution)
	writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}

	err = json.NewEncoder(writer).Encode(dto.TourExecutionDtoFromModel(*createdTourExecution))
	if err != nil {
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusCreated)
}

func (handler *TourExecutionHandler) Update(writer http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	idStr := params["id"]
	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	var tourExecution model.TourExecution
	err = json.NewDecoder(req.Body).Decode(&tourExecution)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	tourExecution.ID = id
	updatedTourExecution, err := handler.TourExecutionService.UpdateTourExecution(&tourExecution)
	writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(writer).Encode(updatedTourExecution)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (handler *TourExecutionHandler) Delete(writer http.ResponseWriter, req *http.Request) {
	idStr := mux.Vars(req)["id"]
	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	err = handler.TourExecutionService.DeleteTourExecution(id)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(http.StatusOK)
}

func (handler *TourExecutionHandler) GetByTouristAndTour(writer http.ResponseWriter, req *http.Request) {
	tourIdStr := mux.Vars(req)["tourId"]
	touristIdStr := mux.Vars(req)["touristId"]
	tourId, err := primitive.ObjectIDFromHex(tourIdStr)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	touristId, err := strconv.ParseInt(touristIdStr, 10, 64)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	executions, err := handler.TourExecutionService.FindByTouristAndTour(tourId, touristId)
	writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	executionDtos := make([]dto.TourExecutionDto, len(executions))
	for i, tour := range executions {
		executionDtos[i] = dto.TourExecutionDtoFromModel(tour)
	}

	err = json.NewEncoder(writer).Encode(executionDtos)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (handler *TourExecutionHandler) GetActiveByTouristAndTour(writer http.ResponseWriter, req *http.Request) {
	tourIdStr := mux.Vars(req)["tourId"]
	touristIdStr := mux.Vars(req)["touristId"]
	tourId, err := primitive.ObjectIDFromHex(tourIdStr)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	touristId, err := strconv.ParseInt(touristIdStr, 10, 64)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	execution, err := handler.TourExecutionService.FindActiveByTouristAndTour(tourId, touristId)
	writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	err = json.NewEncoder(writer).Encode(dto.TourExecutionDtoFromModel(execution))
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
}
