package main

import (
	"context"
	"io/fs"
	"net/http"
	"os"

	"github.com/evanhongo/happy-golang/pkg/logger"
	pb "github.com/evanhongo/happy-golang/rpc/job"
)

func main() {
	client := pb.NewJobServiceProtobufClient("http://localhost:4000", &http.Client{})
	file, err := fs.ReadFile(os.DirFS("."), "asset/test.jpeg")
	if err != nil {
		logger.Error(err)
	}
	resp, err := client.CompressImage(context.Background(), &pb.CompressImageReq{Data: file})
	if err != nil {
		logger.Infof("oh no: %v", err)
		os.Exit(1)
	}
	logger.Infof("%+v", resp)
}
