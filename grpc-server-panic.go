package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/golang/protobuf/ptypes/timestamp"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/tckz/trivial/api"
	"google.golang.org/grpc"
)

type HelloService struct {
}

func TimestampPB(t time.Time) *timestamp.Timestamp {
	ts := &timestamp.Timestamp{
		Seconds: t.Unix(),
		Nanos:   int32(t.Nanosecond()),
	}
	return ts
}

func (s *HelloService) SayHello(ctx context.Context, in *api.HelloRequest) (*api.HelloReply, error) {
	return &api.HelloReply{
		Message: fmt.Sprintf("Hello %s", in.Name),
		Now:     TimestampPB(time.Now()),
	}, nil
}

func (s *HelloService) SayMorning(ctx context.Context, in *api.MorningRequest) (*api.MorningReply, error) {
	panic("panic!!!!")
}

func tryRequestSet(server *grpc.Server) {
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
		if err != nil {
			log.Printf("*** SayMorning: err=%v", err)
		} else {
			log.Printf("SayMorning: res=%s", res)
		}
	}

	// recoveryなしだとここには来ない。panicした段階で終わっている
	{
		res, err := client.SayHello(context.Background(), &api.HelloRequest{Name: "2nd req"})
		if err != nil {
			log.Printf("*** SayHello: err=%v", err)
		} else {
			log.Printf("SayHello: res=%s", res)
		}
	}
}

func main() {

	// recoveryあり -> panic後も継続
	log.Print("with recovery")
	server := grpc.NewServer(
		grpc_middleware.WithUnaryServerChain(grpc_recovery.UnaryServerInterceptor()),
	)
	tryRequestSet(server)

	// recoveryなし -> panic後の処理は実行されない。プロセスが終わる
	log.Print("without recovery")
	server = grpc.NewServer()
	tryRequestSet(server)

	// ゴールしない
	log.Printf("Goal!!")
}
