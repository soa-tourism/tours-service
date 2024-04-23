package handler

import (
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"tours/dto"
	"tours/service"

	"github.com/gorilla/mux"
)

type PublicCheckpointHandler struct {
	CheckpointService *service.PublicCheckpointService
}

func (handler *PublicCheckpointHandler) GetAll(writer http.ResponseWriter, req *http.Request) {
	checkpoints, err := handler.CheckpointService.FindAllPublicCheckpoints()
	writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	checkpointDtos := make([]dto.PublicCheckpointDto, len(checkpoints))
	for i, checkpoint := range checkpoints {
		checkpointDtos[i] = dto.PublicCheckpointDtoFromModel(checkpoint)
	}

	err = json.NewEncoder(writer).Encode(checkpointDtos)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (handler *PublicCheckpointHandler) Get(writer http.ResponseWriter, req *http.Request) {
	idStr := mux.Vars(req)["id"]
	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	checkpoint, err := handler.CheckpointService.FindPublicCheckpoint(id)
	writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	err = json.NewEncoder(writer).Encode(dto.PublicCheckpointDtoFromModel(*checkpoint))
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (handler *PublicCheckpointHandler) Create(writer http.ResponseWriter, req *http.Request) {
	var checkpointDto dto.PublicCheckpointDto
	err := json.NewDecoder(req.Body).Decode(&checkpointDto)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	checkpoint := checkpointDto.MapToModel()
	createdCheckpoint, err := handler.CheckpointService.CreatePublicCheckpoint(checkpoint)
	if err != nil {
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}

	err = json.NewEncoder(writer).Encode(dto.PublicCheckpointDtoFromModel(*createdCheckpoint))
	if err != nil {
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
}

func (handler *PublicCheckpointHandler) Update(writer http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	idStr := params["id"]
	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	var checkpointDto dto.PublicCheckpointDto
	err = json.NewDecoder(req.Body).Decode(&checkpointDto)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	checkpoint := checkpointDto.MapToModel()
	checkpoint.ID = id
	updatedCheckpoint, err := handler.CheckpointService.UpdatePublicCheckpoint(checkpoint)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(writer).Encode(dto.PublicCheckpointDtoFromModel(*updatedCheckpoint))
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
}

func (handler *PublicCheckpointHandler) Delete(writer http.ResponseWriter, req *http.Request) {
	idStr := mux.Vars(req)["id"]
	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	err = handler.CheckpointService.DeletePublicCheckpoint(id)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(http.StatusOK)
}
