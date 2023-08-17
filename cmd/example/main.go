package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"path/filepath"

	pb "github.com/afikrim/learn-grpc-upload-stream/handler/pb/api/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	defaultSrcPath = "/Users/azizfikrimahmudi/Development/Afikrim/learn-grpc-upload-stream/cmd/example/Screen Shot 2023-02-01 at 18.08.29.png"
	defaultDstPath = "Testing/testing.png"
)

func main() {
	var srcPath, dstPath string
	flag.StringVar(&srcPath, "f", defaultSrcPath, "")
	flag.StringVar(&dstPath, "d", defaultDstPath, "")
	flag.Parse()

	log.Printf("starting upload process, src: %+v, dst: %+v", srcPath, dstPath)

	conn, err := grpc.Dial("localhost:9090", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		log.Fatalf("failed to dial grpc host, error: %+v", err)
		return
	}
	defer conn.Close()

	fileBytes, err := os.ReadFile(srcPath)
	if err != nil {
		log.Fatalf("failed to read file, error: %+v", err)
		return
	}

	pbClient := pb.NewUploadStreamV1Client(conn)
	uploadFileStream, err := pbClient.UploadFile(context.Background())
	if err != nil {
		log.Fatalf("failed to start upload file stream, error: %+v", err)
		return
	}

	chunkBytesSize := 2 * 1000 * 1024
	for cursor := 0; cursor < len(fileBytes); cursor += chunkBytesSize {
		endIndex := cursor + chunkBytesSize
		if endIndex > len(fileBytes) {
			endIndex = len(fileBytes) - 1
		}

		req := &pb.UploadFileRequest{
			Path:    dstPath,
			Content: fileBytes[cursor:endIndex],
			Process: pb.UploadFileProcess_UPLOAD_FILE_PROCESS_STORE_FILE,
		}

		err = uploadFileStream.Send(req)
		if err != nil {
			log.Fatalf("failed sending request, cursor: %+v, error: %+v", cursor, err)
			return
		}
	}

	req := &pb.UploadFileRequest{
		Path:        dstPath,
		ContentType: http.DetectContentType(fileBytes),
		Extension:   filepath.Ext(srcPath),
		Process:     pb.UploadFileProcess_UPLOAD_FILE_PROCESS_CREATE_METADATA,
	}

	err = uploadFileStream.Send(req)
	if err != nil {
		log.Fatalf("failed sending request, error: %+v", err)
		return
	}

	res, err := uploadFileStream.CloseAndRecv()
	if err != nil {
		log.Fatalf("failed receiving response, error: %+v", err)
		return
	}

	log.Printf("successfully receving a response, response: %+v", res)
}
