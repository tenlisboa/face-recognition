package controllers

import (
	"github.com/tenlisboa/go-face-recognition/config"
	"github.com/tenlisboa/go-face-recognition/src/domains/entities"
	usecases "github.com/tenlisboa/go-face-recognition/src/domains/use-cases"
)

func RecognitionController() {
	configPaths := entities.Config{
		ModelsDir: config.ModelsDir,
		ImagesDir: config.ImagesDir,
	}

	usecases.RecognizeUsecase(configPaths)
}
