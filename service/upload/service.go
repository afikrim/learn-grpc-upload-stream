package upload

import (
	"context"
	"fmt"
	"path"

	"github.com/afikrim/learn-grpc-upload-stream/core/entity"
	"github.com/afikrim/learn-grpc-upload-stream/core/repository"
	svc "github.com/afikrim/learn-grpc-upload-stream/core/service"
	"github.com/afikrim/learn-grpc-upload-stream/pkg/storage"
)

type (
	service struct {
		storage      storage.Storage
		metadataRepo repository.MetadataRepository
	}
)

func New(storage storage.Storage, metadataRepo repository.MetadataRepository) svc.UploadService {
	return &service{
		storage:      storage,
		metadataRepo: metadataRepo,
	}
}

func (s *service) UploadFile(ctx context.Context, in *entity.UploadFileIn) (*entity.UploadFileOut, error) {
	if in.Process == entity.UploadFileProcessStoreFile {
		return nil, s.storeFile(ctx, in)
	}
	if in.Process == entity.UploadFileProcessUnknown {
		return nil, fmt.Errorf("process is not defined, process: %+v", in.Process)
	}

	err := s.createMetadata(ctx, in)
	if err != nil {
		return nil, err
	}

	generateURLOut, err := s.storage.GenerateURL(ctx, &storage.GenerateURLIn{Path: in.Path})
	if err != nil {
		return nil, err
	}

	return &entity.UploadFileOut{URL: generateURLOut.URL}, nil
}

func (s *service) storeFile(ctx context.Context, in *entity.UploadFileIn) error {
	_, err := s.storage.Store(ctx, &storage.StoreIn{
		Path:     in.Path,
		Content:  in.Content,
		StartIdx: in.LastIdx,
	})
	return err
}

func (s *service) createMetadata(ctx context.Context, in *entity.UploadFileIn) error {
	dir, file := path.Split(in.Path)
	_, err := s.metadataRepo.CreateMetadata(ctx, &entity.CreateMetadataIn{
		Name:        file,
		ContentType: in.ContentType,
		Extension:   in.Extension,
		Size:        in.LastIdx,
		Path:        dir,
	})
	return err
}
