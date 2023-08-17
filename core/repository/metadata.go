package repository

import (
	"context"

	"github.com/afikrim/learn-grpc-upload-stream/core/entity"
)

type (
	MetadataRepository interface {
		CreateMetadata(ctx context.Context, in *entity.CreateMetadataIn) (*entity.CreateMetadataOut, error)
	}
)
