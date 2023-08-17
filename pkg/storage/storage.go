package storage

import "context"

type (
	StoreIn struct {
		Path     string
		Content  []byte
		StartIdx int64
	}

	StoreOut struct {
		EndIdx int64
	}

	RetrieveIn struct {
		Path string
	}

	RetrieveOut struct {
		Content     []byte
		ContentType string
		Size        int
	}

	GenerateURLIn struct {
		Path string
	}

	GenerateURLOut struct {
		URL string
	}

	Storage interface {
		Store(ctx context.Context, in *StoreIn) (*StoreOut, error)
		Retrieve(ctx context.Context, in *RetrieveIn) (*RetrieveOut, error)
		GenerateURL(ctx context.Context, in *GenerateURLIn) (*GenerateURLOut, error)
	}
)

func New(storageType string) Storage {
	switch storageType {
	case "disk-storage":
		return NewFileStorage()
	default:
		return UnimplementedStorage()
	}
}
