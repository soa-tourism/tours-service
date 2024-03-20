package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"tours/dto"
	"tours/service"
)

type TourReviewHandler struct {
	TourReviewService *service.TourReviewService
}

func (handler *TourReviewHandler) GetAllByTourist(writer http.ResponseWriter, req *http.Request) {
	idStr := mux.Vars(req)["id"]
	id, err := strconv.ParseInt(idStr, 10, 64)

	tourReview, err := handler.TourReviewService.FindAllTourReviewsByTourist(id)
	writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(writer).Encode(tourReview)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (handler *TourReviewHandler) GetAllByAuthor(writer http.ResponseWriter, req *http.Request) {
	idStr := mux.Vars(req)["id"]
	id, err := strconv.ParseInt(idStr, 10, 64)

	tourReview, err := handler.TourReviewService.FindAllTourReviewsByAuthor(id)
	writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(writer).Encode(tourReview)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (handler *TourReviewHandler) Get(writer http.ResponseWriter, req *http.Request) {
	idStr := mux.Vars(req)["id"]
	id, err := strconv.ParseInt(idStr, 10, 64)

	tourReview, err := handler.TourReviewService.FindTourReview(id)
	writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	err = json.NewEncoder(writer).Encode(tourReview)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (handler *TourReviewHandler) Create(writer http.ResponseWriter, req *http.Request) {
	var tourReviewDto dto.TourReviewDto
	err := json.NewDecoder(req.Body).Decode(&tourReviewDto)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	imagePaths, err := service.SaveImages(tourReviewDto.Images)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	tour := dto.MapToModel(&tourReviewDto)
	tour.Images = imagePaths

	review, err := handler.TourReviewService.CreateTourReview(tour)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(writer).Encode(review)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
}
