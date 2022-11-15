package recognition

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/Kagami/go-face"
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

func GetReferencesPath(referenceDir string) []string {
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

func GetReferences(recognizer *face.Recognizer, referencesPaths []string) ([]face.Face, error) {
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

func SearchForReferenceFaces(recognizer *face.Recognizer, testingImagePath string, referencesPaths []string) ([]string, error) {
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

func GetSamples(references []face.Face) ([]face.Descriptor, []int32) {
	var samples []face.Descriptor
	var samplesIndexes []int32
	for index, ref := range references {
		samples = append(samples, ref.Descriptor)
		samplesIndexes = append(samplesIndexes, int32(index))
	}

	return samples, samplesIndexes
}
