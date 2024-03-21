package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"tours/dto"
	"tours/service"

	"github.com/gorilla/mux"
)

type TouristPositionHandler struct {
	TouristPositionService *service.TouristPositionService
}

func (handler *TouristPositionHandler) GetAll(writer http.ResponseWriter, req *http.Request) {
	TouristPositions, err := handler.TouristPositionService.FindAllTouristPositions()
	writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	TouristPositionDtos := make([]dto.TouristPositionDto, len(TouristPositions))
	for i, TouristPosition := range TouristPositions {
		TouristPositionDtos[i] = dto.TouristPositionDtoFromModel(TouristPosition)
	}

	err = json.NewEncoder(writer).Encode(TouristPositionDtos)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (handler *TouristPositionHandler) Get(writer http.ResponseWriter, req *http.Request) {
	idStr := mux.Vars(req)["id"]
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	TouristPosition, err := handler.TouristPositionService.FindTouristPosition(id)
	writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	err = json.NewEncoder(writer).Encode(dto.TouristPositionDtoFromModel(*TouristPosition))
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (handler *TouristPositionHandler) Create(writer http.ResponseWriter, req *http.Request) {
	var TouristPositionDto dto.TouristPositionDto
	err := json.NewDecoder(req.Body).Decode(&TouristPositionDto)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	TouristPosition := TouristPositionDto.MapToModel()
	createdTouristPosition, err := handler.TouristPositionService.CreateTouristPosition(TouristPosition)
	if err != nil {
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}

	err = json.NewEncoder(writer).Encode(dto.TouristPositionDtoFromModel(*createdTouristPosition))
	if err != nil {
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
}

func (handler *TouristPositionHandler) Update(writer http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	idStr := params["id"]
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	var TouristPositionDto dto.TouristPositionDto
	err = json.NewDecoder(req.Body).Decode(&TouristPositionDto)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	TouristPosition := TouristPositionDto.MapToModel()
	TouristPosition.Id = id
	updatedTouristPosition, err := handler.TouristPositionService.UpdateTouristPosition(TouristPosition)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(writer).Encode(dto.TouristPositionDtoFromModel(*updatedTouristPosition))
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
}

func (handler *TouristPositionHandler) Delete(writer http.ResponseWriter, req *http.Request) {
	idStr := mux.Vars(req)["id"]
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	err = handler.TouristPositionService.DeleteTouristPosition(id)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(http.StatusOK)
}

// func (handler *TouristPositionHandler) GetByCreator(writer http.ResponseWriter, req *http.Request) {
// 	idStr := mux.Vars(req)["id"]
// 	id, err := strconv.ParseInt(idStr, 10, 64)

// 	pos, err := handler.TouristPositionService.FindPositionByCreator(id)
// 	writer.Header().Set("Content-Type", "application/json")
// 	if err != nil {
// 		writer.WriteHeader(http.StatusNotFound)
// 		return
// 	}

// 	err = json.NewEncoder(writer).Encode(dto.TouristPositionDtoFromModel(*pos))
// 	if err != nil {
// 		writer.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}
// }
