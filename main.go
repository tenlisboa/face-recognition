package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/Kagami/go-face"
)

var (
	modelsDir = filepath.Join(".", "models")
	imagesDir = filepath.Join(".", "images")
)

func isAValidImage(fileName string) bool {
	validFileTypes := []string{".jpg", ".jpeg"}
	for _, fileType := range validFileTypes {
		if strings.Contains(fileName, fileType) {
			return true
		}
	}

	return false
}

func getReferencesPath(referenceDir string) []string {
	refPaths := []string{}

	err := filepath.Walk(referenceDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if isAValidImage(info.Name()) && !info.IsDir() {
			refPaths = append(refPaths, path)
		}

		return nil
	})

	if err != nil {
		log.Fatalf("Error in getReferencesPath: %v", err)
	}

	return refPaths
}

func getReferences(recognizer *face.Recognizer, referencesPaths []string) ([]face.Face, error) {
	faces := []face.Face{}
	for _, refPath := range referencesPaths {
		testFace, err := recognizer.RecognizeSingleFile(refPath)
		if err != nil {
			return nil, err
		}

		faces = append(faces, *testFace)
	}

	return faces, nil
}

func searchForReferenceFaces(recognizer *face.Recognizer, testingImagePath string, referencesPaths []string) ([]string, error) {
	foundFaces, err := recognizer.RecognizeFile(testingImagePath)
	if err != nil {
		return nil, err
	}
	if len(foundFaces) <= 0 {
		return nil, nil
	}

	var matchingReferencesPath []string
	for _, ref := range foundFaces {
		desc := recognizer.Classify(ref.Descriptor)
		matchingReferencesPath = append(matchingReferencesPath, referencesPaths[desc])
	}

	return matchingReferencesPath, nil
}

func getSamples(references []face.Face) ([]face.Descriptor, []int32) {
	var samples []face.Descriptor
	var samplesIndexes []int32
	for index, ref := range references {
		samples = append(samples, ref.Descriptor)
		samplesIndexes = append(samplesIndexes, int32(index))
	}

	return samples, samplesIndexes
}

func main() {
	rec, err := face.NewRecognizer(modelsDir)
	if err != nil {
		log.Fatalf("Can't init face reconizer, %v", err)
	}
	defer rec.Close()

	referencesDir := filepath.Join(imagesDir, "references")
	referencesPaths := getReferencesPath(referencesDir)
	references, err := getReferences(rec, referencesPaths)
	if err != nil {
		log.Fatalf("Can't get references: %v", err)
	}
	samples, indexes := getSamples(references)
	rec.SetSamples(samples, indexes)

	testingImg := filepath.Join(imagesDir, "test.jpg")
	matchingReferences, err := searchForReferenceFaces(rec, testingImg, referencesPaths)
	if err != nil {
		log.Fatalf("Can't search for references: %v", err)
	}
	if len(matchingReferences) <= 0 {
		log.Fatalf("No matching references in image: %s", testingImg)
	}

	fmt.Printf("We found %v references in your picture: %s", len(matchingReferences), strings.Join(matchingReferences[:], ", "))
}
