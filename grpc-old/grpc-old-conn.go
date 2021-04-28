package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"

	"github.com/tckz/go-trivial/grpc-old/api"
	"google.golang.org/grpc"
)

/*
古いgoogle.golang.org/grpcを確認をしたいので、protoc-gen-goの1.0.0を用意すること。
protoc -I. --go_out=plugins=grpc:. api/hello.proto
*/

type HelloService struct {
}

func (s *HelloService) SayHello(ctx context.Context, in *api.HelloRequest) (*api.HelloReply, error) {
	log.Printf("server: SayHello: Name=%s", in.Name)
	return &api.HelloReply{
		Message: fmt.Sprintf("Hello %s", in.Name),
	}, nil
}

func (s *HelloService) SayMorning(ctx context.Context, in *api.MorningRequest) (*api.MorningReply, error) {
	log.Printf("server: SayMorning: Name=%s", in.Name)
	// -> Marshal called with nil でpanic
	return nil, nil
}

// latencyListener 新しいgrpcだと公式に遅延シミュレートするListenerを持っているけど確認したいgRPCが古すぎて新しいgRPCサーバーに接続できないのでAcceptからの遅延要素だけ真似たもの
type latencyListener struct {
	net.Listener

	latencyFunc func() time.Duration
}

func (l latencyListener) Accept() (net.Conn, error) {
	c, err := l.Listener.Accept()
	if err != nil {
		return nil, err
	}

	d := l.latencyFunc()
	log.Printf("server: Accept: sleep %s", d)
	time.Sleep(d)
	return c, nil
}

func (l latencyListener) Close() error {
	return l.Listener.Close()
}

func (l latencyListener) Addr() net.Addr {
	return l.Listener.Addr()
}

var _ net.Listener = (*latencyListener)(nil)

func wrapListener(l net.Listener, f func() time.Duration) net.Listener {
	return &latencyListener{
		Listener:    l,
		latencyFunc: f,
	}
}

// 古いgrpcのClientConnの挙動を確認するためのもの
func main() {
	log.Printf("args=%#v", os.Args)

	// client
	optAddr := flag.String("addr", "", "addr:port of grpc server")
	optDialWithBlock := flag.Bool("dial-with-block", false, "pass WithBlock option to grpc.Dial")
	optDialWithTimeout := flag.Duration("dial-with-timeout", 0, "pass WithTimeout option to grpc.Dial")
	optTimeout := flag.Duration("timeout", time.Second*3, "timeout for request")
	// server
	optBind := flag.String("bind", ":0", "addr:port to bind")
	optAcceptLatency := flag.Duration("accept-latency", 0, "")
	flag.Parse()

	lis, err := net.Listen("tcp", *optBind)
	if err != nil {
		log.Fatalf("*** net.Listen: %v", err)
	}

	lis = wrapListener(lis, func() time.Duration {
		return *optAcceptLatency
	})

	server := grpc.NewServer()
	defer server.GracefulStop()

	svc := &HelloService{}
	api.RegisterGreeterServer(server, svc)

	go func() {
		log.Printf("Serve: %s", lis.Addr().String())
		if err := server.Serve(lis); err != nil && !strings.Contains(err.Error(), "use of closed network connection") {
			log.Fatalf("*** Serve: %v, %T", err, err)
		}
	}()

	dialOpts := []grpc.DialOption{
		grpc.WithInsecure(),
	}
	if *optDialWithTimeout > 0 {
		// ちなみに新しいgrpcではgrpc.Dialのこのオプションはdeprecated。DialContextを使う
		dialOpts = append(dialOpts, grpc.WithTimeout(*optDialWithTimeout))
	}
	if *optDialWithBlock {
		dialOpts = append(dialOpts, grpc.WithBlock())
	}

	addr := *optAddr
	if addr == "" {
		addr = lis.Addr().String()
	}

	log.Printf("Dial: %s", addr)
	conn, err := grpc.Dial(addr, dialOpts...)
	if err != nil {
		log.Fatalf("grpc.Dial: %v", err)
	}
	defer conn.Close()

	// WithBlockなしだと接続は非同期に行われている
	log.Printf("after grpc.Dial")

	cl := api.NewGreeterClient(conn)
	ctx := context.Background()
	if *optTimeout > 0 {
		c, cancel := context.WithTimeout(ctx, *optTimeout)
		defer cancel()
		ctx = c
	}
	res, err := cl.SayHello(ctx, &api.HelloRequest{
		Name: "myname",
	})
	if err != nil {
		log.Printf("SayHello: err=%v", err)
		return
	}

	log.Printf("res=%s", res)
}
