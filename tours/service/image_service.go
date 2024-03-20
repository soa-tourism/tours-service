package service

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type ImageService struct {
	ImageDir string
}

func NewImageService() *ImageService {
	imageService := &ImageService{
		ImageDir: "tours-service/tours/resource/images",
	}
	err := imageService.initializeStorageDirectory()
	if err != nil {
		return nil
	}
	return imageService
}

func SaveImages(imageDataList [][]byte) ([]string, error) {
	imageDir := "resource/images"
	if err := os.MkdirAll(imageDir, 0755); err != nil {
		return nil, err
	}

	var imageNames []string
	for _, imageData := range imageDataList {
		imageName := fmt.Sprintf("image_%d_%d.jpg", time.Now().UnixNano(), rand.Intn(1000))
		imagePath := filepath.Join(imageDir, imageName)
		if err := os.WriteFile(imagePath, imageData, 0644); err != nil {
			return nil, err
		}
		imageNames = append(imageNames, imageName)
	}

	var arrayValue strings.Builder
	arrayValue.WriteString("{")
	for i, name := range imageNames {
		if i > 0 {
			arrayValue.WriteString(",")
		}
		arrayValue.WriteString(fmt.Sprintf(`"%s"`, name))
	}
	arrayValue.WriteString("}")

	return []string{arrayValue.String()}, nil
}

func (is *ImageService) initializeStorageDirectory() error {
	if _, err := os.Stat(is.ImageDir); os.IsNotExist(err) {
		err := os.MkdirAll(is.ImageDir, 0755)
		if err != nil {
			return err
		}
	}
	return nil
}
