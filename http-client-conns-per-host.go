package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/aws/aws-xray-sdk-go/strategy/sampling"
	"github.com/aws/aws-xray-sdk-go/xray"
	"github.com/aws/aws-xray-sdk-go/xraylog"
	vh "github.com/tckz/vegetahelper"
	vegeta "github.com/tsenart/vegeta/lib"
)

// ちょっとしたhttp負荷かけツール
// デフォルトのhttp.ClientでTCPセッションのはられ方とか確認するためのもの。

type nopWriteCloser struct {
	io.Writer
}

func (c nopWriteCloser) Close() error {
	return nil
}

func openResultFile(out string) (io.WriteCloser, error) {
	switch out {
	case "stdout":
		return &nopWriteCloser{os.Stdout}, nil
	default:
		return os.Create(out)
	}
}

var (
	optDuration = flag.Duration("duration", 10*time.Second, "Duration of the test [0 = forever]")
	optOutput   = flag.String("output", "", "/path/to/results.bin or 'stdout'")
	optTarget   = flag.String("target", "http://localhost:8082", "")
	optTimeout  = flag.Duration("timeout", time.Second*3, "timeout for client")
	optWorkers  = flag.Uint64("workers", 30, "number of workers")
	optXray     = flag.Bool("xray", false, "use xray")
	optRate     = &vh.RateFlag{
		Rate: &vegeta.Rate{
			Freq: 30,
			Per:  1 * time.Second,
		}}
)

type samplingAlways struct{}

var _ sampling.Strategy = (*samplingAlways)(nil)

func (s *samplingAlways) ShouldTrace(r *sampling.Request) *sampling.Decision {
	return &sampling.Decision{Sample: true}
}

type xrayLogger struct{}

func (l xrayLogger) Log(level xraylog.LogLevel, msg fmt.Stringer) {
	if level >= xraylog.LogLevelInfo {
		log.Println(msg.String())
	}
}

var _ xraylog.Logger = (*xrayLogger)(nil)

func init() {
	rand.Seed(time.Now().UnixNano())

	flag.Var(optRate, "rate", "Number of requests per time unit")
	flag.Parse()

	if *optXray {
		xray.SetLogger(&xrayLogger{})
		err := xray.Configure(xray.Config{
			SamplingStrategy: &samplingAlways{},
		})
		if err != nil {
			panic(err)
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cl := &http.Client{
		Timeout: *optTimeout,
	}
	if *optXray {
		cl = xray.Client(cl)
	}

	atk := vh.NewAttacker(func(ctx context.Context) (result *vh.HitResult, err error) {
		if *optXray {
			c, seg := xray.BeginSegment(ctx, "http-client-conns-per-host")
			defer seg.Close(err)
			ctx = c
		}

		req, err := http.NewRequest(http.MethodGet, *optTarget, nil)
		if err != nil {
			return nil, err
		}
		req = req.WithContext(ctx)
		res, err := cl.Do(req)
		if err != nil {
			return nil, err
		}

		defer func() {
			defer res.Body.Close()
			io.Copy(ioutil.Discard, res.Body)
		}()

		n, err := io.Copy(ioutil.Discard, res.Body)
		if err != nil {
			return nil, err
		}

		ret := &vh.HitResult{
			SentBytes: 0,
			RecvBytes: uint64(n),
			Code:      uint16(res.StatusCode),
			Error:     "",
		}
		return ret, nil
	}, vh.WithWorkers(*optWorkers))
	res := atk.Attack(ctx, *optRate.Rate, *optDuration, "http-client-conns-per-host")

	out, err := openResultFile(*optOutput)
	if err != nil {
		panic(err)
	}
	defer out.Close()
	enc := vegeta.NewEncoder(out)

	sig := make(chan os.Signal)
	signal.Notify(sig, syscall.SIGINT)

loop:
	for {
		select {
		case s := <-sig:
			log.Printf("Received signal: %s\n", s)
			cancel()
			// keep loop until 'res' is closed.
		case r, ok := <-res:
			if !ok {
				break loop
			}
			if err := enc.Encode(r); err != nil {
				log.Printf("*** Encode: %v\n", err)
				break loop
			}
		}
	}

	cancel()

}
