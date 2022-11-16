# A face recognition project using Go, dlib and go-face

## Setup
- First of all, follow the instructions here into the go-face [README.md](https://github.com/Kagami/go-face/blob/c482b0e1acfb243025792fd8d59d385dca1c011b/README.md)
- Clone the repository inside your go workspace
- Dowload the dlib models
```sh
cd models
wget https://github.com/Kagami/go-face-testdata/raw/master/models/shape_predictor_5_face_landmarks.dat
wget https://github.com/Kagami/go-face-testdata/raw/master/models/dlib_face_recognition_resnet_model_v1.dat
wget https://github.com/Kagami/go-face-testdata/raw/master/models/mmod_human_face_detector.dat
cd ..
```
- Install project dependencies
```sh
go mod tidy
```

## How it works
Inside the `images/references` folder you will place all the images to train the model, those will be your reference.

Inside `main.go` there is a variable called `testingImg`, it's the variable that will stores the image you will test to see if there's any match with the references.

Once you run, it will take the references and match with the provided image, if is't a match it will return a empty slice.
