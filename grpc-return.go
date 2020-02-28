package main

import (
	"context"
	"fmt"
	"log"
	"net"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/pkg/errors"
	"github.com/tckz/trivial/api"
	"google.golang.org/grpc"
)

type HelloService struct {
}

// 応答が非nilでerrも非nil
// -> clientは応答nil、errはサーバーの返却通り、で受け取る
//    errが非nilなら応答の指定は無視される
func (s *HelloService) SayHello(ctx context.Context, in *api.HelloRequest) (*api.HelloReply, error) {
	return &api.HelloReply{
		Message: fmt.Sprintf("Hello %s", in.Name),
	}, errors.New("Wao!")
}

// 応答がnilでerrもnil
// -> Marshal called with nil でpanic
func (s *HelloService) SayMorning(ctx context.Context, in *api.MorningRequest) (*api.MorningReply, error) {
	return nil, nil
}

func main() {
	server := grpc.NewServer(
		grpc_middleware.WithUnaryServerChain(grpc_recovery.UnaryServerInterceptor()),
	)

	lis, err := net.Listen("tcp", ":0")
	if err != nil {
		log.Fatalf("*** net.Listen: %v", err)
	}

	defer server.GracefulStop()

	svc := &HelloService{}
	api.RegisterGreeterServer(server, svc)

	go func() {
		if err := server.Serve(lis); err != nil {
			log.Fatalf("*** Serve: %v", err)
		}
	}()

	conn, err := grpc.Dial(lis.Addr().String(), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("*** grpc.Dial: %v", err)
	}
	defer conn.Close()

	client := api.NewGreeterClient(conn)

	{
		res, err := client.SayMorning(context.Background(), &api.MorningRequest{Name: "1st req"})
		log.Printf("SayMorning: res=%s, err=%v", res, err)
	}

	{
		res, err := client.SayHello(context.Background(), &api.HelloRequest{Name: "2nd req"})
		log.Printf("SayHello: res=%s, err=%v", res, err)
	}

}
