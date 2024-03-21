package handler

import (
	"io"
	"mime"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/mux"
)

type ImageHandler struct {
	ImageDir string
}

func NewImageHandler() *ImageHandler {
	return &ImageHandler{ImageDir: "resource/images"}
}

func (ih *ImageHandler) ServeImage(w http.ResponseWriter, r *http.Request) {
	imageName := mux.Vars(r)["image"]
	imagePath := filepath.Join(ih.ImageDir, imageName)

	imageFile, err := os.Open(imagePath)
	if err != nil {
		http.Error(w, "Image not found", http.StatusNotFound)
		return
	}
	defer imageFile.Close()

	contentType := mime.TypeByExtension(filepath.Ext(imagePath))
	w.Header().Set("Content-Type", contentType)
	w.Header().Set("Content-Security-Policy", "default-src 'self' *")

	if _, err := io.Copy(w, imageFile); err != nil {
		http.Error(w, "Failed to send image", http.StatusInternalServerError)
		return
	}
}
