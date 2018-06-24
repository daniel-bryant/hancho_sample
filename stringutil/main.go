package main

import (
  "log"
  "net"

  "golang.org/x/net/context"
  "google.golang.org/grpc"
  pb "github.com/daniel-bryant/honcho_sample/service_manager/gen/go/stringutil"
  "google.golang.org/grpc/reflection"
)

const (
  port = ":50052"
)

// server is used to implement strutil.StringUtilServer.
type server struct{}

// Reverse implements strutil.StringUtilServer
func (s *server) Reverse(ctx context.Context, in *pb.String) (*pb.String, error) {
  // https://golang.org/doc/code.html#Library
  r := []rune(in.Value)
  for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
    r[i], r[j] = r[j], r[i]
  }
  return &pb.String{Value: string(r)}, nil
}

func main() {
  lis, err := net.Listen("tcp", port)
  if err != nil {
    log.Fatalf("failed to listen: %v", err)
  }
  s := grpc.NewServer()
  pb.RegisterStringUtilServer(s, &server{})
  // Register reflection service on gRPC server.
  reflection.Register(s)
  if err := s.Serve(lis); err != nil {
    log.Fatalf("failed to serve: %v", err)
  }
}
