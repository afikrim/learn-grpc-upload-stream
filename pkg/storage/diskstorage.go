package storage

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"path"
	"strings"
)

const (
	baseURL = "http://localhost:8080/public"

	pathDivider     = "/"
	storageBasePath = "/tmp/learn-grpc-upload-stream"
)

type (
	diskStorage struct{}
)

func NewFileStorage() Storage {
	return &diskStorage{}
}

func (ds *diskStorage) Store(ctx context.Context, in *StoreIn) (*StoreOut, error) {
	purifiedPath := strings.Trim(in.Path, pathDivider)
	absPath := path.Join(storageBasePath, purifiedPath)
	dirPath := path.Dir(absPath)

	if err := os.MkdirAll(dirPath, 0755); err != nil {
		return nil, err
	}

	f, err := os.OpenFile(absPath, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return nil, err
	}

	defer f.Close()

	n, err := f.WriteAt(in.Content, in.StartIdx)
	if err != nil {
		return nil, err
	}

	err = f.Sync()
	if err != nil {
		return nil, err
	}

	return &StoreOut{
		EndIdx: in.StartIdx + int64(n),
	}, nil
}

func (ds *diskStorage) Retrieve(ctx context.Context, in *RetrieveIn) (*RetrieveOut, error) {
	purifiedPath := strings.Trim(in.Path, pathDivider)
	absPath := path.Join(storageBasePath, purifiedPath)

	content, err := os.ReadFile(absPath)
	if err != nil {
		return nil, err
	}

	return &RetrieveOut{
		Content:     content,
		ContentType: http.DetectContentType(content),
		Size:        len(content),
	}, nil
}

func (ds *diskStorage) GenerateURL(ctx context.Context, in *GenerateURLIn) (*GenerateURLOut, error) {
	purifiedPath := strings.Trim(in.Path, pathDivider)
	absPath := path.Join(storageBasePath, purifiedPath)

	_, err := os.Stat(absPath)
	if err != nil {
		return nil, err
	}

	return &GenerateURLOut{
		URL: fmt.Sprintf("%s%s", baseURL, in.Path),
	}, nil
}
