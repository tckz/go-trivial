package main

import (
	"context"
	"fmt"
	"log"
	"net"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/tckz/trivial/api"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

// gRPCのデバッグログ出力を見るもの

type GrpcLogService struct {
	api.GreeterServer
}

func (s *GrpcLogService) SayHello(ctx context.Context, in *api.HelloRequest) (*api.HelloReply, error) {
	log.Printf("serverside SayHello")
	return &api.HelloReply{
		Message: fmt.Sprintf("Hello %s", in.Name),
	}, nil
}

func main() {
	zl, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("*** zap.NewProduction: %v", err)
	}
	grpc_zap.ReplaceGrpcLoggerV2WithVerbosity(zl, 100)

	server := grpc.NewServer(
		grpc_middleware.WithUnaryServerChain(grpc_recovery.UnaryServerInterceptor()),
	)

	lis, err := net.Listen("tcp", ":0")
	if err != nil {
		log.Fatalf("*** net.Listen: %v", err)
	}

	defer func() {
		log.Printf("Before GracefulStop")
		server.GracefulStop()
		log.Printf("After GracefulStop")
	}()

	svc := &GrpcLogService{}
	api.RegisterGreeterServer(server, svc)

	go func() {
		if err := server.Serve(lis); err != nil {
			log.Fatalf("*** Serve: %v", err)
		}
	}()

	log.Printf("Dial: %s", lis.Addr().String())
	conn, err := grpc.Dial(lis.Addr().String(), grpc.WithInsecure())
	log.Printf("After Dial")
	if err != nil {
		log.Fatalf("*** grpc.Dial: %v", err)
	}
	defer func() {
		log.Printf("Before client.Close")
		err := conn.Close()
		log.Printf("After client.Close: err=%v", err)
	}()

	client := api.NewGreeterClient(conn)

	{
		log.Printf("Before SayHello")
		res, err := client.SayHello(context.Background(), &api.HelloRequest{Name: "client"})
		log.Printf("SayHello: res=%s, err=%v", res, err)
	}

}
