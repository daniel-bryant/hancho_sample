package main

import (
  "flag"
  "log"
  "net"

  "golang.org/x/net/context"
  "google.golang.org/grpc"
  pb "github.com/daniel-bryant/honcho_sample/service_manager/gen/go/math"
  "google.golang.org/grpc/reflection"
)

// server is used to implement math.MathServer.
type server struct{}

// SayHello implements math.MathServer
func (s *server) Add(ctx context.Context, in *pb.IntPair) (*pb.Int, error) {
  sum := in.FirstValue + in.SecondValue
  return &pb.Int{Value: sum}, nil
}

func main() {
  port := flag.String("port", "50051", "app port to listen on")
  flag.Parse()

  lis, err := net.Listen("tcp", ":" + *port)
  if err != nil {
    log.Fatalf("failed to listen: %v", err)
  }
  s := grpc.NewServer()
  pb.RegisterMathServer(s, &server{})
  // Register reflection service on gRPC server.
  reflection.Register(s)
  log.Printf("Server listening on port %s", *port)
  if err := s.Serve(lis); err != nil {
    log.Fatalf("failed to serve: %v", err)
  }
}
