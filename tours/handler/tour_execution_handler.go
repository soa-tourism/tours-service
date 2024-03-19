package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"tours/model"
	"tours/service"

	"github.com/gorilla/mux"
)

type TourExecutionHandler struct {
	TourExecutionService *service.TourExecutionService
}

func (handler *TourExecutionHandler) GetAllTourExecutions(writer http.ResponseWriter, req *http.Request) {
	tourExecutions, err := handler.TourExecutionService.FindAllTourExecutions()
	writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(writer).Encode(tourExecutions)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (handler *TourExecutionHandler) GetTourExecution(writer http.ResponseWriter, req *http.Request) {
	idStr := mux.Vars(req)["id"]
	id, err := strconv.ParseInt(idStr, 10, 64)
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

	err = json.NewEncoder(writer).Encode(tourExecution)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (handler *TourExecutionHandler) CreateTourExecution(writer http.ResponseWriter, req *http.Request) {
	var tourExecution model.TourExecution
	err := json.NewDecoder(req.Body).Decode(&tourExecution)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	createdTourExecution, err := handler.TourExecutionService.CreateTourExecution(&tourExecution)
	writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}

	err = json.NewEncoder(writer).Encode(createdTourExecution)
	if err != nil {
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusCreated)
}

func (handler *TourExecutionHandler) UpdateTourExecution(writer http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	idStr := params["id"]
	id, err := strconv.ParseInt(idStr, 10, 64)
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

	tourExecution.Id = id
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

func (handler *TourExecutionHandler) DeleteTourExecution(writer http.ResponseWriter, req *http.Request) {
	idStr := mux.Vars(req)["id"]
	id, err := strconv.ParseInt(idStr, 10, 64)
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
