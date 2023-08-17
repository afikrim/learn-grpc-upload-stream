package main

import (
	"database/sql"
	"log"
	"net"

	v1 "github.com/afikrim/learn-grpc-upload-stream/handler/api/v1"
	pb "github.com/afikrim/learn-grpc-upload-stream/handler/pb/api/v1"
	"github.com/afikrim/learn-grpc-upload-stream/pkg/storage"
	"github.com/afikrim/learn-grpc-upload-stream/repository/metadata"
	"github.com/afikrim/learn-grpc-upload-stream/service/upload"
	_ "github.com/mattn/go-sqlite3"
	"google.golang.org/grpc"
)

func main() {
	db, err := sql.Open("sqlite3", "/Users/azizfikrimahmudi/Development/Afikrim/learn-grpc-upload-stream/test.db")
	if err != nil {
		log.Fatalf("failed to init db connection, error: %v", err)
		return
	}

	lis, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatalf("failed to listen to :9090, error: %v", err)
		return
	}

	strg := storage.New("disk-storage")
	metadataRepo := metadata.New(db)
	svc := upload.New(strg, metadataRepo)

	s := v1.New(svc)
	grpcServer := grpc.NewServer()

	pb.RegisterUploadStreamV1Server(grpcServer, s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
		return
	}
}
