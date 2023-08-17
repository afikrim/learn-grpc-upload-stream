package v1

import (
	"io"

	"github.com/afikrim/learn-grpc-upload-stream/core/service"
	pb "github.com/afikrim/learn-grpc-upload-stream/handler/pb/api/v1"
)

type (
	Server struct {
		pb.UnimplementedUploadStreamV1Server

		uploadSvc service.UploadService
	}
)

func New(uploadSvc service.UploadService) pb.UploadStreamV1Server {
	return &Server{
		uploadSvc: uploadSvc,
	}
}

func (s Server) UploadFile(fileServer pb.UploadStreamV1_UploadFileServer) error {
	var resp *pb.UploadFileResponse

	size := int64(0)
	for {
		req, err := fileServer.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		in := mapUploadFileRequestToEntity(req)
		in.LastIdx = size

		out, err := s.uploadSvc.UploadFile(fileServer.Context(), in)
		if err != nil {
			return err
		}

		size += int64(len(in.Content))
		resp = mapUploadFileOutToPb(out)
	}

	return fileServer.SendAndClose(resp)
}
