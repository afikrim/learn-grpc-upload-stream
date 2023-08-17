package entity

const (
	UploadFileProcessUnknown        = "UNKNOWN"
	UploadFileProcessStoreFile      = "STORE_FILE"
	UploadFileProcessCreateMetadata = "CREATE_METADATA"
)

type (
	UploadFileProcess string
	UploadFileIn      struct {
		Path        string            `json:"path,omitempty"`
		ContentType string            `json:"contentType,omitempty"`
		Extension   string            `json:"extension,omitempty"`
		Content     []byte            `json:"content,omitempty"`
		LastIdx     int64             `json:"lastIdx,omitempty"`
		Process     UploadFileProcess `json:"-"`
	}

	UploadFileOut struct {
		URL string `json:"url,omitempty"`
	}
)
