package service

import (
	"fmt"

	"github.com/evanhongo/happy-golang/pkg/util/os"
	"github.com/google/uuid"
	"github.com/h2non/bimg"
)

type JobService struct {
}

func (service *JobService) CompressImage(data []byte) (string, error) {
	// business logic
	fileName := uuid.New().String() + ".webp"
	err := os.CreateFolder("asset")
	if err != nil {
		return fileName, err
	}
	converted, err := bimg.NewImage(data).Convert(bimg.WEBP)
	if err != nil {
		return fileName, err
	}
	processed, err := bimg.NewImage(converted).Process(bimg.Options{Quality: 60})
	if err != nil {
		return fileName, err
	}

	error := bimg.Write(fmt.Sprintf("./asset"+"/%s", fileName), processed)
	if error != nil {
		return fileName, err
	}

	return fileName, nil
}

func NewJobService() IJobService {
	return &JobService{}
}
