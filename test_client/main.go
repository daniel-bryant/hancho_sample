package main

import (
  "log"
  "time"

  "golang.org/x/net/context"
  "google.golang.org/grpc"
  helloPb "github.com/daniel-bryant/honcho_sample/service_manager/gen/go/helloworld"
  stringPb "github.com/daniel-bryant/honcho_sample/service_manager/gen/go/stringutil"
  mathPb "github.com/daniel-bryant/honcho_sample/service_manager/gen/go/math"
)

const (
  address = "localhost:50051"
)

func testHelloworld() {
  // Set up a connection to the server.
  conn, err := grpc.Dial(address, grpc.WithInsecure())
  if err != nil {
    log.Fatalf("did not connect: %v", err)
  }
  defer conn.Close()
  c := helloPb.NewGreeterClient(conn)

  // Contact the server and print out its response.
  name := "Tester"
  ctx, cancel := context.WithTimeout(context.Background(), time.Second)
  defer cancel()
  r, err := c.SayHello(ctx, &helloPb.HelloRequest{Name: name})
  if err != nil {
    log.Fatalf("could not greet: %v", err)
  }
  log.Printf("Greeting: %s", r.Message)

  r, err = c.SayHelloAgain(context.Background(), &helloPb.HelloRequest{Name: name})
  if err != nil {
    log.Fatalf("could not greet: %v", err)
  }
  log.Printf("Greeting: %s", r.Message)
}

func testStringUtil() {
  // Set up a connection to the server.
  conn, err := grpc.Dial(address, grpc.WithInsecure())
  if err != nil {
    log.Fatalf("did not connect: %v", err)
  }
  defer conn.Close()
  c := stringPb.NewStringUtilClient(conn)

  // Contact the server and print out its response.
  value := "!desrever yllufsseccuS"
  ctx, cancel := context.WithTimeout(context.Background(), time.Second)
  defer cancel()
  r, err := c.Reverse(ctx, &stringPb.String{Value: value})
  if err != nil {
    log.Fatalf("could not reverse: %v", err)
  }
  log.Printf("Reversed: %s", r.Value)
}

func testMath() {
  // Set up a connection to the server.
  conn, err := grpc.Dial(address, grpc.WithInsecure())
  if err != nil {
    log.Fatalf("did not connect: %v", err)
  }
  defer conn.Close()
  c := mathPb.NewMathClient(conn)

  // Contact the server and print out its response.
  var first_value, second_value int32 = 2,3
  ctx, cancel := context.WithTimeout(context.Background(), time.Second)
  defer cancel()
  r, err := c.Add(ctx, &mathPb.IntPair{FirstValue: first_value, SecondValue: second_value})
  if err != nil {
    log.Fatalf("could not add: %v", err)
  }
  log.Printf("Sum: %d", r.Value)
}

func main() {
  testHelloworld()
  testStringUtil()
  testMath()
}
