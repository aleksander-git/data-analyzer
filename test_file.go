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

// Создаем свою структуру, которая должна реализовывать интерфейс pb.MicroserviceServer.
// Здесь взят для примера pb.UnimplementedMicroserviceServer, в котором уже все методы реализованы
// Но эти методы реализованы для заглушки. И если мы берем эти методы, то нам нужно реализовать только те
// которые мы хотим посмотреть, а остальные не нужно
// Сделано чисто для примера
type MicroserviceServer struct {
	pb.UnimplementedMicroserviceServer
}

// GetVersion - возвращает версию нашей программы
// *emptypb.Empty - пустая структура заглушка для rpc proto файлов
func (s *MicroserviceServer) GetVersion(ctx context.Context, _ *emptypb.Empty) (*pb.VersionResponse, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		return &pb.VersionResponse{Version: version}, nil
	}
}

const (
	version    = "0.0.1"
	maxServers = 2
)

func main() {
	// Создаем буферизированный канал. Если в него будет записана ошибка
	// то приложение завершится
	ch := make(chan error, maxServers)

	// Запускаем gRPC сервер локально в горутине,
	// что бы запустить еще gRPC-Gateway сервер
	go func(ch chan error) {
		slog.Info("gRPC сервер запущен на порту 50051")
		port, err := net.Listen("tcp", "localhost:50051")
		if err != nil {
			ch <- err
			return
		}
		s := grpc.NewServer()
		// Получаем все методы gRPC сервера. Нужно для того, чтобы программы для работы с API
		// на подобии Postman, без дополнительного импорта proto файлов знали какие методы доступны.
		reflection.Register(s)

		// Регистрируем наш интерфейс pb.MicroserviceServer в gRPC сервере
		pb.RegisterMicroserviceServer(s, &MicroserviceServer{})

		// Стартуем. Если ошибка - пишем в канал ошибок
		if err = s.Serve(port); err != nil {
			ch <- err
		}
	}(ch)

	// Запускаем наш REST сервер
	go func(ch chan error) {
		// Регистрируем наш интерфейс pb.MicroserviceHandler в gRPC-Gateway сервере
		mux := runtime.NewServeMux()
		// Регистрируем наш сервер для его дальнейшего использования
		err := pb.RegisterMicroserviceHandlerServer(context.Background(), mux, &MicroserviceServer{})
		if err != nil {
			ch <- err
			return
		}

		slog.Info("gRPC-Gateway сервер запущен на порту 8080")
		// Создаем сервер с нашим обработчиком
		s := http.Server{
			Addr:              ":8080",
			Handler:           mux,
			ReadHeaderTimeout: 0, // 0 для того, что бы линтер не ругался
		}
		// Стартуем. Если ошибка - пишем в канал ошибок
		if err = s.ListenAndServe(); err != nil {
			ch <- err
		}
	}(ch)

	for err := range ch {
		panic(err)
	}
}
