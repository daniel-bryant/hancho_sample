package main

import (
  "flag"
  "log"
  "net"

  "golang.org/x/net/context"
  "google.golang.org/grpc"
  pb "github.com/daniel-bryant/honcho_sample/service_manager/gen/go/helloworld"
  "google.golang.org/grpc/reflection"
)

// server is used to implement helloworld.GreeterServer.
type server struct{}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
  return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

func (s *server) SayHelloAgain(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
  return &pb.HelloReply{Message: "Hello again " + in.Name}, nil
}

func main() {
  port := flag.String("port", "50051", "app port to listen on")
  flag.Parse()

  lis, err := net.Listen("tcp", ":" + *port)
  if err != nil {
    log.Fatalf("failed to listen: %v", err)
  }
  s := grpc.NewServer()
  pb.RegisterGreeterServer(s, &server{})
  // Register reflection service on gRPC server.
  reflection.Register(s)
  if err := s.Serve(lis); err != nil {
    log.Fatalf("failed to serve: %v", err)
  }
}
