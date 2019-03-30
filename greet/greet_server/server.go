package main

import (
  "context"
  "log"
  "net"
  "fmt"
  "google.golang.org/grpc"
  "grpc-go-course/greet/greetpb"
   )

type server struct{}

func (*server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
  fmt.Printf("Greet function with the request %v", req)
  firstName := req.GetGreeting().GetFirstName()
  result := "Hello " + firstName
  res := &greetpb.GreetResponse{
    Result: result,
  }
  return res, nil
}

func main() {
  fmt.Println("Hello World")
  lis, err := net.Listen("tcp", "0.0.0.0:50051")
  if err != nil {
    log.Fatalf("Failed to listen: %v", err)
  }

  s:=grpc.NewServer()
  greetpb.RegisterGreetServiceServer(s, &server{})

  if err := s.Serve(lis); err != nil {
    log.Fatalf("failed to server: %v", err)
  }
}