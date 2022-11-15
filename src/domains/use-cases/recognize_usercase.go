package usecases

import (
	"fmt"
	"log"
	"path/filepath"
	"strings"

	"github.com/Kagami/go-face"
	"github.com/tenlisboa/go-face-recognition/src/domains/entities"
	"github.com/tenlisboa/go-face-recognition/src/frameworks/recognition"
)

func RecognizeUsecase(config entities.Config) {
	rec, err := face.NewRecognizer(config.ModelsDir)
	if err != nil {
		log.Fatalf("Can't init face reconizer, %v", err)
	}
	defer rec.Close()

	referencesDir := filepath.Join(config.ImagesDir, "references")
	referencesPaths := recognition.GetReferencesPath(referencesDir)
	references, err := recognition.GetReferences(rec, referencesPaths)
	if err != nil {
		log.Fatalf("Can't get references: %v", err)
	}
	samples, indexes := recognition.GetSamples(references)
	rec.SetSamples(samples, indexes)

	testingImg := filepath.Join(config.ImagesDir, "test.jpg")
	matchingReferences, err := recognition.SearchForReferenceFaces(rec, testingImg, referencesPaths)
	if err != nil {
		log.Fatalf("Can't search for references: %v", err)
	}
	if len(matchingReferences) <= 0 {
		log.Fatalf("No matching references in image: %s", testingImg)
	}

	fmt.Printf("We found %v references in your picture: %s", len(matchingReferences), strings.Join(matchingReferences[:], ", "))
}
