package storage

import "context"

type (
	storage struct{}
)

func UnimplementedStorage() Storage {
	return &storage{}
}

func (s storage) Store(ctx context.Context, in *StoreIn) (*StoreOut, error) {
	//TODO implement me
	panic("implement me")
}

func (s storage) Retrieve(ctx context.Context, in *RetrieveIn) (*RetrieveOut, error) {
	//TODO implement me
	panic("implement me")
}

func (s storage) GenerateURL(ctx context.Context, in *GenerateURLIn) (*GenerateURLOut, error) {
	//TODO implement me
	panic("implement me")
}
