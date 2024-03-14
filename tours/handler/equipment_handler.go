package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"tours/model"
	"tours/service"
)

type EquipmentHandler struct {
	EquipmentService *service.EquipmentService
}

func (handler *EquipmentHandler) GetAll(writer http.ResponseWriter, req *http.Request) {
	equipment, err := handler.EquipmentService.FindAllEquipment()
	writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(writer).Encode(equipment)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (handler *EquipmentHandler) Get(writer http.ResponseWriter, req *http.Request) {
	idStr := mux.Vars(req)["id"]
	id, err := strconv.ParseInt(idStr, 10, 64)

	equipment, err := handler.EquipmentService.FindEquipment(id)
	writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	err = json.NewEncoder(writer).Encode(equipment)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (handler *EquipmentHandler) Create(writer http.ResponseWriter, req *http.Request) {
	var equipment model.Equipment
	err := json.NewDecoder(req.Body).Decode(&equipment)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	equip, err := handler.EquipmentService.CreateEquipment(&equipment)
	err = json.NewEncoder(writer).Encode(equip)
	if err != nil {
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
}

func (handler *EquipmentHandler) Update(writer http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	idStr := params["id"]
	id, err := strconv.ParseInt(idStr, 10, 64)

	var equipment model.Equipment
	err = json.NewDecoder(req.Body).Decode(&equipment)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	equipment.Id = id
	equip, err := handler.EquipmentService.UpdateEquipment(&equipment)
	err = json.NewEncoder(writer).Encode(equip)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
}

func (handler *EquipmentHandler) Delete(writer http.ResponseWriter, req *http.Request) {
	idStr := mux.Vars(req)["id"]
	id, err := strconv.ParseInt(idStr, 10, 64)

	err = handler.EquipmentService.DeleteEquipment(id)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(http.StatusOK)
}

func (handler *EquipmentHandler) GetAvailable(writer http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	tourIdStr := vars["id"]
	tourId, err := strconv.ParseInt(tourIdStr, 10, 64)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	queryValues := req.URL.Query()
	equipmentIdsStr := queryValues["equipmentIds"]
	var equipmentIds []int64
	for _, idStr := range equipmentIdsStr {
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			return
		}
		equipmentIds = append(equipmentIds, id)
	}

	equipment, err := handler.EquipmentService.GetAvailableEquipment(tourId, equipmentIds)
	writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	err = json.NewEncoder(writer).Encode(equipment)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
}
