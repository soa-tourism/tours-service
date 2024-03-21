package handler

import (
	"encoding/json"
	"fmt"
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

func (handler *CheckpointHandler) UpdateCheckpointSecret(writer http.ResponseWriter, req *http.Request) {
	idStr := mux.Vars(req)["id"]
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	var secretDto dto.CheckpointSecretDto
	err = json.NewDecoder(req.Body).Decode(&secretDto)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	secret, _ := secretDto.MapToModel()
	checkpointDto, err := handler.CheckpointService.UpdateCheckpointSecret(id, secret)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(writer).Encode(checkpointDto)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
}
func (handler *CheckpointHandler) UpdateCheckpointEncounter(writer http.ResponseWriter, req *http.Request) {
	// Parse request params
	vars := mux.Vars(req)
	checkpointID := vars["checkpointId"]
	id, err := strconv.ParseInt(checkpointID, 10, 64)
	if err != nil {
		fmt.Println("Error parsing checkpointId:", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	encounterID := vars["encounterId"]
	encounterId, err := strconv.ParseInt(encounterID, 10, 64)
	if err != nil {
		fmt.Println("Error parsing encounterId:", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	isSecretPrerequisite := vars["isSecretPrerequisite"]
	isSecretBool, err := strconv.ParseBool(isSecretPrerequisite)
	if err != nil {
		fmt.Println("Error parsing boolean:", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	// Get checkpoint by id
	checkpoint, err := handler.CheckpointService.FindCheckpoint(id)
	writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		fmt.Println("Can't find checkpoint with id:", id, err)
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	checkpoint.Id = id
	checkpoint.EncounterId = encounterId
	checkpoint.IsSecretPrerequisite = isSecretBool
	// Update checkpoint
	updatedCheckpoint, err := handler.CheckpointService.UpdateCheckpoint(checkpoint)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	updatedCheckpoint.Id = id

	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
	writer.Write([]byte(`{"message": "Checkpoint updated with encounter"}`))
}
