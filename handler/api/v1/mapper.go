package v1

import (
	"github.com/afikrim/learn-grpc-upload-stream/core/entity"
	pb "github.com/afikrim/learn-grpc-upload-stream/handler/pb/api/v1"
)

func mapUploadFileProcessToEntity(pbUploadFileProcess pb.UploadFileProcess) entity.UploadFileProcess {
	switch pbUploadFileProcess {
	case pb.UploadFileProcess_UPLOAD_FILE_PROCESS_STORE_FILE:
		return entity.UploadFileProcessStoreFile
	case pb.UploadFileProcess_UPLOAD_FILE_PROCESS_CREATE_METADATA:
		return entity.UploadFileProcessCreateMetadata
	default:
		return entity.UploadFileProcessUnknown
	}
}

func mapUploadFileRequestToEntity(pbUploadFileRequest *pb.UploadFileRequest) *entity.UploadFileIn {
	return &entity.UploadFileIn{
		Path:        pbUploadFileRequest.GetPath(),
		ContentType: pbUploadFileRequest.GetContentType(),
		Extension:   pbUploadFileRequest.GetExtension(),
		Content:     pbUploadFileRequest.GetContent(),
		Process:     mapUploadFileProcessToEntity(pbUploadFileRequest.GetProcess()),
	}
}

func mapUploadFileOutToPb(uploadFileOut *entity.UploadFileOut) *pb.UploadFileResponse {
	if uploadFileOut == nil {
		return nil
	}
	if uploadFileOut.URL == "" {
		return nil
	}

	return &pb.UploadFileResponse{Url: uploadFileOut.URL}
}
