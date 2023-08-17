package service

import (
	"context"

	"github.com/afikrim/learn-grpc-upload-stream/core/entity"
)

type (
	UploadService interface {
		UploadFile(ctx context.Context, in *entity.UploadFileIn) (*entity.UploadFileOut, error)
	}
)
