package metadata

import (
	"database/sql"
)

const (
	MetadataTable                  = "metadatas"
	MetadataTableNameColumn        = "name"
	MetadataTableContentTypeColumn = "content_type"
	MetadataTableExtensionColumn   = "extension"
	MetadataTableSizeColumn        = "size"
	MetadataTablePathColumn        = "path"
)

type (
	Metadata struct {
		ID          uint64         `json:"id,omitempty"`
		Name        string         `json:"name,omitempty"`
		ContentType string         `json:"contentType,omitempty"`
		Extension   string         `json:"extension,omitempty"`
		Size        int64          `json:"size,omitempty"`
		Path        sql.NullString `json:"path,omitempty"`
		URL         sql.NullString `json:"url,omitempty"`
		CreatedAt   sql.NullTime   `json:"createdAt,omitempty"`
		UpdatedAt   sql.NullTime   `json:"updatedAt,omitempty"`
		DeletedAt   sql.NullTime   `json:"deletedAt,omitempty"`
	}
)
