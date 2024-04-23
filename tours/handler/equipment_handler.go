package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
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
	id, err := primitive.ObjectIDFromHex(idStr)

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
	id, err := primitive.ObjectIDFromHex(idStr)

	var equipment model.Equipment
	err = json.NewDecoder(req.Body).Decode(&equipment)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	equipment.ID = id
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
	id, err := primitive.ObjectIDFromHex(idStr)

	err = handler.EquipmentService.DeleteEquipment(id)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(http.StatusOK)
}

func (handler *EquipmentHandler) GetAvailable(writer http.ResponseWriter, req *http.Request) {
	//vars := mux.Vars(req)
	//tourIdStr := vars["id"]
	//if err != nil {
	//	writer.WriteHeader(http.StatusBadRequest)
	//	return
	//}
	//
	//queryValues := req.URL.Query()
	//equipmentIdsStr := queryValues["equipmentIds"]
	//var equipmentIds []primitive.ObjectID
	//for _, idStr := range equipmentIdsStr {
	//	id, err := primitive.ObjectIDFromHex(idStr)
	//	if err != nil {
	//		writer.WriteHeader(http.StatusBadRequest)
	//		return
	//	}
	//	equipmentIds = append(equipmentIds, id)
	//}
	equipment, err := handler.EquipmentService.FindAllEquipment()
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
