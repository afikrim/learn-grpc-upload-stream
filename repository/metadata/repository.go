package metadata

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/afikrim/learn-grpc-upload-stream/core/entity"
	"github.com/afikrim/learn-grpc-upload-stream/core/repository"
)

type (
	repo struct {
		db *sql.DB
	}
)

func New(db *sql.DB) repository.MetadataRepository {
	return &repo{db: db}
}

func (r *repo) CreateMetadata(ctx context.Context, in *entity.CreateMetadataIn) (*entity.CreateMetadataOut, error) {
	sqlStmt := `INSERT INTO %s(%s,%s,%s,%s,%s) VALUES (?,?,?,?,?)`
	sqlStmt = fmt.Sprintf(
		sqlStmt,
		MetadataTable,
		MetadataTableNameColumn,
		MetadataTableContentTypeColumn,
		MetadataTableExtensionColumn,
		MetadataTableSizeColumn,
		MetadataTablePathColumn,
	)

	stmt, err := r.db.PrepareContext(ctx, sqlStmt)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	res, err := stmt.ExecContext(ctx, in.Name, in.ContentType, in.Extension, in.Size, in.Path)
	if err != nil {
		return nil, err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, err
	}
	if rows != 1 {
		return nil, fmt.Errorf("rows expected to be 1, rows: %v", rows)
	}

	lastId, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	return &entity.CreateMetadataOut{ID: uint64(lastId)}, nil
}
