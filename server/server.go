package main

import (
	"context"
	"net"
	"os"

	ms "github.com/schigh/cposc2017_mathservice"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type server struct{}

func init() {
	log.SetLevel(log.DebugLevel)
}

func main() {
	listener, err := net.Listen("tcp", ":80")
	if err != nil {
		log.Errorf("TCP error: %+v", err)
		os.Exit(-1)
	}

	srv := grpc.NewServer()
	ms.RegisterMathServiceServer(srv, &server{})
	srv.Serve(listener)
}

// Add adds
func (s *server) Add(ctx context.Context, req *ms.AddRequest) (*ms.AddResponse, error) {
	log.Debugf("received call to add %+v and %+v", req.Addend1, req.Addend2)
	sum := req.Addend1 + req.Addend2
	resp := &ms.AddResponse{
		Sum: sum,
	}

	return resp, nil
}

// Average averages
func (s *server) Average(ctx context.Context, req *ms.AverageRequest) (*ms.AverageResponse, error) {
	log.Debugf("received call to average numbers: %+v", req.Numbers)

	count := len(req.Numbers)

	switch count {
	case 0:
		return &ms.AverageResponse{Average: 0}, nil
	case 1:
		return &ms.AverageResponse{Average: float64(req.Numbers[0])}, nil
	}

	var sum float64
	for _, n := range req.Numbers {
		sum += float64(n)
	}

	avg := sum / float64(count)
	return &ms.AverageResponse{Average: avg}, nil
}
