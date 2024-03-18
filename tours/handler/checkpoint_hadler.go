package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"tours/dto"
	"tours/service"

	"github.com/gorilla/mux"
)

type CheckpointHandler struct {
	CheckpointService *service.CheckpointService
}

func (handler *CheckpointHandler) GetAll(writer http.ResponseWriter, req *http.Request) {
	checkpoints, err := handler.CheckpointService.FindAllCheckpoints()
	writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	checkpointDtos := make([]dto.CheckpointDto, len(checkpoints))
	for i, checkpoint := range checkpoints {
		checkpointDtos[i] = dto.CheckpointDtoFromModel(checkpoint)
	}

	err = json.NewEncoder(writer).Encode(checkpointDtos)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (handler *CheckpointHandler) Get(writer http.ResponseWriter, req *http.Request) {
	idStr := mux.Vars(req)["id"]
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	checkpoint, err := handler.CheckpointService.FindCheckpoint(id)
	writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	err = json.NewEncoder(writer).Encode(dto.CheckpointDtoFromModel(*checkpoint))
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (handler *CheckpointHandler) Create(writer http.ResponseWriter, req *http.Request) {
	var checkpointDto dto.CheckpointDto
	err := json.NewDecoder(req.Body).Decode(&checkpointDto)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	checkpoint := checkpointDto.MapToModel()
	createdCheckpoint, err := handler.CheckpointService.CreateCheckpoint(checkpoint)
	if err != nil {
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}

	err = json.NewEncoder(writer).Encode(dto.CheckpointDtoFromModel(*createdCheckpoint))
	if err != nil {
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
}

func (handler *CheckpointHandler) Update(writer http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	idStr := params["id"]
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	var checkpointDto dto.CheckpointDto
	err = json.NewDecoder(req.Body).Decode(&checkpointDto)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	checkpoint := checkpointDto.MapToModel()
	checkpoint.Id = id
	updatedCheckpoint, err := handler.CheckpointService.UpdateCheckpoint(checkpoint)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(writer).Encode(dto.CheckpointDtoFromModel(*updatedCheckpoint))
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
}

func (handler *CheckpointHandler) Delete(writer http.ResponseWriter, req *http.Request) {
	idStr := mux.Vars(req)["id"]
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	err = handler.CheckpointService.DeleteCheckpoint(id)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(http.StatusOK)
}

func (handler *CheckpointHandler) GetByTour(writer http.ResponseWriter, req *http.Request) {
	idStr := mux.Vars(req)["id"]
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	checkpoints, err := handler.CheckpointService.FindCheckpointsByTour(id)
	writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	checkpointDtos := make([]dto.CheckpointDto, len(checkpoints))
	for i, ch := range checkpoints {
		checkpointDtos[i] = dto.CheckpointDtoFromModel(ch)
	}

	err = json.NewEncoder(writer).Encode(checkpointDtos)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
}
