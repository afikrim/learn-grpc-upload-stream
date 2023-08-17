package entity

import "time"

type (
	Metadata struct {
		ID          uint64     `json:"id,omitempty"`
		Name        string     `json:"name,omitempty"`
		ContentType string     `json:"contentType,omitempty"`
		Extension   string     `json:"extension,omitempty"`
		Size        int64      `json:"size,omitempty"`
		Path        string     `json:"path,omitempty"`
		URL         string     `json:"url,omitempty"`
		CreatedAt   *time.Time `json:"createdAt,omitempty"`
		UpdatedAt   *time.Time `json:"updatedAt,omitempty"`
		DeletedAt   *time.Time `json:"deletedAt,omitempty"`
	}

	CreateMetadataIn struct {
		Name        string `json:"name,omitempty"`
		ContentType string `json:"contentType,omitempty"`
		Extension   string `json:"extension,omitempty"`
		Size        int64  `json:"size,omitempty"`
		Path        string `json:"path,omitempty"`
	}

	CreateMetadataOut struct {
		ID uint64 `json:"id,omitempty"`
	}
)
