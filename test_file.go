package main

import (
	"context"
	"log/slog"
	"net"
	"net/http"

	pb "github.com/aleksander-git/data-analyzer/grpc/gen/version/v1"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"
)

type MicroserviceServer struct {
	pb.UnimplementedMicroserviceServer
}

func (s *MicroserviceServer) GetVersion(ctx context.Context, in *emptypb.Empty) (*pb.VersionResponse, error) {
	return &pb.VersionResponse{Version: "0.0.1"}, nil
}

func main() {
	ch := make(chan error, 2)
	go func(ch chan error) {
		slog.Info("gRPC сервер запущен на порту 50051")
		port, err := net.Listen("tcp", ":50051")
		if err != nil {
			ch <- err
			return
		}
		s := grpc.NewServer()
		reflection.Register(s)
		pb.RegisterMicroserviceServer(s, &MicroserviceServer{})
		if err = s.Serve(port); err != nil {
			ch <- err
		}
	}(ch)

	go func(ch chan error) {
		mux := runtime.NewServeMux()
		err := pb.RegisterMicroserviceHandlerServer(context.Background(), mux, &MicroserviceServer{})
		if err != nil {
			ch <- err
			return
		}
		s := http.Server{
			Addr:    ":8080",
			Handler: mux,
		}
		slog.Info("gRPC-Gateway сервер запущен на порту 8080")
		if err = s.ListenAndServe(); err != nil {
			ch <- err
			return
		}
	}(ch)

	for err := range ch {
		panic(err)
	}
}
